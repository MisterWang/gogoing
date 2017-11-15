package test

import (
	"encoding/json"
	"fmt"
)

func Json_test() {
	var j map[string]interface{}
	if err := json.Unmarshal([](byte)(`{"a":"b","c":[{"d":1}]}`), &j); err == nil {
		fmt.Println(j)
	}

	if str, err := json.Marshal(j); err == nil {
		fmt.Println((string)(str))
	}
}
