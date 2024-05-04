package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isVowel(s byte) bool {
	switch string(s) {
	case "a":
		return true
	case "e":
		return true
	case "o":
		return true
	case "i":
		return true
	case "u":
		return true
	}
	return false
}
func maxVowels(s string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	ans := count
	for j := k; j < len(s); j++ {
		if isVowel(s[j]) {
			count++
		}
		if isVowel(s[j-k]) {
			count--
		}
		ans = max(ans, count)
	}

	return ans
}

func main() {}
