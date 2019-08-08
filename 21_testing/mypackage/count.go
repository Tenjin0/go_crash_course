package mypackage

import "sync/atomic"

var count uint32

func Count() uint32 {

	atomic.AddUint32(&count, 1)
	return count
}
