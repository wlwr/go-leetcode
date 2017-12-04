package add_two_numbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := []int{2, 4}
	l2 := []int{5, 6, 4}
	result := []int{7, 0, 5}
	res := addTwoNumbers(generateLinkedList(l1), generateLinkedList(l2))
	if resultList := generateLinkedList(result); !equal(res, resultList) {
		t.Fatalf("%v\n", resultList)
	}
}

func equal(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func generateLinkedList(numbers []int) *ListNode {
	head := new(ListNode)
	walk := head
	for index, nodeVal := range numbers {
		walk.Val = nodeVal
		if index < len(numbers)-1 {
			walk.Next = new(ListNode)
			walk = walk.Next
		}
	}
	return head
}
