package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Htest(url string ) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("???\n");
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
	// return rebots;
}


func Gtest(){
	curtime := time.Now();
	timestamp := curtime.Unix();
	// cs chan string;
	c1 := make(chan string);
	c2 := make(chan string);

	go func(c chan string){
		var res string;
		fmt.Printf("a:\n");
		Htest("http://local.c/curl_sleep1.php");
		fmt.Printf("a end\n");
		c <- res
	}(c1);

	go func(c chan string){
		var res string;
		fmt.Printf("b:\n");
		Htest("http://local.c/curl_sleep1.php");
		fmt.Printf("b end\n");
		c <- res
	}(c2);

	for{
		select{
		case  <- c1:
			fmt.Printf("a\n");
			fmt.Printf("%d\n",time.Now().Unix()-timestamp);
		case <- c2:
			fmt.Printf("b\n");
			fmt.Printf("%d\n",time.Now().Unix()-timestamp);
		}
	}
}