package main

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := data_structure.NewListNode([]int{1, 3, 5, 7})
	list2 := data_structure.NewListNode([]int{2, 4, 6, 8})
	data_structure.PrintListNode(mergeTwoLists(list1, list2))
}
