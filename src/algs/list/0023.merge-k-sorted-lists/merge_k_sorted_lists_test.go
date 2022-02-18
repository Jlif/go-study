package _023_merge_k_sorted_lists

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := data_structure.NewListNodeFromArr([]int{1, 4, 7})
	list2 := data_structure.NewListNodeFromArr([]int{2, 5, 8})
	list3 := data_structure.NewListNodeFromArr([]int{3, 6, 9})
	lists := []ListNode{*list1, *list2, *list3}
	mergedList := mergeKSortedLists(lists)
	println(data_structure.GetListNodeStr(mergedList))
}
