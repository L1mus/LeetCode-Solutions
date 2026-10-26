func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    //
    combined := append(nums1,nums2...)
    slices.Sort(combined)
    var median float64
    if len(combined) % 2 == 0 {
        median = float64((combined[(len(combined)/2)-1] + combined[len(combined)/2]))/2
    }else{
        median = float64(combined[len(combined)/2])
    }

    return median
}