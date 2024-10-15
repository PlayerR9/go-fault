package faults

import (
	faults "github.com/PlayerR9/go-fault"
)

// TowerOfEmbeds returns the tower of embeds of the given fault where
// the first element is the innermost base of the fault and the last one
// is the given fault itself.
//
// The tower stops at the first nil fault.
//
// Parameters:
//   - flt: The fault.
//
// Returns:
//   - []Fault: The tower of embeds.
func TowerOfEmbeds(flt faults.Fault) []faults.Fault {
	var count int

	for flt != nil {
		count++
		flt = flt.Embeds()
	}

	if count == 0 {
		return nil
	}

	tower := make([]faults.Fault, count, count)

	for i := count - 1; i >= 0; i-- {
		tower[i] = flt
		flt = flt.Embeds()
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
func InfoLines(flt faults.Fault) []string {
	tower := TowerOfEmbeds(flt)

	var lines []string

	for _, flt := range tower {
		sub_lines := flt.InfoLines()
		lines = append(lines, sub_lines...)
	}

	return lines
}
