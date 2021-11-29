package main

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := data_structure.NewListNode([]int{1, 3, 5, 7})
	list2 := data_structure.NewListNode([]int{2, 4, 6, 8})
	mergedList := mergeTwoSortedLists(list1, list2)
	ans := data_structure.NewListNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
	println(data_structure.GetListNodeStr(mergedList))
	println(data_structure.GetListNodeStr(ans))
	if data_structure.GetListNodeStr(ans) != data_structure.GetListNodeStr(mergedList) {
		t.Errorf("expect %s, but %s got", data_structure.GetListNodeStr(ans), data_structure.GetListNodeStr(mergedList))
	}
}
