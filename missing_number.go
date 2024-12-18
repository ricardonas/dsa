func missingNumber(nums []int) int {
    n := len(nums)
    sum := 0
   
    for i := 0; i <= len(nums) - 1; i++ {
        sum += nums[i]
    }

    missingSum := ((n * (n + 1)) / 2)

    return missingSum - sum

}
