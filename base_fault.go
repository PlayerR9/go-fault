package faults

import "fmt"

// baseFault is a base class for fault implementations.
type baseFault[C FaultCode] struct {
	// code is the fault code.
	code C

	// msg is the error message.
	msg string

	// context is the fault context.
	context map[string]any
}

// Error implements the Fault interface.
func (f baseFault[C]) Error() string {
	return "(" + f.code.String() + ") " + f.msg
}

// Embeds implements the Fault interface.
func (f baseFault[C]) Embeds() Fault {
	return nil
}

// IsFault checks whether the given fault is of the same type as the target fault.
//
// More specifically, IsFault returns true if the given fault's error code is the same as the
// fault code of the target fault.
func (f baseFault[C]) IsFault(target Fault) bool {
	for target != nil {
		base, ok := target.(*baseFault[C])
		if ok {
			return f.code == base.code
		}

		target = target.Embeds()
	}

	return false
}

// InfoLines returns the info lines of the fault (i.e., any other information that is
// not conveyed by the Error() method).
//
// Returns:
//   - []string: The info lines.
func (bf baseFault[C]) InfoLines() []string {
	var lines []string

	if len(bf.context) > 0 {
		lines = append(lines, "Context:")

		for k, v := range bf.context {
			lines = append(lines, fmt.Sprintf("- %s: %v", k, v))
		}
	}

	return lines
}

// New creates a new Fault with the given fault code and message.
//
// Parameters:
//   - code: The fault code.
//   - msg: The error message.
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func New[C FaultCode](code C, msg string) Fault {
	return &baseFault[C]{
		code: code,
		msg:  msg,
	}
}

// Newf is the same as New, but with a format string.
//
// Parameters:
//   - code: The fault code.
//   - format: The format string.
//   - args: The arguments for the format string.
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func Newf[C FaultCode](code C, format string, args ...any) Fault {
	msg := fmt.Sprintf(format, args...)

	return &baseFault[C]{
		code: code,
		msg:  msg,
	}
}
