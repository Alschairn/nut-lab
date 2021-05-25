package comparisons

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *runtimedata.Frame) {
}

type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *runtimedata.Frame) {
}

type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *runtimedata.Frame) {
}

type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *runtimedata.Frame) {
}

type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *runtimedata.Frame) {
}
