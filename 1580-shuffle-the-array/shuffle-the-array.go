func shuffle(nums []int, n int) []int {
    newNums := make([]int,0,len(nums))
    left:= 0
    for right := n; right < len(nums); right++{
        newNums = append(newNums,nums[left],nums[right])
        left++
    }
    return newNums
}