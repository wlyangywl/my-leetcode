package palindrome_list_234

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "odd length palindrome",
			args: args{
				head: func() *ListNode {
					// 1->2->3->2->1
					node1 := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 2}
					node5 := &ListNode{Val: 1}
					node4.Next = node5
					node3.Next = node4
					node2.Next = node3
					node1.Next = node2
					return node1
				}(),
			},
			want: true,
		},
		{
			name: "odd length not palindrome",
			args: args{
				head: func() *ListNode {
					// 1->2->3->2->1
					node1 := &ListNode{Val: 1}
					node2 := &ListNode{Val: 2}
					node3 := &ListNode{Val: 3}
					node4 := &ListNode{Val: 4}
					node5 := &ListNode{Val: 1}
					node4.Next = node5
					node3.Next = node4
					node2.Next = node3
					node1.Next = node2
					return node1
				}(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
