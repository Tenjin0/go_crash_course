package mypackage

import "sync/atomic"

func Count() uint32 {

	var count uint32

	atomic.AddUint32(&count, 1)
	return count
}
