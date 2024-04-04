package sparseArray

type SparseItem struct {
	row int
	col int
	val any
}

type SparseArray struct {
	rowSize    int
	colSize    int
	defaultVal int
	value      []SparseItem
	filePath   string
}

// SparseArrayFileOperate 定义了 稀疏数组关于文件的操作
type SparseArrayFileOperate interface {
	ArrToFile() (err error)
	FileToArr() (err error)
}

// ISparseArray 定义了 稀疏数组对本身的操作
type ISparseArray interface {
	Add(row int, col int, val any)
	Remove(row int, col int) error
	Reset()
	GetValue() [][]any
	GetRowSize() int
	GetColSize() int
}
