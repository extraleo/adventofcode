package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	p0 := dummy
	for range left - 1 {
		p0 = p0.Next
	}
	var pre, cur *ListNode = nil, p0.Next

	for range right - left + 1 {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	p0.Next.Next = cur
	p0.Next = pre

	return dummy.Next
}
