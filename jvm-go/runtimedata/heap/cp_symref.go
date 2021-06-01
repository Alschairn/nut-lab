package heap

/*
符号引用共性
*/
type SymRef struct {
	cp        *ConstantPool // 放符号引用所在的运行时常量池指针
	className string        // 放类的完全限定名
	class     *Class        // 缓存解析后的类结构体指针
}
