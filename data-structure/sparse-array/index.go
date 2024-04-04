package sparseArray

import (
	"fmt"
)

func RunDemo() {

	// 10x10 的稀疏数组，默认值是 0，存储位置在 ./data-structure/sparse-array/_log/data.txt
	sa := SparseArray{
		rowSize:    10,
		colSize:    10,
		defaultVal: 0,
		filePath:   "./data-structure/sparse-array/_log/data.txt",
	}
	sa.Add(1, 1, "hahaha")
	sa.Add(2, 1, "heiheihei")
	sa.Add(3, 2, "gagaga")
	sa.ShowValue()
	err := sa.Remove(2, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	sa.ShowValue()

	sa.ArrToFile()

	fmt.Println("重置..")
	sa.Reset()
	sa.ShowValue()

	sa.FileToArr()

	arr := sa.GetValue()
	fmt.Println(arr)
}
