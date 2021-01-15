package constants

import (
	"github.com/base"
	"github.com/runtimedata"
)

// 将null引用推入操作数栈顶
type ACONST_NULL struct {base.NoOperandsInstruction}

func (self *ACONST_NULL) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 将double型0推入操作数栈顶
type DCONST_0 struct {base.NoOperandsInstruction}

func (self *DCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// 将double型1推入操作数栈顶
type DCONST_1 struct {base.NoOperandsInstruction}

func (self *DCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// 将float型0推入操作数栈顶
type FCONST_0 struct {base.NoOperandsInstruction}

func (self *FCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// 将float型1推入操作数栈顶
type FCONST_1 struct {base.NoOperandsInstruction}

func (self *FCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// 将float型2推入操作数栈顶
type FCONST_2 struct {base.NoOperandsInstruction}

func (self *FCONST_2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// 将int型-1推入操作数栈顶
type ICONST_M1 struct {base.NoOperandsInstruction}

func (self *ICONST_M1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(-1)
}

// 将int型0推入操作数栈顶
type ICONST_0 struct {base.NoOperandsInstruction}

func (self *ICONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(0)
}

// 将int型1推入操作数栈顶
type ICONST_1 struct {base.NoOperandsInstruction}

func (self *ICONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(1)
}

// 将int型2推入操作数栈顶
type ICONST_2 struct {base.NoOperandsInstruction}

func (self *ICONST_2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(2)
}

// 将int型5推入操作数栈顶
type ICONST_3 struct {base.NoOperandsInstruction}

func (self *ICONST_3) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(3)
}

// 将int型5推入操作数栈顶
type ICONST_4 struct {base.NoOperandsInstruction}

func (self *ICONST_4) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(4)
}

// 将int型5推入操作数栈顶
type ICONST_5 struct {base.NoOperandsInstruction}

func (self *ICONST_5) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(5)
}

// 将long型0推入操作数栈顶
type LCONST_0 struct {base.NoOperandsInstruction}

func (self *LCONST_0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(0)
}

// 将long型1推入操作数栈顶
type LCONST_1 struct {base.NoOperandsInstruction}

func (self *LCONST_1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(1)
}
