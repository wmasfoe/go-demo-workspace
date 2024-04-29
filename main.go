package main

import (
	"fmt"
	"go-demo/assert"
	"go-demo/channel"
	_ "go-demo/chat"
	"go-demo/data-structure/sparse-array"
	_ "go-demo/home"
	"go-demo/leetcode/leetcode0003"
	"go-demo/multi"
	"go-demo/polymorphic"
	"go-demo/sort"
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

	leetcode0003.RunDemo()

	randArr := []int{4, 5, 8, 2, 3, 9, 1}
	newArr := sort.ToSelectSorted(randArr, "asc")
	fmt.Println(randArr, newArr)

}
