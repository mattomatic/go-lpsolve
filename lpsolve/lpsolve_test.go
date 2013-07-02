package lpsolve

import (
    "testing"
    "fmt"
)

func TestZmqStuff(t *testing.T) {
    major, minor, patch := GetZmqVersion()
    fmt.Println("zmg:", major, minor, patch)
}

func TestGetVersion(t *testing.T) {
    info := GetVersion()
    
    if info.Major != 5 {
        t.Error()
    }
    
    if info.Minor != 5 {
        t.Error()
    }
}

func TestMakeLP(t *testing.T) {
    lp := NewLP(2000, 10000)
    
    if lp == nil {
        t.Error()
    }
    
    lp.Delete()
}

func TestPrintLP(t *testing.T) {
    lp := NewLP(5, 10)
    defer lp.Delete()
    lp.Print()
}

func TestLPSolve(t *testing.T) {
    lp := NewLP(2, 2)
    code := lp.Solve()
    
    if code != Optimal {
        t.Error()
    }
    
    lp.Delete()
}