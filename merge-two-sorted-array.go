func merge(nums1 []int, m int, nums2 []int, n int) {
	writeIndex := (m + n) - 1

	for m > 0 && n > 0 {
		num1, num2 := nums1[m-1], nums2[n-1]

		if num1 > num2 {
			nums1[writeIndex] = num1
			m -= 1
		} else {
			nums1[writeIndex] = num2
			n -= 1
		}
		writeIndex -= 1
	}

	for n > 0 {
		nums1[writeIndex] = nums2[n-1]
		writeIndex, n = writeIndex-1, n-1
	}

}
