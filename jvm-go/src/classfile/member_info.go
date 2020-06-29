package classfile

/**
	代表字段和方法
 */
type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          // 描述符
	nameIndex       uint16          // 常量池索引，名称
	descriptorIndex uint16          // 常量池索引，描述
	attributes      []AttributeInfo // 属性表
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp: cp,
		accessFlags:reader.readUint16(),
		nameIndex:reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		attributes:readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}