func maxDistinct(s string) int {
	//constraint huruf kecil semua
	//return int maks substring
	//tidak ada substring yang di awali karakter yang sama
	var reserveCharacter []string
	var distinct int
	var slow int
	for fast := slow; fast < len(s); fast++ {
		if s[slow] != s[fast] && !slices.Contains(reserveCharacter, string(s[fast])) {
			reserveCharacter = append(reserveCharacter, string(s[slow]))
			slow = fast
			distinct++
			continue
		}
	}
	return distinct + 1
}