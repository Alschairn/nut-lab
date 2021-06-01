package heap

import "jvm-go/classfile"

/**
字段信息
 */
type Field struct {
	ClassMember
}

/**
根据class文件的字段信息创建字段
 */
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}