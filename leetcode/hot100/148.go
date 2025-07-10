package hot100

func findMid(root *ListNode) *ListNode {
	fast, slow := root, root
	pre := root
	for fast != nil && fast.Next != nil {
		pre = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	pre.Next = nil
	return slow
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	cur := &ListNode{}
	dummy := cur

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	head2 := findMid(head)
	head = sortList(head)
	head2 = sortList(head2)

	return mergeTwoList(head, head2)
}
