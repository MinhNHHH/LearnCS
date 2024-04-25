package main

// Input: chars = ["a","a","b","b","c","c","c"]
// Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
// Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

func compress(chars []byte) int {
	index := 0
	i := 0

	for i < len(chars) {
		count := 1
		for i < len(chars) && chars[index] == chars[i] {
			count++
			i++
		}
		chars[index] = uint8(count)
		count = 1
	}
	return 1
}

func main() {
}
