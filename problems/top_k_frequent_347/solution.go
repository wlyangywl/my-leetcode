package top_k_frequent_347

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
func topKFrequent(nums []int, k int) []int {
	frequents := make(map[int]int)
	for _, num := range nums {
		if _, ok := frequents[num]; ok {
			frequents[num] = frequents[num] + 1
		} else {
			frequents[num] = 1
		}
	}
	// TODO 再了解一下大小堆
	return nums
	// heap.Pop()
}
