/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList1(head *ListNode) *ListNode {
	 if head == nil || head.Next == nil{
		 return head
	 }
	 newHead:=reverseList(head.Next)
	 head.Next.Next = head
	 head.Next = nil
	 return newHead
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		nxt := curr.Next
		curr.Next = pre
		pre = curr
		curr = nxt
	}
	
	return pre
}
// @lc code=end

