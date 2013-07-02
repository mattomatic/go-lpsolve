package lpsolve
// #cgo pkg-config: libzmq liblpsolve
// #include <zmq.h>
// #include <lp_lib.h>
import "C"

import "fmt"

type VersionInfo struct {
    Major int
    Minor int
    Release int
    Build int
}

type LP struct {
    lprec *C.lprec
}

func NewLP(rows, cols int) (*LP) {
    // err keeps getting set to "no such file or directory" -- ignoring for now
    lprec, _ := C.make_lp(C.int(rows), C.int(cols)) 
    
    if lprec == nil {
        panic("not enough memory to set up structures")
    }
    
    return &LP{lprec}
}

func (lp *LP) Delete() {
    _, err := C.delete_lp(lp.lprec)
    
    if err != nil {
        panic(err.Error())
    }
}

func (lp *LP) Print() {
    _, err := C.print_lp(lp.lprec)
    
    if err != nil {
        panic(err.Error())
    }
}

func (lp *LP) Solve() ResultCode {
    code, err := C.solve(lp.lprec)
    
    if err != nil {
        panic(err.Error())
    }
    
    return ResultCode(code)
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

func GetZmqVersion() (int, int, int) {
    var major, minor, patch C.int
    C.zmq_version(&major, &minor, &patch)
    return int(major), int(minor), int(patch)
}