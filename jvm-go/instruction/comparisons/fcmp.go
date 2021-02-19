package comparisons

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *runtimedata.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

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