func twoSum(nums []int, target int) []int {
	maping := make(map[int]int)

    for idx,v := range nums {
        calculate := target - v
        i,ok := maping[calculate]
        if ok {
            return []int{i,idx}
        }
        maping[v]=idx
    }
   return nil
}