package classfile

import "fmt"

type ClassFile struct {
	magic        uint32          // 魔数，标识作用
	minorVersion uint16          // 次版本号
	majorVersion uint16          // 主版本号
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 类访问标识符
	thisClass    uint16          // 常量池索引，类名
	superClass   uint16          // 常量池索引，超类名，除了Object都有
	interfaces   []uint16        // 常量池索引，类实现的所有接口的名称
	fields       []*MemberInfo   // 字段表
	methods      []*MemberInfo   // 方法表
	attributes   []AttributeInfo // todo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)

	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()

	self.constantPool = readConstantPool(reader)
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)

}

/**
	效验class开头是否符合class文件定义
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
	效验class文件版本号
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	/**
		效验主版本号是否符合Java8范围
	 */
	switch self.majorVersion {

	case 45:
		return

	case 46, 47, 48, 49, 50, 51, 52:
		/**
			次版本后java1.2之后基本不维护，判断0即可
		 */
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/**
	从常量池获取当前类名称
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/**
	从常量池获取父类名称
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

/**
	从常量池获取实现的接口集合
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
