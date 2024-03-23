package multi

import (
	"fmt"
	"time"
)

// 协程

func test() {
	for i := 0; i < 4; i++ {
		fmt.Println("test()", i+1)
		time.Sleep(time.Millisecond * 60)
	}
}

func RunDemo() {

	go test()

	for i := 0; i < 4; i++ {
		fmt.Println("Multi RunDemo()", i+1)
		time.Sleep(time.Millisecond * 60)
	}
}
