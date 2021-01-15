package control

import (
	"github.com/base"
	"github.com/runtimedata"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *runtimedata.Frame) {
	base.Branch(frame, self.Offset)
}