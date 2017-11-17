package main

import (
	"fmt"

	"./test"
)

func main() {
	// test.Json_test()
	// cls(func() {
	// 	fmt.Println("func")
	// })

	// test.Htest()
	fmt.Println(test.Test_reg("s(.*)","ssbssb"))
}

func cls(task func()) {
	task()
}
