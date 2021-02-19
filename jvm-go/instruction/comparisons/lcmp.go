package comparisons

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

// 比较Long变量
type LCMP struct { base.NoOperandsInstruction }

// 将栈顶的两个long变量弹出，进行比较，将比较结果（0、1、-1）推入栈顶
func (self *LCMP) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}