package leetcode0019

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	leftNode, rightNode := head, head
	rightIndex := 1

	for {

		// rightNode指向了最后一个值，leftNode 该删除了
		if rightNode.Next == nil {
			if rightIndex > n {
				leftNode.Next = rightNode
			}
			break
		}

		if rightIndex > n {
			leftNode = leftNode.Next
		}
		rightNode = rightNode.Next
		rightIndex++
	}

	return head
}
