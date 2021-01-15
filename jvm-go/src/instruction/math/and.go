package math

import (
	"github.com/base"
	"github.com/runtimedata"
)

type IAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

type LAND struct{ base.NoOperandsInstruction }

func (self *LAND) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
