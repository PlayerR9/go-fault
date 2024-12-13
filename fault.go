package faults

// Fault is an interface that represents an error.
type Fault interface {
	// Error returns the error message.
	//
	// Returns:
	//   - string: The error message.
	Error() string

	// InstanceOf creates a new Fault with the given message and who shares the
	// same blueprint.
	//
	// Parameters:
	//   - msg: The message for the Fault.
	//
	// Returns:
	//   - Fault: The new Fault. Never returns nil.
	//
	// Unlike with Init of Blueprint, this function returns a new fault that shares
	// the same pointer to the blueprint as the original. Thus, all faults that
	// are generated with this function are loosely comparable.
	InstanceOf(msg string) Fault
}

// NewFault creates a new Fault with the given name and message.
//
// Parameters:
//   - name: The name of the Fault.
//   - msg: The message for the Fault.
//
// Returns:
//   - Fault: The new Fault. Never returns nil.
func NewFault(name, msg string) Fault {
	blueprint := New(name)
	fault := blueprint.Init(msg)
	return fault
}
