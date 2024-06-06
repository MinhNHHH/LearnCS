// We are playing the Guess Game. The game is as follows:

// I pick a number from 1 to n. You have to guess which number I picked.

// Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

// You call a pre-defined API int guess(int num), which returns three possible results:

// -1: Your guess is higher than the number I picked (i.e. num > pick).
// 1: Your guess is lower than the number I picked (i.e. num < pick).
// 0: your guess is equal to the number I picked (i.e. num == pick).

// Return the number that I picked.

// Input: n = 10, pick = 6
// Output: 6

// Input: n = 1, pick = 1
// Output: 1

package main

import "fmt"

func main() {
	fmt.Println(guessNumber(10))
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num, picked int) int {
	if num > picked {
		return -1
	} else if num < picked {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	l, r := 1, n
	for {
		mid := (l + r) / 2
		res := guess(mid, 6)
		if res == 1 {
			l = mid + 1
		} else if res == -1 {
			r = mid - 1
		} else {
			return mid
		}
	}
}
