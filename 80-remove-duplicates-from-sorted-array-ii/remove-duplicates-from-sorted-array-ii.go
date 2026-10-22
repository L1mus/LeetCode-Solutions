func removeDuplicates(nums []int) int {
	duplicate := 1
	k := 1
	n := len(nums)

	if n < 2 {
		return n
	}

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			duplicate++
		} else {
			duplicate = 1
		}

		if duplicate <= 2 {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}