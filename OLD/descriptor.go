package faults

import (
	"fmt"

	"github.com/PlayerR9/go-fault/OLD/internal"
)

// Descriptor is the interface for a fault descriptor.
type Descriptor interface {
	// New creates a new instance of the fault.
	//
	// Returns:
	//   - Fault: The new fault. Never returns nil.
	New() Fault

	fmt.Stringer
}

// faultDescriptor is the descriptor of a fault.
type faultDescriptor[T Enum] struct {
	// code is the code of the fault.
	code T

	// msg is the message of the fault.
	msg string
}

// New implements the Descriptor interface.
func (d faultDescriptor[T]) New() Fault {
	return &baseFault{
		descriptor: d,
		info:       internal.NewInfo(),
	}
}

// String implements the fmt.Stringer interface.
func (d faultDescriptor[T]) String() string {
	return "(" + d.code.String() + ") " + d.msg
}

// New creates a new Fault with the given message.
//
// Parameters:
//   - code: The code of the fault.
//   - msg: The message of the fault.
//
// Returns:
//   - Descriptor: The new fault. Never returns nil.
func New[C Enum](code C, msg string) Descriptor {
	return &faultDescriptor[C]{
		code: code,
		msg:  msg,
	}
}

// FromErr creates a new Fault with the given code and the given error.
// If the error is nil, "something went wrong" is used.
//
// Parameters:
//   - code: The code of the fault.
//   - err: The error to create the fault from.
//
// Returns:
//   - Descriptor: The new fault.
func FromErr[C Enum](code C, err error) Descriptor {
	var msg string

	if err == nil {
		msg = "something went wrong"
	} else {
		msg = err.Error()
	}

	return &faultDescriptor[C]{
		code: code,
		msg:  msg,
	}
}
