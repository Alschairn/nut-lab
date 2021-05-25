package stores

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

// Store float into local variable
type FSTORE struct{ base.Index8Instruction }

func (self *FSTORE) Execute(frame *runtimedata.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (self *FSTORE_0) Execute(frame *runtimedata.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (self *FSTORE_1) Execute(frame *runtimedata.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (self *FSTORE_2) Execute(frame *runtimedata.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (self *FSTORE_3) Execute(frame *runtimedata.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
