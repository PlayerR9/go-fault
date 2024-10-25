package fault

import (
	"fmt"
	"strings"
)

type FaultCode interface {
	~int

	String() string
}

type Fault interface {
	Error() string
}

type baseFault[C FaultCode] struct {
	code    C
	msg     string
	context map[string]any
}

func (f baseFault[C]) Error() string {
	return "(" + f.code.String() + ") " + f.msg
}

func New[C FaultCode](code C, msg string) Fault {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		msg = "something went wrong"
	}

	return &baseFault[C]{
		code: code,
		msg:  msg,
	}
}

func (f baseFault[C]) Embeds() Fault {
	return nil
}

func (f baseFault[C]) IsFault(target Fault) bool {
	if target == nil {
		return false
	}

	b, ok := target.(*baseFault[C])
	if !ok {
		return false
	}

	return b.code == f.code && strings.EqualFold(f.msg, b.msg)
}

func (f baseFault[C]) InfoLines() []string {
	var lines []string

	if len(f.context) > 0 {
		lines = append(lines, "Context:")

		for k, v := range f.context {
			lines = append(lines, fmt.Sprintf("- %s: %v", k, v))
		}
	}

	return lines
}

func (f *baseFault[C]) WithContext(k string, v any) Fault {
	if f.context == nil {
		f.context = make(map[string]any)
	}

	f.context[k] = v

	return f
}
