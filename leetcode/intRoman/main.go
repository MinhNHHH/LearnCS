package main

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	i := 0

	for num > 0 {
		for num >= values[i] {
			res += symbols[i]
			num -= values[i]
		}
		i++
	}

	return res
}

func main() {

}
