package fltif

import (
	flt "github.com/PlayerR9/go-fault"
	flts "github.com/PlayerR9/go-fault/faults"
)

// NewAssertFail creates a new fault with the code AssertFail and the message
// given by the given string.
//
// Parameters:
//   - msg: The message.
//
// Returns:
//   - Fault: The new fault. Never returns nil.
func NewAssertFail(msg string) flt.Fault {
	return flt.New(flts.AssertFail, msg)
}
