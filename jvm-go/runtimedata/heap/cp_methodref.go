package heap

import "jvm-go/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

/*
创建实例
 */
func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
