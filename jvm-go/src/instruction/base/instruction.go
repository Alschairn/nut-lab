package base

import (
	"github.com/runtimedata"
)


/**
	指令接口
 */
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *runtimedata.Frame)
}


type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 跳转指令
type BranchInstruction struct {
	Offset int // 跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 存取局部变量表
type Index8Instruction struct {
	Index uint // 局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 访问运行时常量池
type Index16Instruction struct {
	Index uint // 常量池索引
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

