package main

import (
	// "fmt"
	// "sync"

	"./test"
)

func main() {
	test.Json_test()
	// cls(func() {
	// 	fmt.Println("func")
	// })

	// test.Htest()
	// fmt.Println(test.Test_reg("s(.*)","ssbssb"))
	// Gtest();

	// test.Gtest();
}

func cls(task func()) {
	task()
}
