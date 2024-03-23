package main

import (
	"fmt"
	"go-demo/assert"
	_ "go-demo/home"
	"go-demo/multi"
	"go-demo/polymorphic"
)

func main() {

	fmt.Println("hi")

	polymorphic.RunDemo()
	assert.RunDemo()

	fmt.Println("@@@@@@@@@@@@@@@")

	//home.RunSystem()
	fmt.Println("\n\n")

	multi.RunDemo()
}
