package comparisons

import (
	"instruction"
	"github.com/runtimedata"
)

type FCMPG struct{ instruction.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *runtimedata.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ instruction.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *runtimedata.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *runtimedata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}