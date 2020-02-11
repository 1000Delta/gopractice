package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("", "", 1024, 768)
	defer ui.Close()

	// 文件方式调用html文件
	f, err := os.OpenFile("./ui/index.html", os.O_RDONLY, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
	}
	ui.Load("data: text/html, " + url.PathEscape(string(data)))
	// 也可通过文件服务器方式调用html文件

	// 绑定 go 函数到 js body对象
	// body.start
	ui.Bind("start", func() string {
		return fmt.Sprintln("窗体已加载")
	})
	ui.Bind("printHello", func() string {
		return "Hello"
	})

	<-ui.Done()
}
