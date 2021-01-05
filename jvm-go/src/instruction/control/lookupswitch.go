package control

import (
	"instruction"
	"github.com/runtimedata"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs int32
	matchOffsets []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *instruction.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *runtimedata.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			instruction.Branch(frame, int(offset))
			return
		}
	}
	instruction.Branch(frame, int(self.defaultOffset))
}