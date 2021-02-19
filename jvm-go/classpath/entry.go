package classpath

import (
	"os"
	"strings"
)

// 定义常量，存放系统分隔符
const pathListSeparator = string(os.PathListSeparator)

/**
	定义类路径接口
 */
type Entry interface {

	/*
		根据className初始化类路径实例
		className: 类相对路径
	 */
	readClass(className string) ([]byte, Entry, error)

	/**
		返回变量的字符串表示（类似Java toString方法）
	 */
	String() string
}

/**
	根据不同参数创建不同类型Entry实例
 */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.Contains(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			return newZipEntry(path)
	}
	return newDirEntry(path)
}
