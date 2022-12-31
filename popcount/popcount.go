package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// n << x = n * 2^x
// y >> z = y / 2^z

// PopCount returns the population count (number of set bits) of x.
// func PopCount(x uint64) (results int) {
// 	for i := 0; i <= 8; i++ {
// 		results += int(pc[byte(x>>(i*8))])
// 	}
// 	return results
// }

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) (ret int) {
	for i := 0; i < 64; i++ {
		fmt.Println("232", x&1)
		if x&1 == 1 {
			ret++
		}
		x >>= 1

	}
	return ret
}
