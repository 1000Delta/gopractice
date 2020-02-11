package main

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(10).LRU().Build()
	// 存入数据
	gc.Set("data1", "set data")
	// 存入有期限数据
	gc.SetWithExpire("data2", "dataWithExpire", time.Minute * 5)
	// 通过回调自动存入数据
	gcl := gcache.New(10).ARC().LoaderFunc(func(key interface{}) (interface{}, error) {
		// 调用数据库逻辑
		if key == "key" {
			return "dataAutoLoad", nil
		}
		return "dataIsNotExist", nil
	}).Build()
	// gc
	value, err := gc.Get("data1")
	if err != nil {
		fmt.Println("报错")
	}
	fmt.Println(value)
	value, err = gc.Get("data3")
	if err != nil {
		fmt.Println("报错")
	}
	fmt.Println(value)
	// gc1 auto load
	value, err = gc.Get("data1")
	if err != nil {
		fmt.Println("报错")
	}
	fmt.Println(value)
	value, err = gcl.Get("key")
	if err != nil {
		fmt.Println("报错")
	}
	fmt.Println(value)
}
