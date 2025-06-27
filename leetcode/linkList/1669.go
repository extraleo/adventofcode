package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{Next: list1}
	cur := dummy
	for range a {
		cur = cur.Next
	}
	nxt := cur.Next
	cur.Next = list2
	for range b - a {
		nxt = nxt.Next
	}
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = nxt.Next
	return dummy.Next
}
