package channel

import "fmt"

func writeFactorialFn(n int, valueChan chan map[int]int, exitChan chan bool) {
	res := 1
	// 计算 n!
	for i := 1; i <= n; i++ {
		res *= i
	}
	//push 到 chan
	valueChan <- map[int]int{n: res}
	exitChan <- true
}

// 算1-20的阶乘，放到 map 中
func factorialDemo() {
	mapIntChan := make(chan map[int]int, 20)
	exitChan := make(chan bool, 20)

	// 开20个协程，批量写入 chan
	for i := 1; i <= 20; i++ {
		go writeFactorialFn(i, mapIntChan, exitChan)
	}

	//再开一个协程，专门用来关闭 chan，以便终止 range
	go func() {
		for i := 0; i < 20; i++ {
			<-exitChan
		}
		close(mapIntChan)
		close(exitChan)
	}()

	for value := range mapIntChan {
		fmt.Println("阶乘chan value === ", value)
	}
}

func isPrime(num int) bool {
	res := true
	if 1 >= num {
		res = false
	}
	for i := 2; i < num; i++ {
		if 0 == num%i {
			res = false
			break
		}
	}
	return res
}

// 向 valueChan push 数据
func putNumList(valueChan chan int) {
	for i := 2; i <= 800; i++ {
		valueChan <- i
	}

	// push 完 close，以便终止 4 个 primeNumList
	close(valueChan)
}

func primeNumList(primeChan chan int, valueChan chan int, exitChan chan bool) {
	for value := range valueChan {
		if flag := isPrime(value); flag {
			primeChan <- value
		}
	}

	exitChan <- true
}

// 开4个协程，计算 1-800之间的素数，放到 slice 中
func primeDemo() {
	//包含 1-800 数字的 chan
	valueChan := make(chan int, 800)
	//包含计算的素数的chan
	primeChan := make(chan int, 800)
	//标识协程计算完成的 chan
	exitChan := make(chan bool, 4)

	putNumList(valueChan)

	// 开4个协程
	for i := 0; i < 4; i++ {
		primeNumList(primeChan, valueChan, exitChan)
	}

	go func() {
		// 四个协程都完成，关闭 primeChan
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()

	for i := range primeChan {
		fmt.Println("素数有：", i)
	}
}

func selectDemo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("selectErrorDemo 出现 panic", err)
		}
	}()
	strChan := make(chan string, 5)
	for i := 1; i <= 5; i++ {
		strChan <- fmt.Sprintf("%d", i)
	}

	//使用 for 遍历 没有close的 chan，程序会死锁deadlock，可以使用 select
forWrap:
	for {
		select {
		case v := <-strChan:
			fmt.Println("strChan", v)
		//strChan 取不到数据时，走default
		default:
			break forWrap
		}
	}
}

func RunDemo2() {
	factorialDemo()
	primeDemo()
	selectDemo()
}
