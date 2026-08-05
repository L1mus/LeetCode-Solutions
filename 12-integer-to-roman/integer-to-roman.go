func intToRoman(num int) string {
	symbol := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for idx, v := range symbol {
		if num == v {
			return idx
		}
	}

	var result string
	numStr := strconv.Itoa(num)
	satuan := len(numStr)
	var destructuringNum []string

	for idx, _ := range numStr {
		transform := string(numStr[idx])
		for i := 1; i < satuan; i++ {
			transform += "0"
		}
		if transform == "00" {
			satuan--
			continue
		} else if transform == "000" {
			satuan--
			continue
		}
		satuan--
		destructuringNum = append(destructuringNum, transform)
	}
	for _, v := range destructuringNum {
		value, _ := strconv.Atoi(v)
		if value%1000 == 0 {
			nilai := value / 1000
			for i := 0; i < nilai; i++ {
				result += "M"
				continue
			}
		} else if value%100 == 0 {
			for value != 0 {
				if value == 900 {
					result += "CM"
					value -= 900
					continue
				} else if value%500 >= 0 && value%500 != value {
					result += "D"
					value -= 500
				} else if value == 400 {
					result += "CD"
					value -= 400
					continue
				} else {
					result += "C"
					value -= 100
				}
				continue
			}
		} else if value%10 == 0 {
			for value != 0 {
				if value == 90 {
					result += "XC"
					value -= 90
					continue
				} else if value%50 >= 0 && value%50 != value {
					result += "L"
					value -= 50
				} else if value == 40 {
					result += "XL"
					value -= 40
					continue
				} else {
					result += "X"
					value -= 10
				}
				continue
			}
		} else {
			for value != 0 {
				if value == 9 {
					result += "IX"
					value -= 9
					continue
				} else if value%5 >= 0 && value%5 != value {
					result += "V"
					value -= 5
				} else if value == 4 {
					result += "IV"
					value -= 4
					continue
				} else {
					result += "I"
					value -= 1
				}
				continue
			}
		}
	}
	return result
}