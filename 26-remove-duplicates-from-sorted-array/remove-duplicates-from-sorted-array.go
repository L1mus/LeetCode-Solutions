func removeDuplicates(nums []int) int {
	n := len(nums)
	k := 0
	duplicate := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, existing := duplicate[nums[i]]; existing {
			continue
		} else {
			nums[k] = nums[i]
			k++
		}
		duplicate[nums[i]] = i
	}
	return k
}