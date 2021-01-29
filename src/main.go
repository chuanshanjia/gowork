package main

import (
	"fmt"
)

//  比main先执行
func init() {
	fmt.Println("init")
}

func main() {
	var str1 string
	str1 = `
多行字符串需要使用反引号，多用于内嵌源码和内嵌数据
在反引号中的所有代码不会被编译器识别，而只是作为字符串的一部分
`
	fmt.Println(str1)
	/*输出*/
	fmt.Println("Hello world!")

	fmt.Printf("%T, %o \n", 123, 123)
}
