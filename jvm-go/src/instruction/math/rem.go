package math

import (
	"instruction"
	"github.com/runtimedata"
	"math"
)

type DREM struct{ instruction.NoOperandsInstruction }

func (self *DREM) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct{ instruction.NoOperandsInstruction }


type IREM struct{ instruction.NoOperandsInstruction }

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

type LREM struct{ instruction.NoOperandsInstruction }

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
