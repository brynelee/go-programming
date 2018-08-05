package main

import "fmt"

func main(){
	str := "Hello, World!"
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
