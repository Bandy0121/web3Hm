package main

import "fmt"

func main() {
	a, b := 1, 2
	var c int
	c = a + b
	fmt.Println("c = a + b, c =", c)

	plusAssignment(c, a)
	subAssignment(c, a)
	mulAssignment(c, a)
	divAssignment(c, a)
	modAssignment(c, a)
	leftMoveAssignment(c, a)
	rightMoveAssignment(c, a)
	andAssignment(c, a)
	orAssignment(c, a)
	norAssignment(c, a)
}

func plusAssignment(c, a int) {
	c += a // c = c + a
	fmt.Println("c += a, c =", c)
}

func subAssignment(c, a int) {
	c -= a // c = c - a
	fmt.Println("c -= a, c =", c)
}

func mulAssignment(c, a int) {
	c *= a // c = c * a
	fmt.Println("c *= a, c =", c)
}

func divAssignment(c, a int) {
	c /= a // c = c / a
	fmt.Println("c /= a, c =", c)
}

func modAssignment(c, a int) {
	c %= a // c = c % a
	fmt.Println("c %= a, c =", c)
}

func leftMoveAssignment(c, a int) {
	c <<= a // c = c << a
	fmt.Println("c <<= a, c =", c)
}

func rightMoveAssignment(c, a int) {
	c >>= a // c = c >> a
	fmt.Println("c >>= a, c =", c)
}

func andAssignment(c, a int) {
	c &= a // c = c & a
	fmt.Println("c &= a, c =", c)
}

func orAssignment(c, a int) {
	c |= a // c = c | a
	fmt.Println("c |= a, c =", c)
}

func norAssignment(c, a int) {
	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c =", c)
}
