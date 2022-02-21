package _141_linked_list_cycle

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestHasCycle(t *testing.T) {
	list1 := data_structure.NewListNodeFromArr([]int{1, 4, 7})
	println(hasCycle(list1))

	list2 := data_structure.NewListNodeFromArr([]int{})
	println(hasCycle(list2))

	tmp := &ListNode{
		Val:  3,
		Next: nil,
	}
	list3 := &ListNode{
		Val:  0,
		Next: tmp,
	}
	tmp.Next = list3
	println(hasCycle(list3))

}
