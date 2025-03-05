package main

import (
	"fmt"
	// _ "init_order/pkg1"
)

// const mainName string = "main"

// var mainVar string = getMainVar()

// func init() {
// 	fmt.Println("main init method invoked")
// }

func main() {
	// fmt.Println("main method invoked!")
	// 十六进制
	// var a uint8 = 0xF
	// var b uint8 = 0xf

	// // 八进制
	// var c uint8 = 017
	// var d uint8 = 0o17
	// var e uint8 = 0o17

	// // 二进制
	// var f uint8 = 0b1111
	// var g uint8 = 0b1111

	// // 十进制
	// var h uint8 = 15

	// var float1 float32 = 10
	// float2 := 10.0
	// float1 = float32(float2)
	/* 复数：complex64，complext128
	   整型与浮点数日常中常见的数字都是实数，复数是实数的延伸，可以通过两个部分构成，一个实部，一个虚部，常见的声明形式如下：
	   var z complex64 = a + bi
	   a 和 b 均为实数，i 为虚数单位，当 b = 0 时，z 就是常见的实数。
	   当 a = 0 且 b ≠ 0 时，将 z 为纯虚数。 */
	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1) // c2与c3是等价的
	fmt.Println(c1 == complex64(c2))
	fmt.Println(complex128(c1) == c2)
	fmt.Println(c2 == c3)
	fmt.Println(c2 == c3)
	x := real(c2) //实部
	y := imag(c2) //虚部
	fmt.Println(x)
	fmt.Println(y)

	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)

	var _bytes []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var _s string = string(_bytes)
	fmt.Println(_s)
	/* `rune` 是 `int32` 的内置别名，可以把 `rune` 和 `int32` 视为同一种类型。但 rune 是特殊的整数类型。
	   在 Go 中，一个 rune 值表示一个 Unicode 码点。一般情况下，一个 Unicode 码点可以看做一个 Unicode 字符。有些比较特殊的 Unicode 字符有多个 Unicode 码点组成。
	*/
	var r1 rune = 'a'
	var r2 rune = '世'
	fmt.Println(r1) //97
	fmt.Println(r2) //19990

	var srune string = "abc，你好，世界！"
	var runes []rune = []rune(srune)
	fmt.Println(runes)
	fmt.Println(len(runes))

	var s1 string = "Hello\nworld!\n"
	var s2 string = `Hello
world!
`
	fmt.Println(s1 == s2)

	comparebyte_rune_string()
}

func comparebyte_rune_string() {
	var s string = "Go语言"
	var bytes []byte = []byte(s)
	var runes []rune = []rune(s)

	fmt.Println("string length: ", len(s))
	fmt.Println("bytes length: ", len(bytes))
	fmt.Println("runes length: ", len(runes))

	fmt.Println("string sub: ", s[0:7])
	fmt.Println("bytes sub: ", string(bytes[0:7]))
	fmt.Println("runes sub: ", string(runes[0:3]))
}

// func getMainVar() string {
// 	fmt.Println("main.getMainVar method invoked!")
// 	return mainName
// }
