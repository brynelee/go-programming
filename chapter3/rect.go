package main

import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main(){
	fmt.Println(*NewRect(3,5,200,400))
}