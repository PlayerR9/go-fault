package faults

import "slices"

// faultT is for private use.
type faultT struct{}

// Faults is the namespace for operating on faults.
var Faults faultT

func init() {
	Faults = faultT{}
}

// AddContext adds a key-value pair to the context of the given fault.
//
// Parameters:
//   - fault: The fault to add the context to.
//   - key: The key for the context entry.
//   - value: The value associated with the key.
//
// Panics if the fault is nil or if no AddContext() method is found in the fault's tower of embeds.
func (faultT) AddContext(fault Fault, key string, value any) {
	if fault == nil {
		panic("no fault to add context to")
	}

	var f interface{ AddContext(key string, value any) }
	var ok bool

	for {
		f, ok = fault.(interface{ AddContext(key string, value any) })
		if ok {
			break
		}

		fault = fault.Embeds()
		if fault == nil {
			panic("no AddContext() method found in fault's tower of embeds")
		}
	}

	f.AddContext(key, value)
}

// Throw appends a frame to the stack trace of the given fault.
//
// Parameters:
//   - fault: The fault to which the frame will be appended.
//   - frame: The frame to append to the fault's stack trace.
//
// Returns:
//   - Fault: The fault with the appended frame. Nil if the fault was nil.
//
// Panics if the fault is nil or if no AppendFrame() method is found in the fault's tower of embeds.
func (faultT) Throw(fault Fault, frame string) Fault {
	if fault == nil {
		return nil
	}

	target := fault

	var f interface{ AppendFrame(frame string) }
	var ok bool

	for {
		f, ok = target.(interface{ AppendFrame(frame string) })
		if ok {
			break
		}

		target = target.Embeds()
		if target == nil {
			panic("no AppendFrame() method found in fault's tower of embeds")
		}
	}

	f.AppendFrame(frame)

	return fault
}

// TowerOfEmbeds returns the tower of embeds of the given fault.
//
// Parameters:
//   - fault: The fault to return the tower of embeds for.
//
// Returns:
//   - []Fault: The tower of embeds of the fault. Nil if the fault was nil.
//
// The returned list is in the order of the fault being passed in last, and the
// base of the fault first.
func (faultT) TowerOfEmbeds(fault Fault) []Fault {
	if fault == nil {
		return nil
	}

	tower := []Fault{fault}

	for {
		fault = fault.Embeds()
		if fault == nil {
			break
		}

		tower = append(tower, fault)
	}

	slices.Reverse(tower)

	return tower
}

// InfoLines returns the info lines of the given fault, by traversing the
// fault's tower of embeds and concatenating the info lines of each fault in
// the tower.
//
// Parameters:
//   - fault: The fault to return the info lines for.
//
// Returns:
//   - []string: The info lines of the fault. Nil if the fault was nil.
func (faultT) InfoLines(fault Fault) []string {
	if fault == nil {
		return nil
	}

	tower := Faults.TowerOfEmbeds(fault)

	var lines []string

	for _, f := range tower {
		lines = append(lines, f.Lines()...)
	}

	return lines
}

// try recovers from panics and stores the recovered value in the given fault pointer.
//
// Parameters:
//   - fault: A pointer to a fault to store the recovered value in.
//   - fn: The function to run.
func try(fault *Fault, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case string:
			*fault = New(Unknown, r).New()
		case Fault:
			*fault = r
		case error:
			*fault = FromErr(Unknown, r).New()
		default:
			*fault = NewErrPanic(r)
		}
	}()

	fn()
}

// Try runs the given function and returns the fault that was recovered, or nil if
// no fault was recovered.
//
// Parameters:
//   - fn: The function to run.
//
// Returns:
//   - Fault: The fault that was recovered, or nil if none.
func (faultT) Try(fn func()) Fault {
	if fn == nil {
		return nil
	}

	var fault Fault

	try(&fault, fn)

	return fault
}

// Get returns the value associated with the given key in the context of the given
// fault. If the fault is nil, or if the key is not in the context, then the second
// return value is false, and the first return value is undefined.
//
// Returns:
//   - any: The value associated with the key.
//   - bool: True if the key is in the context, false otherwise.
func (faultT) Get(fault Fault, key string) (any, bool) {
	if fault == nil {
		return nil, false
	}

	var f interface{ Get(key string) (any, bool) }
	var ok bool

	for {
		f, ok = fault.(interface{ Get(key string) (any, bool) })
		if ok {
			break
		}

		fault = fault.Embeds()
		if fault == nil {
			return nil, false
		}
	}

	v, ok := f.Get(key)
	return v, ok
}

// Get returns the value associated with the given key in the context of the given
// fault, or a zero value of type T if the key is not in the context. If the fault
// is nil, or if the key is not in the context, then the second return value is
// nil. Otherwise, the second return value is the fault.
//
// Returns:
//   - T: The value associated with the key, or a zero value of type T if the key
//     is not in the context.
//   - Fault: The fault, or nil if the fault is nil or if the key is not in the
//     context.
func Get[T any](fault Fault, key string) (T, Fault) {
	if fault == nil {
		return *new(T), NewErrNoSuchKey(key)
	}

	a, ok := Faults.Get(fault, key)
	if !ok {
		return *new(T), NewErrNoSuchKey(key)
	}

	v, ok := a.(T)
	if !ok {
		return *new(T), NewErrWrongKey(key, *new(T), a)
	}

	return v, fault
}
