package heap

import "jvm-go/classfile"

/**
字段和方法都属于类的成员，它们有一些相同的信息（访问标志、名字、描述符）
避免重复，创建如下结构体存放
 */
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

/**
从class文件复制数据
 */
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}


