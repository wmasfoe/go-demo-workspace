package assert

import "fmt"

type Point struct {
	x int
	y int
}

func RunDemo() {
	var any interface{}

	point := Point{
		x: 1,
		y: 1,
	}

	any = point
	fmt.Println("any = point", any)

	var b Point
	// 直接 b=any 是编译失败的，需要 .(type) 断言
	b = any.(Point)
	fmt.Println("b = any", b)

	var any2 interface{}
	f := 1.1

	any2 = f

	// float 默认是 64，这里断言32，断言失败了
	if y, ok := any2.(float32); ok {
		fmt.Println(y)
	} else {
		fmt.Println("断言失败")
	}


}
