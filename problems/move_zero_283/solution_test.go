package move_zero_283

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int // ✅ 添加期望结果
	}{
		{
			name: "move zero 1",
			args: args{
				nums: []int{1, 2, 0, 2, 1, 0, 0},
			},
			want: []int{1, 2, 2, 1, 0, 0, 0},
		},
		{
			name: "move zero 2",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "move zero 3",
			args: args{
				nums: []int{0, 1},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() got = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
