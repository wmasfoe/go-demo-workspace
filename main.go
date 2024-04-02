package main

import (
	"fmt"
	"go-demo/assert"
	"go-demo/channel"
	_ "go-demo/chat"
	_ "go-demo/home"
	"go-demo/multi"
	"go-demo/polymorphic"
	sparseArray "go-demo/sparse-array"
)

func main() {

	fmt.Println("hi")

	polymorphic.RunDemo()
	assert.RunDemo()

	fmt.Println("@@@@@@@@@@@@@@@")

	//home.RunSystem()
	fmt.Println("\n\n")

	multi.RunDemo()

	channel.RunDemo()
	channel.RunDemo2()

	//chat.RunDemo()

	sparseArray.RunDemo()

}
