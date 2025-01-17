// 3194. Minimum Average of Smallest and Largest Elements
func minimumAverage(nums []int) float64 {
	sort.Ints(nums)

	left, right := 0, len(nums)-1

	min := math.MaxFloat64

	for left <= right {
		calc := float64(nums[left]+nums[right]) / float64(2)

		if calc < min {
			min = calc
		}
		
		left++
		right--
	}

	return min
}
