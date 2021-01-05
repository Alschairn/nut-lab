package math

import (
	"instruction"
	"github.com/runtimedata"
)

// todo	实现
type D2F struct{ instruction.NoOperandsInstruction }


type D2I struct{ instruction.NoOperandsInstruction }

func (self *D2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// todo 实现
type D2L struct{ instruction.NoOperandsInstruction }

