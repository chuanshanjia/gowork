package main

import (
	"fmt"
	"github.com/chuanshanjia/gowork/goinaction/code/chapter2/sample/search"
)

//  比main先执行
func init() {
	fmt.Println("init")
}

func main() {
	// Perform the search for the specified term
	search.Run("president")
}
