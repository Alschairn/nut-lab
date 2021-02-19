package constants

import (
	"base"
	"runtimedata"
)

type NOP struct { base.NoOperandsInstruction }

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// noting to do
}
