package stack

import (
	"instruction"
	"github.com/runtimedata"
)

// Swap the top two operand stack values
type SWAP struct{ instruction.NoOperandsInstruction }

func (self *SWAP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}