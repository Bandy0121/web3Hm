package main

import "fmt"

func main() {
	// str1 := "abc123"
	// for index := range str1 {
	// 	fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	// }

	// str2 := "测试中文"
	// for index := range str2 {
	// 	fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	// }
	// fmt.Printf("len(str2) = %d\n", len(str2))

	// runesFromStr2 := []rune(str2)
	// bytesFromStr2 := []byte(str2)
	// fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	// fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))

	// array := [...]int{1, 2, 3}
	// slice := []int{4, 5, 6}

	// // 方法1：只拿到数组的下标索引
	// for index := range array {
	// 	fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	// }
	// for index := range slice {
	// 	fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	// }
	// fmt.Println()

	// // 方法2：同时拿到数组的下标索引和对应的值
	// for index, value := range array {
	// 	fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
	// 	fmt.Printf("array -- index=%d range value=%d \n", index, value)
	// }
	// for index, value := range slice {
	// 	fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
	// 	fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	// }
	// fmt.Println()

	hash := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}

	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	//range 关键字在迭代映射集合时，其中的 key 是乱序的
	/* 	key=c, value=4
	   	key=a, value=1
	   	key=f, value=2
	   	key=z, value=3
	   	key=z, value=3
	   	key=c, value=4
	   	key=a, value=1
	   	key=f, value=2 */
}
