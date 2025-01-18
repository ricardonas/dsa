func distinctAverages(nums []int) int {
    sort.Ints(nums)
	l, r := 0, len(nums)-1

	hasher := make(map[float64]*struct{})

	for l < r {

		avg := float64(nums[l] + nums[r]) / 2.0

		if _, ok := hasher[avg]; !ok {
			hasher[avg] = nil
		}
		
		l++
		r--
	}

	return len(hasher)
}
