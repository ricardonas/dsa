func merge(nums1 []int, m int, nums2 []int, n int) {
	pivot := (m + n) - 1

	for m > 0 && n > 0 {
		num1, num2 := nums1[m-1], nums2[n-1]

		if num1 > num2 {
			nums1[pivot] = num1
			m -= 1
		} else {
			nums1[pivot] = num2
			n -= 1
		}
		pivot -= 1
	}

	for n > 0 {
		nums1[pivot] = nums2[n-1]
		pivot, n = pivot-1, n-1
	}

}
