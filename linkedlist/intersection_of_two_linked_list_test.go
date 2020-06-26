package linkedlist

// leetcode 160

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var i, j int
	head1, head2 := headA, headB
	for headA.Next != nil {
		i++
		headA = headA.Next
	}
	for headB.Next != nil {
		j++
		headB = headB.Next
	}
	if i > j {
		i, j = j, i
		head1, head2 = head2, head1
	}
	for n := 0; n < j-i; n++ {
		head2 = head2.Next
	}
	for head1 != nil && head2 != nil && head1 != head2 {
		head1, head2 = head1.Next, head2.Next
	}
	return head1
}
