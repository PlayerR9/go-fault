package internal

import "errors"

var (
	// ErrNilReceiver occurs when a receiver is nil.
	//
	// Format:
	// 	"receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrNilReceiver = errors.New("receiver must not be nil")
}
