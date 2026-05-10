package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // no incial "falso"
	current := dummy
	carry := 0 // vai um

	for l1 != nil || l2 != nil || carry != 0 { // equanto tiver valor para somar
		sum := carry

		if l1 != nil { // soma valores da lista l1 até quando tiver vazio
			sum += l1.Val
			l1 = l1.Next // proximo valor
		}

		if l2 != nil { // soma valores da lista l2 até quanto tiver vazio
			sum += l2.Val
			l2 = l2.Next // proximo valor
		}

		digit := sum % 10
		carry = sum / 10

		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummy.Next
}
