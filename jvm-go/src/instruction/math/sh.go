package math

import (
	"instruction"
	"github.com/runtimedata"
)

type ISHL struct{ instruction.NoOperandsInstruction } // int左位移

func (self *ISHL) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// int变量只有32位，所以只取v2的前5个比特就足够表示位移位数了
	// Go语言位移操作符右侧必须是无符号整数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}


type ISHR struct{ instruction.NoOperandsInstruction } // int算术右位移

func (self *LSHR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

type IUSHR struct{ instruction.NoOperandsInstruction } // int逻辑右位移

func (self *IUSHR) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// Go语言并没有Java语言中的>>>运算符，为了达到无符号位移的目的，需要先把v1转成无符号整数，位移操作之后，再转回有符号整数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}


type LSHL struct{ instruction.NoOperandsInstruction } // long左位移

type LSHR struct{ instruction.NoOperandsInstruction } // long算术右位移

type LUSHR struct{ instruction.NoOperandsInstruction } // long逻辑右位移