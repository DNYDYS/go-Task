package main

import "fmt"

// 面向对象第一题
func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}
func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Width + rec.Height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.r
}

// 长方形
type Rectangle struct {
	Width  float64
	Height float64
}

// 圆
type Circle struct {
	r float64
}

// 面向对象 第二题
type Person struct {
	Name string
	age  int
}

type Employee struct {
	person    Person
	EmplyeeId int
}

func (e Employee) PrintInfo() {
	fmt.Println("此员工的名字=", e.person.Name, " 年龄=", e.person.age, " 员工编号=", e.EmplyeeId)
}

func main() {
	// 面向对象第一题
	rectangle := Rectangle{Width: 2, Height: 3}
	circle := &Circle{r: 5}
	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())

	// 第二题
	person := Person{"YoLanDa", 18}
	employee := Employee{person, 1}
	employee.PrintInfo()

}
