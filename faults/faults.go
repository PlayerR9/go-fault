package faults

import (
	"errors"

	flt "github.com/PlayerR9/go-fault"
	"github.com/PlayerR9/go-fault/faults/internal"
)

var (
	// ErrInvalidFault occurs when an invalid fault is used.
	ErrInvalidFault error
)

func init() {
	ErrInvalidFault = errors.New("invalid fault")
}

// Embeds calls the Embeds method on the given fault and returns the result.
//
// Parameters:
//   - fault: The fault to call the Embeds method on.
//
// Returns:
//   - flt.Fault: The result of the Embeds method. Nil if the fault is nil or if it does not have an Embeds method.
func Embeds(fault flt.Fault) flt.Fault {
	if fault == nil {
		return nil
	}

	inner, ok := internal.Embeds[flt.Fault](fault)
	if !ok {
		return nil
	}

	return inner
}

// Innermost retrieves the innermost embedded element of type flt.Fault from the provided element.
//
// The function continuously calls the Embeds method to access nested elements until it encounters
// an element that either does not have an Embeds method or is equivalent to the zero value of the
// specified type.
//
// If the given element is nil, the function returns nil.
//
// If no innermost element is found, the function panics with an ErrInvalidFault error.
//
// Parameters:
//   - fault: The element from which to extract the innermost embedded element.
//
// Returns:
//   - flt.Fault: The innermost embedded element of type flt.Fault.
func Innermost(fault flt.Fault) flt.Fault {
	if fault == nil {
		return nil
	}

	inner, ok := internal.Innermost[flt.Fault](fault)
	if !ok {
		panic(ErrInvalidFault)
	}

	return inner
}

// TowerOfEmbeds returns the tower of embeds of the given fault where
// the first element is the innermost base of the fault and the last one
// is the given fault itself.
//
// The tower stops at the first nil fault.
//
// If the given fault is nil, the function returns nil.
//
// If no tower is found, the function panics with an ErrInvalidFault error.
//
// Parameters:
//   - fault: The fault.
//
// Returns:
//   - []flt.Fault: The tower of embeds.
func TowerOfEmbeds(fault flt.Fault) []flt.Fault {
	if fault == nil {
		return nil
	}

	tower, ok := internal.TowerOfEmbeds[flt.Fault](fault)
	if !ok {
		panic(ErrInvalidFault)
	}

	return tower
}

// InfoLines returns the info lines of the given fault and any of its embeds; starting from the innermost base.
//
// Parameters:
//   - flt: The fault.
//
// Returns:
//   - []string: The info lines.
func InfoLines(fault flt.Fault) []string {
	if fault == nil {
		return nil
	}

	tower, ok := internal.TowerOfEmbeds[flt.Fault](fault)
	if !ok {
		panic(ErrInvalidFault)
	}

	var lines []string

	for _, elem := range tower {
		e, ok := elem.(interface{ InfoLines() []string })
		if ok {
			sublines := e.InfoLines()
			lines = append(lines, sublines...)
		}
	}

	if len(lines) == 0 {
		return nil
	}

	lines = lines[:len(lines):len(lines)]

	return lines
}

// Is checks whether the given fault is of the same type as the target fault.
//
// If the target fault is nil, Is returns false.
//
// Otherwise, Is returns true if the given fault is of the same type as the target fault,
// i.e. if the given fault's error code is the same as the fault code of the target fault.
//
// If the given fault does not have an error code, Is checks whether the given fault
// is the same as the target fault or whether the given fault has the target fault as its base.
//
// If the given fault has an error code and the target fault does not, Is returns false.
//
// Parameters:
//   - fault: The fault to check.
//   - target: The target fault.
//
// Returns:
//   - bool: True if the given fault is of the same type as the target fault, false otherwise.
func Is(fault flt.Fault, target flt.Fault) bool {
	if fault == nil {
		return false
	}

	for target != nil {
		if fault == target {
			return true
		}

		f, ok := fault.(interface{ IsFault(flt.Fault) bool })
		if ok && f.IsFault(target) {
			return true
		}

		target, ok = internal.Embeds[flt.Fault](target)
		if !ok {
			break
		}
	}

	return false
}

func WithContext(fault flt.Fault, key string, value any) (flt.Fault, bool) {
	if fault == nil {
		return nil, false
	}

	fault, ok := internal.WithContext[flt.Fault](fault, key, value)
	return fault, ok
}
