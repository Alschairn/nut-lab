package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) // 打开zip文件
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File { // 遍历zip中的文件
		if f.Name == className { // 寻找class文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self ,nil
		}
	}
	return nil, nil, errors.New("Class not found: " + className) // 没有找到classs文件，返回异常
}

func (self *ZipEntry) String() string {
	return self.absPath
}