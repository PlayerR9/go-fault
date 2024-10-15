package faults

type StdFaultCode int

const (
	Unknown StdFaultCode = iota - 1
	FatalErr
)

func FromString(s string) Fault {
	if s == "" {
		return nil
	}

	return &baseFault[StdFaultCode]{
		code: Unknown,
		msg:  s,
	}
}

func FromError(err error) Fault {
	if err == nil {
		return nil
	}

	return &baseFault[StdFaultCode]{
		code: Unknown,
		msg:  err.Error(),
	}
}

func NewErrPanic(value any) Fault {
	fault := New(FatalErr, "A panic has occurred")

	return fault
}
