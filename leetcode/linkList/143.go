package main


func reorderList(head *ListNode)  {
    mid := _midNode(head)
		head2 := _reverseList(mid)
		for head2.Next != nil{
			nxt := head.Next
			nxt2 := head2.Next
			head.Next = head2
			head2.Next = nxt
			head = nxt
			head2 = nxt2

		}

}


func _midNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func _reverseList(head *ListNode) *ListNode{
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}