func moveZeroes(nums []int) {
	l := 0
	for r := l + 1; r < len(nums); r++ {
		if nums[l] != 0 {
			l++
		} else if nums[l] == 0 && nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		} else {
			continue
		}
	}
}
