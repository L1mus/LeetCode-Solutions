func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for idx, v := range nums {
		requireNumber := target - v
		if _, ok := mapNums[requireNumber]; ok {
			return []int{mapNums[requireNumber], idx}
		}
		mapNums[v] = idx
	}
	return []int{}
}