package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestModel struct {
	gorm.Model
	Name string `json:"name" gorm:"index"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// 迁移数据库
	db.AutoMigrate(&TestModel{})

	testData := &TestModel{
		Name: "test",
	}

	// 插入数据
	if err := db.Create(testData).Error; err != nil {
		fmt.Println(testData.ID)
	}

	// 读取数据
	data := []*TestModel{}
	if err := db.Find(&data).Error; err != nil {
		fmt.Println(err.Error())
	}

	for _, d := range data {
		fmt.Printf("%d %s\n", d.ID, d.Name)
	}
}
