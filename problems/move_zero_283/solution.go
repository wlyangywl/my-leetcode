package move_zero_283

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。

func moveZeroes(nums []int) {
	var j int
	for i, num := range nums {
		if num != 0 {
			// 和pos位置做交换
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			j++
		}
	}
}
