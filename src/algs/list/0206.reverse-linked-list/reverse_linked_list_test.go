package _206_reverse_linked_list

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	list := data_structure.NewListNodeFromArr([]int{1, 4, 7, 10})
	println(data_structure.GetListNodeStr(reverseLinkedList(list)))
}
