package faults

import "fmt"

// Enum is an interface that must be implemented by all enums.
type Enum interface {
	~int

	fmt.Stringer
}

// StdCode is the type of a standard error code.
type StdCode int

const (
	// Unknown is the standard error code for an unknown error.
	Unknown StdCode = iota // Unknown

	// FatalError is the standard error code for a fatal error.
	FatalError // Fatal Error

	// OperationFail occurs when a function fails.
	OperationFail // Operation Failed
)

var (
	// ErrPanic is the descriptor for a panic error.
	ErrPanic Descriptor

	// ErrNoSuchKey is the descriptor for a no such key error.
	ErrNoSuchKey Descriptor

	// ErrWrongKey is the descriptor for a wrong key error.
	ErrWrongKey Descriptor
)

func init() {
	ErrPanic = New(FatalError, "a panic has occurred")

	ErrNoSuchKey = New(OperationFail, "no such key")

	ErrWrongKey = New(OperationFail, "key refers to a value of a different type")
}

// NewErrPanic creates a new Fault indicating a panic.
//
// Parameters:
//   - value: The value of the panic.
//
// Returns:
//   - Fault: A Fault with context indicating a panic.
func NewErrPanic(value any) Fault {
	fault := ErrPanic.New()

	Faults.AddContext(fault, "Value", value)

	return fault
}

// NewErrNoSuchKey creates a new Fault indicating that a specified key does not exist.
//
// Parameters:
//   - key: The key that was not found.
//
// Returns:
//   - Fault: A Fault with context indicating the missing key.
func NewErrNoSuchKey(key string) Fault {
	fault := ErrNoSuchKey.New()

	Faults.AddContext(fault, "Key", key)

	return fault
}

// NewErrWrongKey creates a new Fault indicating that a specified key refers to a value
// of a different type.
//
// Parameters:
//   - key: The key that was not found.
//   - expectedType: The expected type of the value.
//   - actualType: The actual type of the value.
//
// Returns:
//   - Fault: A Fault with context indicating the wrong key.
func NewErrWrongKey(key string, expectedType, actualType any) Fault {
	fault := ErrWrongKey.New()

	Faults.AddContext(fault, "Key", key)

	if expectedType != nil {
		Faults.AddContext(fault, "Expected Type", fmt.Sprintf("%T", expectedType))
	} else {
		Faults.AddContext(fault, "Expected Type", "nil")
	}

	if actualType != nil {
		Faults.AddContext(fault, "Actual Type", fmt.Sprintf("%T", actualType))
	} else {
		Faults.AddContext(fault, "Actual Type", "nil")
	}

	return fault
}
