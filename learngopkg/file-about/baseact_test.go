package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
	"time"
)

const (
	dir1 = "tmp"
	dir2 = "tmpcopy"
)

func TestGetWorkDir(t *testing.T) {
	pwd, _ := os.Getwd()
	fmt.Println("当前工作目录：" + pwd)
}
func TestChangeDir(t *testing.T) {
	pwd, _ := os.Getwd()
	fmt.Println("目录1：" + pwd)
	if err := os.Chdir(dir1); err != nil {
		fmt.Println(err.Error())
	}
	pwd, _ = os.Getwd()
	fmt.Println("目录2：" + pwd)
	if err := os.Chdir(dir2); err != nil {
		fmt.Println(err.Error())
	}
	pwd, _ = os.Getwd()
	fmt.Println("目录3：" + pwd)
}

func TestGetDirFiles(t *testing.T) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func TestCreateDir(t *testing.T) {
	err := os.Mkdir("tmp", os.ModeDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("创建目录: " + dir1)
}

func TestCreateFile(t *testing.T) {
	// os.Create
	nFile, err := os.Create(dir1+"/file1")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer nFile.Close()
}


func TestRenameDir(t *testing.T) {
	err := os.Rename(dir1, dir2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("重命名：%s -> %s\n", dir1, dir2)
	}
}

func TestRemove(t *testing.T) {
	fs, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptn, _ := regexp.Compile("tmp*?")
	for _, f := range fs {
		if ptn.MatchString(f.Name()) {
			err = os.Remove(f.Name())
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(time.Millisecond * 500)
			fmt.Println("删除文件: " + f.Name())
		}
	}
}