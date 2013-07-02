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
	lp.SetVerbosity(Critical) // by default the verbosity is very high
	return lp
}

func (lp *LP) Delete() {
	_, err := C.delete_lp(lp.lprec)
	check(err)
}

func (lp *LP) Print() {
	_, err := C.print_lp(lp.lprec)
	check(err)
}

func (lp *LP) SetValue(row, col int, value Real) {
	_, err := C.set_mat(lp.lprec, C.int(row), C.int(col), C.REAL(value))
	check(err)
}

func (lp *LP) SetConstraintType(row int, ctype ConstraintType) {
	_, err := C.set_constr_type(lp.lprec, C.int(row), C.int(ctype))
	check(err)
}

func (lp *LP) Solve() ResultCode {
	code, err := C.solve(lp.lprec)
	check(err)
	return ResultCode(code)
}

func (lp *LP) SetVerbosity(level Verbosity) {
	_, err := C.set_verbose(lp.lprec, C.int(level))
	check(err)
}

func GetVersion() *VersionInfo {
	var major, minor, release, build C.int
	_, err := C.lp_solve_version(&major, &minor, &release, &build)

	fmt.Println(err)

	if err != nil {
		panic("error getting version")
	}

	return &VersionInfo{int(major), int(minor), int(release), int(build)}
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
