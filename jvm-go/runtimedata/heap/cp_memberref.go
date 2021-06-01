package heap

import "jvm-go/classfile"

type MemberRef struct {
	SymRef
	name       string // 字段名
	descriptor string // 字段描述符
}

/*
从class文件内存储的字段或方法常量中提取数据
 */
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
