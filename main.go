package main

import (
	"fmt"

	"./test"
)

func main() {
	test.Json_test()
	cls(func() {
		fmt.Println("func")
	})
}

func cls(task func()) {
	task()
}