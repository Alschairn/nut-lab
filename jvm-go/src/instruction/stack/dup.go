package stack

import (
	"github.com/base"
	"github.com/runtimedata"
)

// 复制栈顶单个变量
type DUP struct{ base.NoOperandsInstruction }

func (self *DUP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

type DUP_X1 struct{ base.NoOperandsInstruction }

type DUP_X2 struct{ base.NoOperandsInstruction }

type DUP2 struct{ base.NoOperandsInstruction }

type DUP2_X1 struct{ base.NoOperandsInstruction }

type DUP2_X2 struct{ base.NoOperandsInstruction }

