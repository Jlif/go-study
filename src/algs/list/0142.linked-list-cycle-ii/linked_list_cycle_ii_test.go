package _142_linked_list_cycle_ii

import (
	"fmt"
	"tech.jiangchen/go-study/src/algs/data_structure"
	"testing"
)

func TestHasCycleII(t *testing.T) {
	list1 := data_structure.NewListNodeFromArr([]int{1, 4, 7})
	fmt.Printf("%+v\n", detectCycle(list1))

	list2 := data_structure.NewListNodeFromArr([]int{})
	fmt.Printf("%+v\n", detectCycle(list2))

	tmp := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	list3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: tmp,
			},
		},
	}
	tmp.Next.Next = list3.Next.Next
	fmt.Printf("%+v\n", detectCycle(list3))
}
