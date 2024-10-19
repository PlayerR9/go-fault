// Code generated by "stringer -type=StdFaultCode -linecomment"; DO NOT EDIT.

package faults

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown - -1]
	_ = x[FatalErr-0]
	_ = x[OperationFail-1]
	_ = x[BadParameter-2]
}

const _StdFaultCode_name = "UNKNOWNFATALOperation FailedBad Parameter"

var _StdFaultCode_index = [...]uint8{0, 7, 12, 28, 41}

func (i StdFaultCode) String() string {
	i -= -1
	if i < 0 || i >= StdFaultCode(len(_StdFaultCode_index)-1) {
		return "StdFaultCode(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _StdFaultCode_name[_StdFaultCode_index[i]:_StdFaultCode_index[i+1]]
}
