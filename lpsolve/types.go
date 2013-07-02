package lpsolve 

type ResultCode int // result code returned by lpsolve
const (
    OutOfMemory ResultCode = -2
    Error ResultCode = -1
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

type Constraint int
const (
    LE Constraint = iota
    GE
    EQ
)