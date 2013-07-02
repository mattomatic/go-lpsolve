package lpsolve

import (
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

	lp.Solve()
	lp.Print()

}
