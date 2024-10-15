package faults

func try(flt *Fault, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case Fault:
			*flt = r
		case string:
			*flt = FromString(r)
		case error:
			*flt = FromError(r)
		default:

		}
	}()

	fn()
}

func Try(fn func()) Fault {
	if fn == nil {
		return nil
	}

	var flt Fault

	try(&flt, fn)

	return flt
}
