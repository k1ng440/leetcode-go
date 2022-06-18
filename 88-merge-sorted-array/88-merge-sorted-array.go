package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Do not return anything, modify nums1 in-place instead

	pointer1 := m - 1
	pointer2 := n - 1

	for pointer1 >= 0 && pointer2 >= 0 {
		index := pointer1 + pointer2 + 1

		if nums1[pointer1] >= nums2[pointer2] {
			nums1[index] = nums1[pointer1]
			pointer1--
			continue
		}

		nums1[index] = nums2[pointer2]
		pointer2--
	}

	for pointer2 >= 0 {
		nums1[pointer2] = nums2[pointer2]
		pointer2--
	}
}
