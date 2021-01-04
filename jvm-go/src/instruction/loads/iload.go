package loads

import (
	"instruction"
	"github.com/runtimedata"
)

func _iload(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

// Load int from local variable
type ILOAD struct{ instruction.Index8Instruction }

func (self *ILOAD) Execute(frame *runtimedata.Frame) {
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *ILOAD_0) Execute(frame *runtimedata.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *ILOAD_1) Execute(frame *runtimedata.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *ILOAD_2) Execute(frame *runtimedata.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *ILOAD_3) Execute(frame *runtimedata.Frame) {
	_iload(frame, 3)
}
