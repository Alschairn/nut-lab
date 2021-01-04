package constants

import (
	"instruction"
	"github.com/runtimedata"
)

type NOP struct {instruction.NoOperandsInstruction}

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// noting to do
}
