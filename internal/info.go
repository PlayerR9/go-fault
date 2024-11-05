package internal

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

// Info is the information of a fault.
type Info struct {
	// context is the context of the fault.
	context map[string]any

	// frames is the stack frame of the fault.
	frames []string

	// timestamp is the timestamp of the fault.
	timestamp time.Time
}

// NewInfo creates a new Info with the current timestamp and empty context.
//
// Returns:
//   - *Info: The new Info. Never returns nil.
func NewInfo() *Info {
	return &Info{
		timestamp: time.Now(),
	}
}

// Get returns the value associated with the given key in the context of the fault.
//
// Returns:
//   - any: The value associated with the key.
//   - bool: True if the key is in the context, false otherwise.
func (f Info) Get(key string) (any, bool) {
	if len(f.context) == 0 {
		return nil, false
	}

	return f.context[key], true
}

// AddContext adds a key-value pair to the context of the fault.
//
// Parameters:
//   - key: The key for the context entry.
//   - value: The value associated with the key.
//
// Panics if the receiver is nil.
func (f *Info) AddContext(key string, value any) {
	if f == nil {
		panic(ErrNilReceiver)
	}

	if f.context == nil {
		f.context = make(map[string]any)
	}

	f.context[key] = value
}

// AppendFrame appends a frame to the stack trace of the fault.
//
// Parameters:
//   - frame: The frame to append.
//
// Panics if the receiver is nil.
func (f *Info) AppendFrame(frame string) {
	if f == nil {
		panic(ErrNilReceiver)
	}

	f.frames = append(f.frames, frame)
}

// Lines returns the info lines of the fault.
//
// The returned lines are in the following format:
//
// Timestamp: <timestamp>
//
// Stack Trace:
//   - <frame> <- <frame> <- ...
//
// Context:
//   - <key>: <value>
func (f Info) Lines() []string {
	var lines []string

	lines = append(lines, fmt.Sprintf("Timestamp: %s", f.timestamp.Format(time.RFC3339)))

	if len(f.frames) > 0 {
		lines = append(lines, "Stack Trace:")

		frames := make([]string, len(f.frames))
		copy(frames, f.frames)
		slices.Reverse(frames)

		lines = append(lines, "- "+strings.Join(frames, " <- "))
	}

	if len(f.context) > 0 {
		lines = append(lines, "Context:")

		for key, value := range f.context {
			lines = append(lines, fmt.Sprintf("- %s: %v", key, value))
		}
	}

	return lines
}
