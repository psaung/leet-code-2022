package main


func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	pre := fakeHead
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}
	return fakeHead.Next
}