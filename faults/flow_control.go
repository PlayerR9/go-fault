package faults

import (
	flt "github.com/PlayerR9/go-fault"
)

func try(fault *flt.Fault, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case flt.Fault:
			*fault = r
		case string:
			*fault = FromString(r)
		case error:
			*fault = FromError(r)
		default:

		}
	}()

	fn()
}

func Try(fn func()) flt.Fault {
	if fn == nil {
		return nil
	}

	var flt flt.Fault

	try(&flt, fn)

	return flt
}
