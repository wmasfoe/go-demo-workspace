package sort

import "fmt"

func ToSelectSorted(arr []int, mode string) []int {
	if mode != "asc" && mode != "desc" {
		fmt.Printf("不支持 %v 排序方式\n", mode)
		return arr
	}

	leftIndex := -1
	for {
		leftIndex++
		rightIndex := leftIndex
		minVal := arr[rightIndex]
		minValIndex := rightIndex
		for ; rightIndex < len(arr); rightIndex++ {
			currVal := arr[rightIndex]
			if minVal > currVal {
				minVal = arr[rightIndex]
				minValIndex = rightIndex
			}
		}

		temp := arr[leftIndex]
		arr[leftIndex] = minVal
		arr[minValIndex] = temp
		if leftIndex == (len(arr) - 1) {
			break
		}
	}

	return arr
}

func SelectSort(arr *[]int, mode string) {

	if mode != "asc" && mode != "desc" {
		fmt.Printf("不支持 %v 排序方式\n", mode)
		return
	}

	res := ToSelectSorted(*arr, mode)
	*arr = res
}
