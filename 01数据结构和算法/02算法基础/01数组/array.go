package _1数组

// Array 用切片模拟一个数组
type Array struct {
	data   []int
	length int
}

// NewArray 创建数组的工厂方法
func NewArray(cap uint) *Array {
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

func (arr *Array) Len() int {
	return arr.length
}

func (arr *Array) Find(index uint) (int, error) {
	return arr.data[index], nil
}

func (arr *Array) Insert(index uint, v int) error {

}

func (arr *Array) InsertToTail(v int) error {

}

func (arr *Array) Delete(index uint) (int, error) {

}
