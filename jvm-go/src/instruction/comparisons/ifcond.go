package comparisons

import (
	"github.com/base"
	"github.com/runtimedata"
)

type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }
type IFLT struct{ base.BranchInstruction }
type IFLE struct{ base.BranchInstruction }
type IFGT struct{ base.BranchInstruction }
type IFGE struct{ base.BranchInstruction }