package lpsolve

import (
	"math"
	"testing"
)

const (
	Tolerance = 1e-9 // floating point error tolerance
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
	lp := NewLP(10000)
	defer lp.Delete()

	if lp == nil {
		t.Error()
	}
}

func TestPrintLP(t *testing.T) {
	lp := NewLP(10)
	defer lp.Delete()

	lp.Print()
}

func TestLPSolve(t *testing.T) {
	lp := NewLP(2)
	defer lp.Delete()

	code, err := lp.Solve()

	if err != nil {
		t.Error()
	}

	if code != Optimal {
		t.Error()
	}
}

func TestLPSolveApiExample(t *testing.T) {
	lp := NewLP(2) // ok so apparently we need to know this ahead of time!
	defer lp.Delete()

	c1 := NewConstraint(LE, 15000, []Real{120, 210})
	c2 := NewConstraint(LE, 4000, []Real{110, 30})
	c3 := NewConstraint(LE, 75, []Real{1, 1})
	ob := NewObjective(Maximize, []Real{143, 60})

	lp.AddConstraint(c1)
	lp.AddConstraint(c2)
	lp.AddConstraint(c3)
	lp.SetObjective(ob)

	code, _ := lp.Solve()
	variables, _ := lp.GetVariables()

	if code != Optimal {
		t.Error()
	}

	if !floatEquals(variables[0], 21.875) {
		t.Error()
	}

	if !floatEquals(variables[1], 53.125) {
		t.Error()
	}
}

func TestResolve(t *testing.T) {
	lp := NewLP(2) // ok so apparently we need to know this ahead of time!
	defer lp.Delete()

	c1 := NewConstraint(LE, 15000, []Real{120, 210})
	c2 := NewConstraint(LE, 4000, []Real{110, 30})
	c3 := NewConstraint(LE, 75, []Real{1, 1})
	ob := NewObjective(Maximize, []Real{143, 60})

	lp.AddConstraint(c1)
	lp.AddConstraint(c2)
	lp.AddConstraint(c3)
	lp.SetObjective(ob)

	code, _ := lp.Solve()
	variables, _ := lp.GetVariables()

	if code != Optimal {
		t.Error()
	}

	if !floatEquals(variables[0], 21.875) {
		t.Error()
	}

	if !floatEquals(variables[1], 53.125) {
		t.Error()
	}
	
	lp.PrintConstraints()
	
	c1.Bound = 13000
	
	lp.Solve()
	lp.PrintConstraints()
}

func floatEquals(a, b Real) bool {
	return math.Dim(float64(a), float64(b)) <= Tolerance
}
