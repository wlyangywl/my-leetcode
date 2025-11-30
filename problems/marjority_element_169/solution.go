package marjority_element_169

func majorityElement(nums []int) int {
	var count int
	var major int
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if num == major {
			count++
		} else {
			count--
		}
	}
	return major
}
