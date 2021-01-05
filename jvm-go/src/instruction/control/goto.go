package control

import (
	"instruction"
	"github.com/runtimedata"
)

type GOTO struct{ instruction.BranchInstruction }

func (self *GOTO) Execute(frame *runtimedata.Frame) {
	instruction.Branch(frame, self.Offset)
}