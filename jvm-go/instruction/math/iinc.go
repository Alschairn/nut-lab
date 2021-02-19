package math

import (
	"jvm-go/instruction/base"
	"jvm-go/runtimedata"
)

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

// 从字节码读取操作数
func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

// 从局部变量表读取变量，加上常量值，协会局部变量表
func (self *IINC) Execute(frame *runtimedata.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}