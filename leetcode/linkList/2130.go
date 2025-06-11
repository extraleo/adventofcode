package main

func pairSum(head *ListNode) int {
	mid := findMid(head)
	sec := reverseLink(mid)
	ans := 0
	for sec != nil {
		ans = max(ans, head.Val+sec.Val)
		head = head.Next
		sec = sec.Next
	}
	return ans
}