package comparisons

import (
	"instruction"
	"github.com/runtimedata"
)

type IF_ACMPEQ struct{ instruction.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		instruction.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ instruction.BranchInstruction }


