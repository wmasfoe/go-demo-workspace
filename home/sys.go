package home

import "fmt"

func RunSystem() {

	defer func() {
		fmt.Println("感谢使用~")
	}()

	// 用户选择的操作
	opt := ""

exit:
	for {

		fmt.Println("记账软件，选择你的功能：")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记收入")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出软件")
		fmt.Println("请输入：")

		_, err := fmt.Scanln(&opt)
		if err != nil {
			fmt.Println("Scanln error ", err)
		}

		switch opt {
		case "1":
			fmt.Println("当前收支明细为：")
		case "2":
			fallthrough
		case "3":
			fmt.Println("登记输入输出：")
		case "4":
			break exit
		default:
			fmt.Println("非法输入")
		}
	}

}
