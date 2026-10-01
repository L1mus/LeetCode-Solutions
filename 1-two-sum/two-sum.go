func twoSum(nums []int, target int) []int {
    for idx,v := range nums{
        for i := idx+1 ; i < len(nums);i++ {
            result := v + nums[i] 
            if result == target{
                return []int{idx,i}
            }
        }
    }
    return []int{}
}