package golib

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)

}

func IsExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func IsDirExists(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

func CheckDir(dirname string, mode os.FileMode) error {
	ok := IsDirExists(dirname)
	if !ok {
		err := os.MkdirAll(dirname, mode)
		if err != nil {
			return errors.New(fmt.Sprintf("%s:make dir err!", dirname))
		}
	}
	return nil
}
