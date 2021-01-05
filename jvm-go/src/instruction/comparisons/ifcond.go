package comparisons

import (
	"instruction"
	"github.com/runtimedata"
)

type IFEQ struct{ instruction.BranchInstruction }

func (self *IFEQ) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		instruction.Branch(frame, self.Offset)
	}
}

type IFNE struct{ instruction.BranchInstruction }
type IFLT struct{ instruction.BranchInstruction }
type IFLE struct{ instruction.BranchInstruction }
type IFGT struct{ instruction.BranchInstruction }
type IFGE struct{ instruction.BranchInstruction }