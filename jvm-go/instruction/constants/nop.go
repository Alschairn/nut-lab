package constants

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// really do nothing
}
