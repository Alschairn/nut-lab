package extended

import (
	"base"
	"runtimedata"
)

type IFNULL struct{ base.BranchInstruction } // Branch if reference is null

func (self *IFNULL) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}


type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNONNULL) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
