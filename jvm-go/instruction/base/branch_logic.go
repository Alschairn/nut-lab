package base

import "jvm-go/runtimedata"

func Branch(frame *runtimedata.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}