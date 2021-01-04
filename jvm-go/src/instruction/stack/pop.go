package stack

import (
	"instruction"
	"github.com/runtimedata"
)

// 栈顶变量弹出，用于弹出占用一个操作数栈位置的变量
type POP struct{ instruction.NoOperandsInstruction }

func (self *POP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// 栈顶变量弹出，用于弹出占用两个个操作数栈位置的变量
type POP2 struct{ instruction.NoOperandsInstruction }

func (self *POP2) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}