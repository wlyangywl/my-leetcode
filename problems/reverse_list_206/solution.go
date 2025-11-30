package reverse_list_206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
* 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解题思路： 核心定义三指针
// cur指向待反转的链表，同时兼顾当前反转的值
// pre指向已经反转的链表
// next指向待反转的链表首节点的下一个节点

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
