package main

import "fmt"

func canChange(start string, target string) bool {
	waitingL, waitingR := 0, 0

	for i := 0; i < len(start); i++ {
		curr := string(start[i])
		goal := string(target[i])

		if curr == "R" {
			if waitingL > 0 {
				return false
			}
			waitingR++
		}

		if goal == "L" {
			if waitingR > 0 {
				return false
			}
			waitingL++
		}

		if goal == "R" {
			if waitingR == 0 {
				return false
			}
			waitingR--
		}
		if curr == "L" {
			if waitingL == 0 {
				return false
			}
			waitingL--
		}
	}
	return waitingL == 0 && waitingR == 0
}

func main() {
	fmt.Println(canChange("_L__R__R_", "L______RR"))
}
