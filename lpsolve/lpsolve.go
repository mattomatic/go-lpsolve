package lpsolve

// #cgo pkg-config: liblpsolve
// #include <lp_lib.h>
import "C"

import "fmt"

type VersionInfo struct {
	Major   int
	Minor   int
	Release int
	Build   int
}

type Constraint struct {
	Values []Real
	Type   ConstraintType
	Bound  Real
}

type LP struct {
	lprec *C.lprec
}

func NewLP(rows, cols int) *LP {
	// err keeps getting set to "no such file or directory" -- ignoring for now
	lprec, _ := C.make_lp(C.int(rows), C.int(cols))

	if lprec == nil {
		panic("not enough memory to set up structures")
	}

	lp := &LP{lprec}
	//lp.SetVerbosity(Critical) // by default the verbosity is very high
	return lp
}

func (lp *LP) Delete() error {
	_, err := C.delete_lp(lp.lprec)
	return err
}

func (lp *LP) Print() error {
	_, err := C.print_lp(lp.lprec)
	return err
}

func (lp *LP) AddConstraint(c *Constraint) error {
	_, err := C.add_constraint(lp.lprec, (*C.REAL)(&c.Values[0]), C.int(c.Type), C.REAL(c.Bound))
	return err
}

func (lp *LP) SetValue(row, col int, value Real) error {
	_, err := C.set_mat(lp.lprec, C.int(row), C.int(col), C.REAL(value))
	return err
}

func (lp *LP) SetRh(row int, value Real) error {
	_, err := C.set_rh(lp.lprec, C.int(row), C.REAL(value))
	return err
}

func (lp *LP) SetConstraintType(row int, ctype ConstraintType) error {
	_, err := C.set_constr_type(lp.lprec, C.int(row), C.int(ctype))
	return err
}

func (lp *LP) SetObjective(col int, value Real) error {
	_, err := C.set_obj(lp.lprec, C.int(col), C.REAL(value))
	//todo: this function returns bool!
	return err
}

func (lp *LP) SetObjectiveFunction(values []Real) error {
    _, err := C.set_obj_fn(lp.lprec, (*C.REAL)(&values[0]))
    return err
}

func (lp *LP) GetVariable(col int) (Real, error) {
	value, err := C.get_var_primalresult(lp.lprec, C.int(col))
	return Real(value), err
}

func (lp *LP) SetMaximize() error {
	_, err := C.set_maxim(lp.lprec)
	return err
}

func (lp *LP) SetMinimize() error {
	_, err := C.set_minim(lp.lprec)
	return err
}

func (lp *LP) Solve() (SolverStatus, error) {
	code, err := C.solve(lp.lprec)
	return SolverStatus(code), err
}

func (lp *LP) SetVerbosity(level Verbosity) error {
	_, err := C.set_verbose(lp.lprec, C.int(level))
	return err
}

func GetVersion() *VersionInfo {
	var major, minor, release, build C.int
	_, err := C.lp_solve_version(&major, &minor, &release, &build)
	check(err)

	if err != nil {
		panic("error getting version")
	}

	return &VersionInfo{int(major), int(minor), int(release), int(build)}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
}
