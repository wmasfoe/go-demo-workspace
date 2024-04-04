package sparseArray

import (
	"bufio"
	"errors"
	"fmt"
	log "go-demo/logger"
	"os"
	"regexp"
	"strconv"
)

// ArrToFile 稀疏数组保存到文件中
func (arr *SparseArray) ArrToFile() error {
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

// FileToArr 从指定文件读取
func (arr *SparseArray) FileToArr() error {
	file, err := os.OpenFile(arr.filePath, os.O_RDONLY, 0666)
	if err != nil {
		arr.value = []SparseItem{}
		log.Logger.Warn(fmt.Sprintf("读取的文件 %v 不存在，已经初始化为 []", arr.filePath))
		return nil
	}
	defer func() {
		file.Close()
	}()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\w+)`)

	// 清空稀疏数组 value，for 循环重新append
	arr.Reset()

	readIndex := 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		matches := re.FindStringSubmatch(lineText)

		rowValue, _ := strconv.Atoi(matches[1])
		colValue, _ := strconv.Atoi(matches[2])
		valValue := matches[3]

		item := SparseItem{
			row: rowValue,
			col: colValue,
			val: valValue,
		}

		arr.value = append(arr.value, item)

		readIndex++
	}

	arr.ShowValue()

	return nil
}

func (arr *SparseArray) Reset() {
	arr.value = []SparseItem{}
}

func (arr *SparseArray) Add(row int, col int, val any) {
	item := SparseItem{
		row: row,
		col: col,
		val: val,
	}
	arr.value = append(arr.value, item)
}
func (arr *SparseArray) Remove(row int, col int) error {
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
func (arr *SparseArray) ShowValue() {
	fmt.Println("当前内存中的内容是：")
	for _, nodeValue := range arr.value {
		fmt.Printf("%v, %v = %v\n", nodeValue.row, nodeValue.col, nodeValue.val)
	}
	fmt.Println(arr.value)
}

func (arr *SparseArray) GetRowSize() int {
	return arr.rowSize
}
func (arr *SparseArray) GetColSize() int {
	return arr.colSize
}
func (arr *SparseArray) GetValue() [][]any {
	row, col := arr.rowSize, arr.colSize

	resArr := make([][]any, row)
	for i := range resArr {
		resArr[i] = make([]any, col)
	}

	for _, nodeValue := range arr.value {
		resArr[nodeValue.row][nodeValue.col] = nodeValue.val
	}

	return resArr
}
