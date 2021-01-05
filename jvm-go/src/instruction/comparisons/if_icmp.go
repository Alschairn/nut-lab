package comparisons

import (
	"instruction"
	"github.com/runtimedata"
)

type IF_ICMPEQ struct{ instruction.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		instruction.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct{ instruction.BranchInstruction }
type IF_ICMPLT struct{ instruction.BranchInstruction }
type IF_ICMPLE struct{ instruction.BranchInstruction }
type IF_ICMPGT struct{ instruction.BranchInstruction }
type IF_ICMPGE struct{ instruction.BranchInstruction }