package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// head = [1,2,3,4,5], n = 2
	var dummy = &ListNode{Val: -1, Next: head}
	var fast, slow = head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

// https://leetcode.cn/problems/middle-of-the-linked-list/description/
func middleNode(head *ListNode) *ListNode {
	// 1,2,3,4,5
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// https://leetcode.cn/problems/linked-list-cycle/description/
func hasCycle(head *ListNode) bool {
	// head = [3,2,0,-4], pos = 1
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// l1 = [1,2,4], l2 = [1,3,4]
	dummy := &ListNode{Val: -1, Next: nil}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next //base and break origin link
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

// n + m = m + n
// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}
		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}

// head insert
// https://leetcode.cn/problems/reverse-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for ; head != nil; head = head.Next {
		tail = &ListNode{Val: head.Val, Next: tail}
	}
	return tail
}

func genList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	nodes := make([]*ListNode, n)
	for i, v := range nums {
		nodes[i] = &ListNode{v, nil}
	}
	for i := 0; i < n-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	return nodes[0]
}

func genCycleList(nums []int, pos int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	nodes := make([]*ListNode, n)
	for i, v := range nums {
		nodes[i] = &ListNode{v, nil}
	}
	for i := 0; i < n-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos > 0 && pos < n-1 {
		nodes[n-1].Next = nodes[pos]
	}
	return nodes[0]
}

func printList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Print(node.Val)
	}
	fmt.Println()
}
