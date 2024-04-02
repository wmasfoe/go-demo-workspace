package sparseArray

import (
	"bufio"
	"errors"
	"fmt"
	log "go-demo/logger"
	"os"
)

func (arr *SparseArray) arrToFile() error {
	// 打开文件， os.O_CREATE|os.O_WRONLY 表示没有文件创建文件，并且以写入的方式打开，0666 表示所有用户可读写(rwx)
	file, err := os.OpenFile(arr.filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Logger.Error(err.Error())
	}
	defer func() {
		file.Close()
	}()

	writer := bufio.NewWriter(file)

	for _, nodeVal := range arr.value {
		str := fmt.Sprintf("%v \t %v \t %v\r\n", nodeVal.row, nodeVal.col, nodeVal.val)
		_, err := writer.WriteString(str)
		if err != nil {
			log.Logger.Error(err.Error())
		}
	}
	flushErr := writer.Flush()
	if flushErr != nil {
		log.Logger.Error(flushErr.Error())
	}

	return nil
}
func (arr *SparseArray) fileToArr() error {
	return nil
}

func (arr *SparseArray) add(row int, col int, val any) {
	item := SparseItem{
		row: row,
		col: col,
		val: val,
	}
	arr.value = append(arr.value, item)
}
func (arr *SparseArray) remove(row int, col int) error {
	//要删除的元素的下标，默认-1
	tarIndex := -1
	temp := (*arr).value
	for i, nodeValue := range arr.value {
		if nodeValue.row == row && nodeValue.col == col {
			tarIndex = i
		}
	}
	if tarIndex == -1 {
		return errors.New("[warn]: remove 未找到对应的元素")
	} else {
		// 拼接 tarIndex 的前半部分数组
		arr.value = temp[0:tarIndex]
		// 判断是否是最后一个
		if len(temp)-1 > tarIndex {
			// 品牌 tarIndex 的后半部分数组
			arr.value = append(arr.value, temp[tarIndex+1:]...)
		}
		return nil
	}
}
func (arr *SparseArray) showValue() {
	fmt.Println("当前内存中的内容是：")
	for _, nodeValue := range arr.value {
		fmt.Printf("%v, %v = %v\n", nodeValue.row, nodeValue.col, nodeValue.val)
	}
	fmt.Println(arr.value)
}

func RunDemo() {

	// 10x10 的稀疏数组，默认值是 0，存储位置在 assets/sparse-array.txt
	sa := SparseArray{
		rowSize:    10,
		colSize:    10,
		defaultVal: 0,
		filePath:   "./sparse-array/_log/data.txt",
	}
	sa.add(1, 1, "hahaha")
	sa.add(2, 1, "heiheihei")
	sa.add(3, 2, "gagaga")
	sa.showValue()
	err := sa.remove(2, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	sa.showValue()

	sa.arrToFile()
}
