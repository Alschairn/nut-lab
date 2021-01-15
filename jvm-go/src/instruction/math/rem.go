package math

import (
	"github.com/base"
	"github.com/runtimedata"
	"math"
)

type DREM struct{ base.NoOperandsInstruction }

func (self *DREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct{ base.NoOperandsInstruction }


type IREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}
