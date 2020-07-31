package constants

import (
	"instruction"
	"runtimedata"
)

type NOP struct {
	instruction.NoOperandsInstruction
}

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// noting to do
}
