package constants

import (
	base2 "base"
	"runtimedata"
)

type BIPUSH struct {val int8} // 从操作数中获取一个byte型整数，扩展为int型，推入栈顶

func (self *BIPUSH) FetchOperands(reader *base2.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *runtimedata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {val int16} // 从操作数中获取一个short型证书，扩展为int型，推入栈顶

func (self *SIPUSH) FetchOperands(reader *base2.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *runtimedata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}