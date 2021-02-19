package math

import (
	"base"
	"runtimedata"
)

// todo	实现
type D2F struct{ base.NoOperandsInstruction }


type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// todo 实现
type D2L struct{ base.NoOperandsInstruction }

