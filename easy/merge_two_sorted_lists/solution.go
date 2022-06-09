package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy ListNode
	tail := &dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}

func main() {
	// list 1
	ListOneNode3 := ListNode{Val: 4, Next: nil}
	ListOneNode2 := ListNode{Val: 2, Next: &ListOneNode3}
	ListOneNode1 := ListNode{Val: 1, Next: &ListOneNode2}

	// list 2
	ListTwoNode3 := ListNode{Val: 4, Next: nil}
	ListTwoNode2 := ListNode{Val: 3, Next: &ListTwoNode3}
	ListTwoNode1 := ListNode{Val: 1, Next: &ListTwoNode2}

	mergeTwoLists(&ListOneNode1, &ListTwoNode1)
}
