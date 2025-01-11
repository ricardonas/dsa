import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for idx, value := range nums {
		// Skip a value if it has already been checked.
		if idx > 0 && value == nums[idx-1] {
			continue
		}

		l, r := idx+1, len(nums)-1

		for l < r {
			sum := value + nums[l] + nums[r]

			if sum < 0 {
				l += 1
			} else if sum > 0 {
				r -= 1
			} else {
				res = append(res, []int{value, nums[l], nums[r]})
				l += 1

				// Skip values that have already been checked.
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}

			}

		}

	}

	return res
}
