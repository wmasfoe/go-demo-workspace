package multi

import (
	"fmt"
	"time"
)

// 协程

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test()", i+1)
		time.Sleep(time.Second)
	}
}

func RunDemo() {

	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("Multi RunDemo()", i+1)
		time.Sleep(time.Second)
	}
}
