package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p1 *int
	var p2 *string

	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	example1()
	example2()
	example3()
}

// 使用指针访问值
func example1() {
	var p1 *int
	i := 1
	p1 = &i
	fmt.Println(*p1 == i)
	*p1 = 2
	fmt.Println(i)
}

// 修改指针指向的值
func example2() {
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}

// 指针、unsafe.Pointer 和 uintptr
// 注意，这个操作非常危险，并且结果不可控，在一般情况下是不需要进行这种操作。
func example3() {
	a := "Hello, world!"
	fmt.Println(&a)
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(c)
	fmt.Println(*c)
}
