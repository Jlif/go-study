package _092_reverse_linked_list_ii

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestReverseLinkedListII(t *testing.T) {
	list := data_structure.NewListNodeFromArr([]int{1, 4, 7, 10})
	println(data_structure.GetListNodeStr(list))

	list = reverseN(list, 2)
	println(data_structure.GetListNodeStr(list))

	list = reverseBetween(list, 2, 3)
	println(data_structure.GetListNodeStr(list))
}
