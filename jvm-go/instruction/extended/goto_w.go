package extended

import (
	"base"
	"runtimedata"
)

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *runtimedata.Frame) {
	base.Branch(frame, self.offset)
}