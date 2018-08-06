package main

import "fmt"

func main(){

	// 数组是值类型
	var a = [3]int{1, 2, 3}
	var b = a
	b[1]++
	fmt.Println(a, b)

	var c = [3]int{1, 2, 3}
	var d = &c
	d[1]++
	fmt.Println(c, *d)
}
