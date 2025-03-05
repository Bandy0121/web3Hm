package main

import "fmt"

type A struct {
	a string
}

func (a A) string() string {
	return a.a
}

func (a A) stringA() string {
	return a.a
}

func (a A) setA(v string) {
	a.a = v
}

func (a *A) stringPA() string {
	return a.a
}

func (a *A) setPA(v string) {
	a.a = v
}

type B struct {
	A
	b string
}

func (b B) string() string {
	return b.b
}

func (b B) stringB() string {
	return b.b
}

type C struct {
	B
	a string
	b string
	c string
	d []byte
}

func (c C) string() string {
	return c.c
}

func (c C) modityD() {
	c.d[2] = 3
}

func callStructMethod() {
	var a A
	a = A{
		a: "a",
	}
	a.string()
}

func NewC() C {
	return C{
		B: B{
			A: A{
				a: "ba",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "c",
		d: []byte{1, 2, 3},
	}
}

func main() {
	c := NewC()
	cp := &c
	fmt.Println(c.string())
	fmt.Println(c.stringA())
	fmt.Println(c.stringB())

	fmt.Println(cp.string())
	fmt.Println(cp.stringA())
	fmt.Println(cp.stringB())

	c.setA("1a")
	fmt.Println("------------------c.setA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setA("2a")
	fmt.Println("------------------cp.setA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	c.setPA("3a")
	fmt.Println("------------------c.setPA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setPA("4a")
	fmt.Println("------------------cp.setPA")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.modityD()
	fmt.Println("------------------cp.modityD")
	fmt.Println(cp.d)
}
