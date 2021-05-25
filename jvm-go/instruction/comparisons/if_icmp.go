package comparisons

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

type IF_ICMPEQ struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *runtimedata.Frame) {
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (self *IF_ICMPLT) Execute(frame *runtimedata.Frame) {
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (self *IF_ICMPLE) Execute(frame *runtimedata.Frame) {
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (self *IF_ICMPGT) Execute(frame *runtimedata.Frame) {
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPGE) Execute(frame *runtimedata.Frame) {
}
