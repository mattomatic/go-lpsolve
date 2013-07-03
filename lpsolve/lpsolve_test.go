package lpsolve

import (
	"fmt"
	"testing"
)

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
	defer lp.Delete()

	if lp == nil {
		t.Error()
	}
}

func TestPrintLP(t *testing.T) {
	lp := NewLP(5, 10)
	defer lp.Delete()

	lp.Print()
}

func TestLPSolve(t *testing.T) {
	lp := NewLP(2, 2)
	defer lp.Delete()

	code, err := lp.Solve()

	if err != nil {
		t.Error()
	}

	if code != Optimal {
		t.Error()
	}
}

func TestLPSolveTwo(t *testing.T) {
	lp := NewLP(3, 2)
	defer lp.Delete()

	lp.SetMaximize()

	lp.SetValue(1, 1, 120)
	lp.SetValue(1, 2, 210)
	lp.SetRh(1, 15000)
	lp.SetConstraintType(1, LE)

	lp.SetValue(2, 1, 110)
	lp.SetValue(2, 2, 30)
	lp.SetRh(2, 4000)
	lp.SetConstraintType(2, LE)

	lp.SetValue(3, 1, 1)
	lp.SetValue(3, 2, 1)
	lp.SetRh(3, 75)
	lp.SetConstraintType(3, LE)

	lp.SetObjective(1, 143)
	lp.SetObjective(2, 60)

	code, err := lp.Solve()

	if code != Optimal {
		t.Error()
	}

	if err != nil {
		t.Error()
	}

	x, _ := lp.GetVariable(1)
	y, _ := lp.GetVariable(2)

	fmt.Println("x", x, "y", y)
}

func TestLPSolveThree(t *testing.T) {
	lp := NewLP(0, 2) // ok so apparently we need to know this ahead of time!
	defer lp.Delete()

	c1 := &Constraint{[]Real{2, 120, 210}, LE, 15000}
	c2 := &Constraint{[]Real{2, 110, 30}, LE, 4000}
	c3 := &Constraint{[]Real{2, 1, 1}, LE, 75}
	
	lp.SetMaximize()	
	lp.AddConstraint(c1)
	lp.AddConstraint(c2)
	lp.AddConstraint(c3)
	lp.SetObjectiveFunction([]Real{2, 143, 60})

	code, _ := lp.Solve()
	
	lp.Print()
	if code != Optimal {
	    t.Error()
	}
	
	//lp.Delete()
}
