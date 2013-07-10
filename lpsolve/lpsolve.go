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
	Values []Real // note: values should be n + 1 (column 0 is not considered!)
	Type   ConstraintType
	Bound  Real
}

type LP struct {
	lprec     *C.lprec
	variables []Real
}

// ====================================

func NewLP(cols int) *LP {
	// err keeps getting set to "no such file or directory" -- ignoring for now
	lprec, _ := C.make_lp(C.int(0), C.int(cols))
	variables := make([]Real, cols, cols)

	if lprec == nil {
		panic("not enough memory to set up structures")
	}

	lp := &LP{lprec, variables}
	lp.SetVerbosity(Critical) // by default the verbosity is very high

	return lp
}

func (lp *LP) Delete() error {
	_, err := C.delete_lp(lp.lprec)
	return err
}

func (lp *LP) AddConstraint(c *Constraint) error {
	_, err := C.add_constraint(lp.lprec, (*C.REAL)(&c.Values[0]), C.int(c.Type), C.REAL(c.Bound))
	return err
}

func (lp *LP) SetObjectiveFunction(values []Real) error {
	_, err := C.set_obj_fn(lp.lprec, (*C.REAL)(&values[0]))
	return err
}

func (lp *LP) GetVariables() ([]Real, error) {
	_, err := C.get_variables(lp.lprec, (*C.REAL)(&lp.variables[0]))
	return lp.variables, err
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

// ====================================

func (lp *LP) Print() error {
	_, err := C.print_lp(lp.lprec)
	return err
}

func (lp *LP) PrintDuals() error {
	_, err := C.print_duals(lp.lprec)
	return err
}

func (lp *LP) PrintConstraints() error {
	_, err := C.print_constraints(lp.lprec, C.int(len(lp.variables)))
	return err
}

func (lp *LP) PrintTableau() error {
	_, err := C.print_tableau(lp.lprec)
	return err
}

func (lp *LP) PrintSolution() error {
	_, err := C.print_solution(lp.lprec, C.int(len(lp.variables)))
	return err
}

// ====================================

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
