package merge_two_list_21

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	cur1 := list1
	cur2 := list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur.Next = &ListNode{
				Val: cur1.Val,
			}
			cur1 = cur1.Next
		} else if cur2.Val < cur1.Val {
			cur.Next = &ListNode{
				Val: cur2.Val,
			}
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	return res.Next
}
