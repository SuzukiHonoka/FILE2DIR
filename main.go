package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const pfn = "main.go"


func main() {
	var args = os.Args
	var dir, _ = os.Getwd()
	var spath = "."
	if len(args) > 1 {
		spath = args[1]
	}
	if spath == "." {
		spath = dir
	}
	file, _ := ioutil.ReadDir(dir)
	for _, f := range file {
		if !f.IsDir() {
			name := f.Name()
			if name != pfn {
				fext := filepath.Ext(name)
				ext := ""
				if len(fext) > 1 {
					ext = strings.ToUpper(fext[1:])
				}
				abss := filepath.Join(spath, name)
				absd := filepath.Join(spath, ext)
				status, err := movefile(abss, absd)
				if status {
					fmt.Println("Move:", abss, absd, "Succeed.")
				} else {
					fmt.Println(err.Error())
				}
			}
		}
	}
}

func movefile(src string, dst string) (bool, error) {
	if !exists(dst) {
		err := os.Mkdir(dst, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	err := os.Rename(src, filepath.Join(dst, path.Base(src)))
	if err != nil {
		return false, err
	}
	return true, nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
