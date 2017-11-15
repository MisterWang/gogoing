package test

import "fmt"

func test_chan() {
	c := make(chan string)
	go channel(c)

	c <- "hello"
	str := <-c
	fmt.Println(str)
	// fmt.Println(<-c)
}

func channel(c chan string) {
	fmt.Println(<-c)
	c <- "123"
}
