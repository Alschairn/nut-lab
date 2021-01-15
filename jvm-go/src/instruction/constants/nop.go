package constants

import (
	"github.com/base"
	"github.com/runtimedata"
)

type NOP struct {base.NoOperandsInstruction}

func (self *NOP) Execute(frame *runtimedata.Frame) {
	// noting to do
}
