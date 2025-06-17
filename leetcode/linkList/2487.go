package main

// 递推来写
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
			return head
	}
	node := removeNodes(head.Next)
	if node.Val > head.Val{
			return node
	}
	head.Next = node
	return head
}

// 反转 -> 迭代 -> 反转

func removeNodesII(head *ListNode) *ListNode {
	h1 := reverse2487(head)
	cur := h1
	for cur.Next != nil{
		if cur.Next.Val < cur.Val{
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return reverse2487(h1)
}

func reverse2487(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil{
		nxt := cur.Next
		cur.Next = pre
		pre=cur
		cur = nxt
	}
	return pre
}