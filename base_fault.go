package faults

// baseFault is a base implementation of the Fault interface.
type baseFault struct {
	// blueprint is the blueprint of the Fault.
	blueprint Blueprint

	// msg is the error message.
	msg string
}

// Error implements Fault.
func (f baseFault) Error() string {
	return f.msg
}

// InstanceOf implements Fault.
func (f baseFault) InstanceOf(msg string) Fault {
	return &baseFault{
		blueprint: f.blueprint,
		msg:       msg,
	}
}

// Blueprint returns the blueprint of the fault.
//
// Returns:
//   - Blueprint: The blueprint of the fault. Never returns nil.
func (f baseFault) Blueprint() Blueprint {
	return f.blueprint
}
