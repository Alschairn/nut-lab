package extended

import (
	"instruction"
	"github.com/runtimedata"
)

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *instruction.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *runtimedata.Frame) {
	instruction.Branch(frame, self.offset)
}