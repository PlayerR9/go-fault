package faults

import "github.com/PlayerR9/go-fault/OLD/internal"

// Fault is the interface for a fault.
type Fault interface {
	// Error returns the message of the fault.
	//
	// Returns:
	//   - string: The message of the fault.
	Error() string

	// Embeds returns the base of the fault.
	//
	// Returns:
	//   - Fault: The base of the fault. Nil if the fault is the base.
	Embeds() Fault

	// Lines returns the info lines of the fault, if any.
	//
	// Returns:
	//   - []string: The info lines of the fault.
	Lines() []string
}

// baseFault is the base implementation of the Fault interface.
type baseFault struct {
	// descriptor is the fault descriptor.
	descriptor Descriptor

	// info is the information of the fault.
	info *internal.Info
}

// Error implements the Fault interface.
func (b baseFault) Error() string {
	return b.descriptor.String()
}

// Embeds implements the Fault interface.
func (b baseFault) Embeds() Fault {
	return nil
}

// Lines implements the Fault interface.
func (b baseFault) Lines() []string {
	return b.info.Lines()
}

// AddContext adds a key-value pair to the fault's context.
//
// Parameters:
//   - key: The key for the context entry.
//   - value: The value associated with the key.
func (b *baseFault) AddContext(key string, value any) {
	if b == nil {
		panic(internal.ErrNilReceiver)
	}

	b.info.AddContext(key, value)
}

// AppendFrame appends a frame to the fault's stack trace.
//
// Parameters:
//   - frame: The frame to append.
func (b *baseFault) AppendFrame(frame string) {
	if b == nil {
		panic(internal.ErrNilReceiver)
	}

	b.info.AppendFrame(frame)
}

// Get returns the value associated with the given key in the fault's context.
//
// Returns:
//   - any: The value associated with the key.
//   - bool: True if the key is in the context, false otherwise.
func (b *baseFault) Get(key string) (any, bool) {
	if b == nil {
		panic(internal.ErrNilReceiver)
	}

	v, ok := b.info.Get(key)
	return v, ok
}
