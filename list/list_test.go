package list

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		list *ListNode
		pos  int
	}{
		{genList([]int{1, 2, 3, 4, 5}), 2},
		{genList([]int{1}), 1},
		{genList([]int{1, 2}), 2},
	}
	for _, c := range cases {
		printList(c.list)
		printList(removeNthFromEnd(c.list, c.pos))
	}
}

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		list *ListNode
	}{
		{genList([]int{1, 2, 3, 4, 5})},
		{genList([]int{1, 2})},
	}

	for _, c := range cases {
		printList(c.list)
		printList(middleNode(c.list))
	}
}

func TestHhasCycle(t *testing.T) {
	cases := []*ListNode{
		genCycleList([]int{3, 2, 0, -4}, 1),
	}
	for _, c := range cases {
		fmt.Println(hasCycle(c))
	}
}

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		list1, list2 *ListNode
	}{
		{genList([]int{1, 2, 3}), genList([]int{1, 3, 4})},
	}
	for _, c := range cases {
		printList(c.list1)
		printList(c.list2)
		printList(mergeTwoLists(c.list1, c.list2))
	}
}

func TestReverseList(t *testing.T) {
	cases := []struct {
		list *ListNode
	}{
		{genList([]int{1, 2, 3, 4, 5})},
	}
	for _, c := range cases {
		printList(c.list)
		printList(reverseList(c.list))
	}
}
