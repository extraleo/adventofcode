package main

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		nxt := cur.Next.Val
		if cur.Next.Next.Val == nxt {
			for cur.Next != nil && cur.Next.Val == nxt {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
