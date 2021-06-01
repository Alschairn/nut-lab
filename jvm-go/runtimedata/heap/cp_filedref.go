package heap

import "jvm-go/classfile"

type FieldRef struct {
	MemberRef
	field *Field // 解析后的字段指针
}

/*
创建实例
 */
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

