package classfile

type ConstantStringInfo struct {
	cp          ConstantPool // 常量池
	stringIndex uint16       // 字符串索引
}

/**
	获取常量池索引
 */
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

/**
	根据索引从常量池获取字符串
 */
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
