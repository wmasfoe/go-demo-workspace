package channel

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func RunDemo() {
	intCh := make(chan int, 1)
	intCh <- 1
	//关闭 ch
	close(intCh)
	fmt.Println("intCh == ", <-intCh)

	personCh := make(chan Person, 10)

	// 传入 100 个person
	for i := 0; i < 10; i++ {
		personCh <- Person{Name: "xiaoming" + fmt.Sprintf("%v", i), Age: i, Address: "bej"}
	}

	// 取出
	close(personCh)
	for i := range personCh {
		fmt.Printf("名字：%v, 年龄：%v, 地址：%v\n", i.Name, i.Age, i.Address)
	}
}
