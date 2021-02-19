package runtimedata

type Frame struct {
	lower        *Frame        // 链表
	localVars    LocalVars     // 局部变量表指针
	operandStack *OperandStack // 操作数栈指针
	thread       *Thread
	nextPC       int
}


func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) Thread() *Thread {
	return self.thread
}