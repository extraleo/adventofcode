package hot100

type ListNode struct {
     Val int
     Next *ListNode
 }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	q, p := headA, headB
	for p != q {
			if p != nil {
					p = p.Next
			}else{
					p = headB
			}

			if q != nil{
					q = q.Next
			}else{
					q = headA
			}
	}
	return p
}
