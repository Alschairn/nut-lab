package classfile

type ConstantClassInfo struct {
	cp        ConstantPool // 常量池
	nameIndex uint16       // 类名常量池索引
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
