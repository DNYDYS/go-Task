package main

import (
	"fmt"
	"task1/util"
)

type Employees struct {
	Id         int
	Name       string  `grom:"column:name"`
	Department string  `grom:"column:department"`
	Salary     float32 `grom:"column:salary"`
}

type Book struct {
	Id     int
	Title  string  `grom:"column:title"`
	Author string  `grom:"column:author"`
	Price  float32 `grom:"column:price"`
}

func main() {
	db := util.ConnDB()

	//题目一
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employees []Employees
	if err := db.Debug().Where("department", "技术部").Find(&employees).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("技术部的员工信息", employees)

	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employee Employees
	if err := db.Debug().Where(max("salary")).Find(&employee).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("工资最高的员工信息：", employee)

	// 题目二
	//定义一个 Book 结构体，包含与 books 表对应的字段。
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	var book []Book
	if err := db.Debug().Where("price > ?", 90).Find(&book).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("价格大于90元的书的信息：", book)

}
