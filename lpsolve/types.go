package lpsolve

type SolverStatus int

const (
	UnknownError SolverStatus = iota - 5
	DataIgnored
	NoBFP
	NoMemory
	NotRun
	Optimal
	SubOptimal
	Infeasible
	Unbounded
	Degenerate
	NumericalFailure
	UserAbort
	Timeout
)

// ====================================

type Verbosity int

const (
	Neutral Verbosity = iota
	Critical
	Severe
	Important
	Normal
	Detailed
	Full
)

// ====================================

type ConstraintType int

const (
	FR ConstraintType = iota // unconstrained
	LE                       // less than
	GE                       // greater than
	EQ                       // equal to
)

// ====================================

type ObjectiveType int

const (
	Minimize ObjectiveType = iota
	Maximize
)

// ====================================

type Real float64
