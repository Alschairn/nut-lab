package classpath

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/**
	定义结构体
 */
type DirEntry struct {
	absDir string
}

/**
	无专门构造函数，统一使用new开头标识
 */
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) // 转换为绝对路径
	if err != nil {
		panic(err) // 转换失败，程序崩溃？？
	}
	return &DirEntry{absDir} // 构建对象，返回
}

/**
	不需要显式声明实现，方法匹配即可
 */
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, className) // 拼接成完整的文件路径
	fmt.Printf("entry_dir readClass filename: %v \n", filename)
	data, err := ioutil.ReadFile(filename) // 读取文件
	return data, self, err // 返回
}

func (self *DirEntry) String() string{
	return self.absDir
}