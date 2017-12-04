package add_two_numbers

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	walk := head
	carry := 0
	for l1 != nil || l2 != nil {
		var l1Val, l2Val, val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + carry
		if sum >= 10 {
			carry = 1
			val = sum - 10
		} else {
			carry = 0
			val = sum
		}
		walk.Val = val
		if l1 != nil || l2 != nil {
			walk.Next = new(ListNode)
			walk = walk.Next
		}
	}
	if carry > 0 {
		walk.Next = new(ListNode)
		walk.Next.Val = carry
	}
	return head
}
