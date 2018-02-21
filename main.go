// my-test project main.go
package main

import (
	"fmt"
	"syscall"
	//_ "github.com/antonlahti/go-winapi"
	//"reflect"
)

func main() {
	fmt.Println("Hello World!")
	test_make_and_new()
	test_print()
	loadLib()
}

type point struct {
	x int
	y int
}

func test_make_and_new() {
	//内建函数 new 用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针
	//也可以为自定义类型来分配内存
	var i1 int = 1
	fmt.Printf("i1值为：%d\n", i1)
	var i2 *int = new(int)
	fmt.Printf("i2值为：%d\n", *i2)
	fmt.Printf("i2地址值为：%v\n", i2)
	*i2 = 2
	fmt.Printf("i2值为：%d\n", *i2)
	fmt.Printf("i2地址值为：%v\n", i2)

	//用于初始化自定义结构
	my_point := new(point)
	fmt.Printf("my_point值为：%v\n", *my_point)
	my_point.x = 5
	my_point.y = 10
	fmt.Printf("my_point地址值为：%v\n", &my_point)
	fmt.Printf("my_point地址值为：%v\n", my_point)

	//make 只能用于初始化slice,map,chan, new分配返回的是指针，即类型*T；make返回引用，即T
	slice := make([]int, 10)
	fmt.Printf("slice tpye:%T\n", slice)
	fmt.Printf("slice content:%v\n", slice)

}

func test_print() {
	// %v 以默认的方式打印变量的值，自动推导类型，对于自定义的结构体也适用
	i := 32.0
	fmt.Printf("i=%v\n", i)
	// %+v 打印结构体的字段名称
	p := point{3, 4}
	fmt.Printf("p:%+v\n", p)
	//%T 打印变量的类型
	fmt.Printf("type of i:%T\n", i)
	fmt.Printf("type of p:%T\n", p)
	// %+d 带符号的整型，fmt.Printf("%+d", 255)输出+255
	fmt.Printf("i=%+d\n", 32)
	//%q 打印单引号
	fmt.Printf("带引号：%q\n", "带单引号的字符串。。。")
	// 不带引号的
	fmt.Printf("不带引号的字符串:%s\n", "大噶好，我系古天乐")
	// 布尔值
	fmt.Printf("%t\n", true)
	// 二进制表示，该占位符不能打印布尔值
	fmt.Printf("%b\n", 32)
	// 单个字符打印
	fmt.Printf("%c\n", 65)
	// 16进制打印
	fmt.Printf("%x\n", &i)
	//输出指针的值
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", &p)
}

func loadLib() {
	ole32, err := syscall.LoadLibrary("ole32.dll")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ole32)
}
