// 167. Two Sum II - Input Array Is Sorted - Two pointers
func twoSum(nums []int, target int) []int {
	res := []int{}
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			res = append(res, left+1, right+1)
			break
		}

		if sum > target {
			right -= 1
		} else {
			left += 1
		}
	}

	return res
}
