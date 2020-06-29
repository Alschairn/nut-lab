package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex uint16

	/**
	类型描述符
		1. 基本类型byte、short、char、int、long、float和double的描述符是单个字母，分别对应B、S、C、I、J、F和D。
		注意，long的描述符是J而不是L
		2. 引用类型的描述符是L＋类的完全限定名＋分号。
		3. 数组类型的描述符是[＋数组元素类型描述符。
		4. 数组类型的描述符是[＋数组元素类型描述符。
	字段描述符
		1. 数组类型的描述符是[＋数组元素类型描述符。
	方法描述符
		（分号分隔的参数类型描述符）+返回值类型描述符，其中void返回值由单个字母V表示。
	 */
	descriptorIndex uint16 // 描述符
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
