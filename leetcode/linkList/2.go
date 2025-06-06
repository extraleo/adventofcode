package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 尾插法, 需要一个 p0 节点
	p0 := &ListNode{}
	cur := p0
	carry := 0 // 进位
	for l1 != nil || l2 != nil || carry != 0{
		if l1 != nil{
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil{
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry%10}
		cur = cur.Next
		carry /= 10
	}
	return p0.Next
}
