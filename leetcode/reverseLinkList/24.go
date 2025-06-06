package main

func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    p0 := dummy
    n := 0
    for cur:=head; cur != nil; cur = cur.Next{
        n++
    }
    var pre, cur *ListNode = nil, p0.Next
    for ;n >= 2; n -= 2{
        for range(2){
            next := cur.Next
            cur.Next = pre
            pre = cur
            cur = next
        }
        nxt := p0.Next
        p0.Next.Next = cur
        p0.Next = pre
        p0 = nxt
    }
    return dummy.Next
}