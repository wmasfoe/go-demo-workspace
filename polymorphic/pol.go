package polymorphic

import (
	"fmt"
)

type USB interface {
	start()
	stop()
}

type Phone struct {
	Name string
}
func (p Phone) start()  {
	fmt.Println(p.Name, "开始工作")
}
func (p Phone) stop()  {
	fmt.Println(p.Name, "断开连接~")
}
func (p Phone) call(tel uint) {
	fmt.Println("打电话给: ", tel)
}


type Camera struct {
	Name string
}
func (c Camera) start() {
	fmt.Println(c.Name, "开始工作")
}
func (c Camera) stop() {
	fmt.Println(c.Name, "断开连接~~")
}

/*
USB 是多态的，可以被多个 struct 实现。
而 ConnectComputer 这个方法，可以传入实现了 USB 接口的参数，故 USB 多态
如果后续 Phone 或者 Camera 没有实现 USB，那么这个方法会编译错误
*/
func ConnectComputer(usb USB) {
	defer usb.stop()
	usb.start()
	fmt.Println("balabalabala~~~√∆çå∂√∆˚hdkjs˚")

	// 如果可以断言为 Phone ，调用它的 call 方法
	if usingUsb, isPhone := usb.(Phone); isPhone {
		usingUsb.call(12345678901)
	}
}

type Student struct {
	Name string
}

func judgeType (v interface{}) {
	switch v.(type) {
	case Student:
		fmt.Println("Student")
	case *Student:
		fmt.Println("*Student point")

	}
}

func RunDemo() {

	iphone := Phone{
		Name: "iPhone18 Pro",
	}
	snnoy := Camera{
		Name: "索尼X8",
	}

	ConnectComputer(iphone)
	ConnectComputer(snnoy)

	fmt.Println("=======")

	stu := Student{Name: "小明"}
	stuP := &Student{Name: "小红"}

	judgeType(stuP)
	judgeType(stu)
}


