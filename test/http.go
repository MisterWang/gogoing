package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Htest() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
