package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	hashSt := make(map[int]bool, len(nums))
	for _, k := range nums {
		hashSt[k] = true
	}

	cur := dummy

	for cur != nil && cur.Next != nil {
		if hashSt[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
