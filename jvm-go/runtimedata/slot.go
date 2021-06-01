package runtimedata

import "jvm-go/runtimedata/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
