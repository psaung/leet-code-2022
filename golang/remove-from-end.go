package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := &ListNode{}
	node.Next = head

	slow, fast := node, node

	for ; n > 0; n-- {
		fast = fast.Next
	}

	for ; fast.Next != nil; slow, fast = slow.Next, fast.Next {}

	slow.Next = slow.Next.Next

	return node.Next
}