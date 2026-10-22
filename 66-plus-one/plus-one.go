func plusOne(digits []int) []int {
	plus := 1
	n := len(digits) - 1
	for i := n; i >= 0; i-- {
		result := digits[i] + plus
		digits[i] = result % 10
		plus = result / 10
	}
	if plus > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}