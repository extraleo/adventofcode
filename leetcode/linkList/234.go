package main

func isPalindrome(head *ListNode) bool {
	mid := findMid(head)
	sec := reverseLink(mid)
	for sec != nil {
		if head.Val != sec.Val {
			return false
		}
		head = head.Next
		sec = sec.Next
	}
	return true
}
func findMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseLink(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
