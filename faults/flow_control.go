package faults

import (
	"fmt"

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
			*fault = flt.New(Unknown, r)
		case error:
			*fault = flt.New(Unknown, r.Error())
		default:
			*fault = flt.New(Unknown, fmt.Sprintf("panic: %v", r))
		}
	}()

	fn()
}

func Try(fn func()) flt.Fault {
	if fn == nil {
		return nil
	}

	var fault flt.Fault

	try(&fault, fn)

	return fault
}
