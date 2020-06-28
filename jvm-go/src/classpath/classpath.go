package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type Classpath struct {

	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

/**
	根据参数初始化类路径
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/**
	初始化系统类路径
 */
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	fmt.Printf("parseBootAndExtClasspath jreLibPath: %v \n", jreLibPath)
	self.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	fmt.Printf("parseBootAndExtClasspath jreExtPath: %v \n", jreExtPath)
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/**
	初始化用户类路径
	如果没有提供 -classpath/-cp 实现，则使用当前目录作为用户类路径
 */
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/**
	寻找jre目录
	优先使用用户输入的-Xjre选项作为jre目录
	未找到在当前目录下寻找jre目录
	未找到尝试使用环境变量JAVA_HOME
 */
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

/**
	判断路径是否存在
 */
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	fmt.Printf("classpath readClass classname: %v \n", className)
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
