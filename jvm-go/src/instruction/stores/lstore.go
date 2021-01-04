package stores

import (
	"instruction"
	"github.com/runtimedata"
)

// Store long into local variable
type LSTORE struct{ instruction.Index8Instruction }

func (self *LSTORE) Execute(frame *runtimedata.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_0) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_1) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_2) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_3) Execute(frame *runtimedata.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}