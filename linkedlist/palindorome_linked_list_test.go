package linkedlist


func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}
	slow, fast := head, head
	for fast!=nil && fast.Next != nil{
		slow, fast = slow.Next, fast.Next.Next
	}
	if fast != nil{
		slow = slow.Next
	}
	r:=reverseList(slow)
	for r != nil{
		if head.Val != r.Val{
			return false
		}
		r, head = r.Next, head.Next
	}
	return true
}