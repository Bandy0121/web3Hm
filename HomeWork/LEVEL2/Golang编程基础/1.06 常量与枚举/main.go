package main

// 枚举的本质就是一系列的常量。所以 Go 中使用 const 定义枚举，比如：
// const (
// 	Male   = "Male"
// 	Female = "Female"
// )

func main() {
	// 方式1
	const a int = 1

	// 方式2
	const b = "test"

	// 方式3
	const c, d = 2, "hello"

	// 方式4
	const e, f bool = true, false

	// 方式5
	const (
		h    byte = 3
		i         = "value"
		j, k      = "v", 4
		l, m      = 5, false
	)

	const (
		n = 6
	)
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

// 除了直接定义值以外，还会使用类型别名，让常量定义的枚举类型的作用显得更直观，比如：
func example1() {
	type ConnState int
	const (
		StateNew ConnState = iota
		StateActive
		StateIdle
		StateHijacked
		StateClosed
	)
}
func (g *Gender) String() string {
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g *Gender) IsMale() bool {
	return *g == Male
}

// 修改指针指向的值
func example2() {
	type Month int
	const (
		January Month = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
}

// 指针、unsafe.Pointer 和 uintptr
// 注意，这个操作非常危险，并且结果不可控，在一般情况下是不需要进行这种操作。
func example3() {
	const pre int = 1
	const a int = iota
	const (
		b int = iota
		c
		d
		e
	)
	const (
		f = 2
		g = iota
		h
		i
	)
	//iota 仅能与 const 关键字配合使用
	type Gender byte
	const (
		Male Gender = iota
		Female
	)
}
