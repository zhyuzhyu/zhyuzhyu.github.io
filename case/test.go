package main

import (
	"fmt"
	"os"

	"github.com/stretchr/testify/assert"
)

func main() {
	// 创建一个新的 Assert 对象
	a := assert.New(nil)

	// 检查一个条件是否为真
	value := "hello"
	if !a.Equal("hello", value) {
		fmt.Println("Value is not equal to 'hello'")
		os.Exit(1)
	}

	// 检查一个条件是否为假
	if a.Equal("world", value) {
		fmt.Println("Value should not be equal to 'world'")
		os.Exit(1)
	}

	fmt.Println("All assertions passed!")
}