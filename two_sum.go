func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if idx, exists := numMap[diff]; exists {
			return []int{idx, i}
		}
		numMap[v] = i
	}

	return nil
}
