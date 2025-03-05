package main

import "fmt"

func main() {
	var m1 map[string]string
	fmt.Println("m1 length:", len(m1))

	m2 := make(map[string]string)
	fmt.Println("m2 length:", len(m2))
	fmt.Println("m2 =", m2)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 length:", len(m3))
	fmt.Println("m3 =", m3)

	m4 := map[string]string{}
	fmt.Println("m4 length:", len(m4))
	fmt.Println("m4 =", m4)

	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("m5 length:", len(m5))
	fmt.Println("m5 =", m5)
}
