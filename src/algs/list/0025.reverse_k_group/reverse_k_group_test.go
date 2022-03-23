package _025_reverse_k_group

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	list := data_structure.NewListNodeFromArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	newList := reverseKGroup(list, 2)
	println(data_structure.GetListNodeStr(newList))
}
