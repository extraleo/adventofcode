package main

func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil && cur.Next != nil; {
		if cur.Val != cur.Next.Val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}

	}
	return head
}
