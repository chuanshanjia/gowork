package main

import (
	"fmt"
)

func main() {
	var map_test = make(map[string]string)
	key := "name"
	map_test[key] = "shanjia"
	if _, exists := map_test["ok"]; exists {
		fmt.Printf("key[%s] exists\n", "ok")
	} else {
		fmt.Printf("key[%s] not exists\n", "ok")
	}

	if value, exists := map_test[key]; exists {
		fmt.Printf("key[%s] exists, value[%s]\n", key, value)
	} else {
		fmt.Printf("key[%s] not exists\n", key)
	}

	value, ok := map_test[key]
	if ok {
		fmt.Printf("key[%s] exists, value[%s]\n", key, value)
	}
}
