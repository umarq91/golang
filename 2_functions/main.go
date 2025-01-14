package main

import (
	"fmt"
)


func main() {
	a, b := swap(21, 40)
	fmt.Println(a, b)
	total, product := calc(30, 50)
	fmt.Println(total, product)
loops()
}

// multiple returned functions
func swap(a int, b int) (int, int) {
	return a, b
}

// Named returned ,  This allows you to directly assign values to the named return variables without needing to use the return keyword explicitly.
func calc(a int, b int) (total int, product int) {
	total = a + b
	product = a * b
	return
}

func loops(){
	for i:= 0 ; i<10 ; i++ {
		fmt.Println(i)
	}
}

func whileLoop(){
	sum:=0
	for sum<10{
		fmt.Println(sum)
	}
}