package _1数组

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 */

type Array struct {
	data   []int
	length uint
}

// NewArray 为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (arr *Array) Len() uint {
	return arr.length
}

//判断索引是否越界
func (arr *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(arr.data)) {
		return true
	}
	return false
}

// Find 通过索引查找数组，索引范围[0,n-1]
func (arr *Array) Find(index uint) (int, error) {
	if arr.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return arr.data[index], nil
}

// Insert 插入数值到索引index上
func (arr *Array) Insert(index uint, v int) error {
	if arr.Len() == uint(cap(arr.data)) {
		return errors.New("full array")
	}
	if index != arr.length && arr.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := arr.length; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = v
	arr.length++
	return nil
}

func (arr *Array) InsertToTail(v int) error {
	return arr.Insert(arr.Len(), v)
}

// Delete 删除索引index上的值
func (arr *Array) Delete(index uint) (int, error) {
	if arr.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := arr.data[index]
	for i := index; i < arr.Len()-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.length--
	return v, nil
}

// Print 打印数列
func (arr *Array) Print() {
	var format string
	for i := uint(0); i < arr.Len(); i++ {
		format += fmt.Sprintf("|%+v", arr.data[i])
	}
	fmt.Println(format)
}
