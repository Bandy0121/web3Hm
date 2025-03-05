package main

import "fmt"

type A struct {
	a     string
	bytes [2]byte
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

func value(a A, value string) {
	a.a = value
}

func point(a *A, value string) {
	a.a = value
}

func main() {
	a := A{
		a: "a",
	}

	value(a, "any")

	point(&a, "any")

	pa := &a

	// a *A
	// a.setPA("pa")

	// a A
	fmt.Println(a.string())
	// a A
	fmt.Println(a.stringA())
	// a *A
	fmt.Println(a.stringPA())

	// a A
	fmt.Println(pa.string())
	// a A
	fmt.Println(pa.stringA())
	// a *A
	fmt.Println(pa.stringPA())
}
