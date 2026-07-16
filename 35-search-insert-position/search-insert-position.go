func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		findNumber := int(math.Floor(float64(left+right) / 2))
		if nums[findNumber] >= target {
			right = findNumber
		} else {
			left = findNumber + 1
		}
	}
    if nums[left] == target{
        return left
    }else if nums[left] < target{
        return left + 1
    }else if nums[left] > target{
        return left
    }
	return 0
}