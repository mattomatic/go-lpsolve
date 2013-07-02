package lpsolve

type ResultCode int // result code returned by lpsolve
const (
	OutOfMemory ResultCode = -2
	Error       ResultCode = -1
)
const (
	Optimal ResultCode = iota
	Suboptimal
	InfeasibleModel
	UnboundedModel
	DegenerateModel
	NumericalFailure
	UserAborted
	Timeout
	Presolved
	ProcFail
	ProcBreak
	FeasFound
	NoFeasFound
)

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

type ConstraintType int

const (
	FR ConstraintType = iota // unconstrained
	LE                       // less than
	GE                       // greater than
	EQ                       // equal to
)

type Real float64
