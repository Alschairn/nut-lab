package heap

import (
	"fmt"
	"jvm-go/classfile"
	"jvm-go/classpath"
)

/*
类加载器
*/
type ClassLoader struct {
	cp       *classpath.Classpath // Classpath指针，用于搜索和读取class文件
	classMap map[string]*Class    // 记载已经加载的类数据，key是类的全限定名
}

/*
创建实例
 */
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp: cp,
		classMap: make(map[string]*Class),
	}
}

/*
法把类数据加载到方法区
 */
func (self *ClassLoader) LoadClass(name string) *Class {
	// 查找classMap，看类是否已经被加载
	if class, ok := self.classMap[name]; ok {
		return class // 类已经加载
	}
	return self.loadNonArrayClass(name)
}

/*
加载非数据类
 */
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 1. 找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
	// 2. 解析class文件，生成虚拟机可以使用的类数据，放入方法区
	class := self.defineClass(data)
	// 3. 进行链接
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}


func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data) // 将class文件转化为class结构体
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class) // 递归调用LoadClass()方法加载类的每一个直接接口
	self.classMap[class.name] = class
	return class
}
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}


func link(class *Class) {
	verify(class) // 验证类信息
	prepare(class) // 变量分配空间并给予初始值
}
func verify(class *Class) {
	// todo Java虚拟机规范4.10节详细介绍了类的验证算法，可以自己实现。
}
