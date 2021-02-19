package control

import (
	"base"
	"runtimedata"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *runtimedata.Frame) {
	base.Branch(frame, self.Offset)
}