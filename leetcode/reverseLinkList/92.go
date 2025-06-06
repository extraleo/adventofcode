package main

import "fmt"


func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	var pre, cur *ListNode = nil, dummy
	for range left {
		cur = cur.Next
	}
    fmt.Println(cur.Val)
	for range(right-left){
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	dummy.Next.Next = cur
	dummy.Next = pre

	return dummy.Next
}


func main(){
	hea
	for i:=range 5 {
		
	}
	reverseBetween()
}