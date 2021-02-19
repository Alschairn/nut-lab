package classfile

/**
	 用于标识类、接口、字段、方法已不建议使用
 */
type DeprecatedAttribute struct {
	MarkerAttribute
}

/**
	用于标记源文件中不存在，由编译器生成的类成员
 */
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

