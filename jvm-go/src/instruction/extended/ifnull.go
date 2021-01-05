package extended

import (
	"instruction"
	"github.com/runtimedata"
)

type IFNULL struct{ instruction.BranchInstruction } // Branch if reference is null

func (self *IFNULL) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		instruction.Branch(frame, self.Offset)
	}
}


type IFNONNULL struct{ instruction.BranchInstruction } // Branch if reference not null

func (self *IFNONNULL) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		instruction.Branch(frame, self.Offset)
	}
}
