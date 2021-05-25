package control

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *runtimedata.Frame) {
	base.Branch(frame, self.Offset)
}
