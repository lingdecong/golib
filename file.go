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

func IsExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func IsDirExists(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

func CheckDir(file string) error {
	ok := IsDirExists(file)
	if !ok {
		err := os.MkdirAll(file, 0777)
		if err != nil {
			return errors.New(fmt.Sprintf("%s:make dir err!", file))
		}
	}
	return nil
}
