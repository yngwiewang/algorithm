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

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	var endA, endB *ListNode 
	for h := headA; h != nil; h = h.Next {
		endA = h
		lenA++
	}
	for h := headB; h != nil; h = h.Next {
		endB = h
		lenB++
	}
	if endA != endB{
		return nil
	}
    
	head1, head2 := headA, headB
	if lenA < lenB {
		lenA, lenB = lenB, lenA
		head1, head2 = head2, head1
	}
	for i := 0; i < lenA - lenB; i++ {
		head1 = head1.Next
	}
	for {
		if head1 == head2 {
			return head1
		}
		head1, head2 = head1.Next, head2.Next
	}
}
