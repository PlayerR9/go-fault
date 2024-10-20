package fltif

import (
	"fmt"

	flt "github.com/PlayerR9/go-fault"
)

// Cond asserts whether the condition is true. If not, it returns a fault with an
// ErrAssertFailed error.
//
// Parameters:
//   - cond: The condition to assert.
//   - msg: The message of the error. If empty, "something went wrong" is used.
//
// Example:
//
//	Cond("foo" == "bar", "foo must be bar") // Panics: foo must be bar
func Cond(cond bool, msg string) flt.Fault {
	if cond {
		return nil
	}

	return NewAssertFail(msg)
}

// Condf is the same as Cond, but with a format string.
//
// Parameters:
//   - cond: The condition to assert.
//   - format: The format.
//   - args: The arguments of the format.
//
// Example:
//
//	Condf("foo" = "bar", "%q must be %q", "foo", "bar") // Panics: "foo" must be "bar"
func Condf(cond bool, format string, args ...any) flt.Fault {
	if cond {
		return nil
	}

	return NewAssertFail(fmt.Sprintf(format, args...))
}

// Err asserts whether the error is nil. If not, it panics with an
// ErrAssertFailed error.
//
// Parameters:
//   - inner: The error to assert.
//   - format: The format of the function call that returned the error.
//   - args: The arguments of the function call.
//
// Example:
//
//	func MyFunc(a, b string) error {
//		return errors.New("something went wrong")
//	}
//
//	err := MyFunc("foo", "bar")
//	Err(err, "MyFunc(%q, %q)", "foo", "bar") // Panics: MyFunc("foo", "bar") = something went wrong
func Err(inner error, format string, args ...any) flt.Fault {
	if inner == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = " + inner.Error()

	return NewAssertFail(msg)
}

// True is just syntactic sugar for the Condf function but used for when functions
// return a boolean result to indicate success instead of an error.
//
// Parameters:
//   - ok: The boolean to assert.
//   - format: The format of the function call that returned the boolean.
//   - args: The arguments of the function call.
//
// Example:
//
//	func MyFunction(a, b string) bool {
//		return a == b
//	}
//
//	ok := MyFunction("foo", "bar")
//	True(ok, "MyFunction(%q, %q)", "foo", "bar") // Panics: MyFunction("foo", "bar") = false
func True(ok bool, format string, args ...any) flt.Fault {
	if ok {
		return nil
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = false"

	return NewAssertFail(msg)
}

// False is the same as True, but checks for false instead of true.
//
// Parameters:
//   - ok: The boolean to assert.
//   - format: The format of the function call that returned the boolean.
//   - args: The arguments of the function call.
//
// Example:
//
//	func MyFunction(a, b string) bool {
//		return a != b
//	}
//
//	ok := MyFunction("foo", "bar")
//	False(ok, "MyFunction(%q, %q)", "foo", "bar") // Panics: MyFunction("foo", "bar") = true
func False(ok bool, format string, args ...any) flt.Fault {
	if !ok {
		return nil
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = true"

	return NewAssertFail(msg)
}

// NotNil asserts whether the variable is non-nil according to the IsNiler interface. If
// not, it panics with an ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//
// Example:
//
//	type MyStruct struct {}
//
//	func (ms *MyStruct) IsNil() bool {
//		return ms == nil
//	}
//
//	var ms *MyStruct
//	NotNil(ms, "ms") // Panics: ms = nil
func NotNil(v interface {
	IsNil() bool
}, name string) flt.Fault {
	if v != nil && !v.IsNil() {
		return nil
	}

	if name == "" {
		name = "variable"
	}

	return NewAssertFail(name + " = nil")
}

// NotZero asserts whether the variable is not its zero value. If not, it
// panics with an ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, "variable" is used.
//
// Example:
//
//	v := 0
//
//	NotZero[int](v, "v") // Panics: v = 0
func NotZero[T comparable](v T, name string) flt.Fault {
	zero := *new(T)

	if v != zero {
		return nil
	}

	if name == "" {
		name = "variable"
	}

	msg := fmt.Sprintf("%s = %v", name, v)

	return NewAssertFail(msg)
}

// Type asserts whether the variable is of type T. If not, it panics with an
// ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//   - allow_nil: Whether to allow the variable to be nil.
//
// Example:
//
//	v := "foo"
//	Type[int](v, "v", false) // Panics: v = string, want int
func Type[T any](v any, name string, allow_nil bool) flt.Fault {
	if name == "" {
		name = "variable"
	}

	if v == nil && !allow_nil {
		return NewAssertFail(name + " = nil")
	} else if v == nil {
		return nil
	}

	_, ok := v.(T)
	if !ok {
		msg := fmt.Sprintf("%s = %T, want %T", name, v, *new(T))

		return NewAssertFail(msg)
	}

	return nil
}

// Deref asserts whether the variable is both non-nil and is of type T. If
// not, it panics with an ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//
// Returns:
//   - T: The dereferenced variable.
//
// Example:
//
//	var v *int
//	_ = Deref[string](v, "v") // Panics: v = *int, want string
func Deref[T any](v any, name string) (T, flt.Fault) {
	if name == "" {
		name = "variable"
	}

	if v == nil {
		msg := fmt.Sprintf("%s = nil, want %T", name, *new(T))

		return *new(T), NewAssertFail(msg)
	}

	switch v := v.(type) {
	case *T:
		return *v, nil
	case T:
		return v, nil
	default:
		msg := fmt.Sprintf("%s = %T, want %T", name, v, *new(T))

		return *new(T), NewAssertFail(msg)
	}
}

// New is a syntactic sugar asserting constructors. It asserts whether the
// constructor does not return an error and the result is non-nil. If not, it
// panics with an ErrAssertFailed error.
//
// Parameters:
//   - res: The result of the constructor.
//   - inner: The error returned by the constructor.
//
// Example:
//
//	type MyStruct struct {}
//
//	func (my_struct *MyStruct) IsNil() bool {
//		return my_struct == nil
//	}
//
//	func NewMyStruct() (*MyStruct, error) {
//		return nil, nil
//	}
//
//	res := New(NewMyStruct()) // Panics: *MyStruct = nil
func New[T interface{ IsNil() bool }](res T, err error) (T, flt.Fault) {
	if err != nil {
		return *new(T), NewAssertFail("err = " + err.Error())
	}

	if res.IsNil() {
		msg := fmt.Sprintf("%T = nil", *new(T))

		return *new(T), NewAssertFail(msg)
	}

	return res, nil
}

// Conv asserts whether the variable is of type T. If not, it panics with an
// ErrAssertFailed error. Unlike with Type(), this returns the converted
// type as well.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//
// Example:
//
//	v := "foo"
//	res := Conv[int](v, "v") // Panics: v = string, want int
func Conv[T any](v any, name string) (T, flt.Fault) {
	if name == "" {
		name = "variable"
	}

	if v == nil {
		return *new(T), NewAssertFail(name + " = nil")
	}

	val, ok := v.(T)
	if !ok {
		msg := fmt.Sprintf("%s = %T, want %T", name, v, *new(T))

		return *new(T), NewAssertFail(msg)
	}

	return val, nil
}
