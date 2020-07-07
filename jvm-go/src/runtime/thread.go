package runtime

type Thread struct {
	pc    int
	stack *Stack // java虚拟机栈指针
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}


func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/**
	返回方法当前帧
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
