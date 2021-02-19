package constants

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

type NOP struct { base.NoOperandsInstruction }

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// noting to do
}
