// https://www.hackerrank.com/challenges/find-digits

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%v", &t)
	for i := 0; i < t; i++ {
		var n int
		var divisibleDigits int = 0
		fmt.Scanf("%v", &n)
		for i := n; i > 0; i /= 10 {
			digit := i % 10
			if (digit != 0 && n % digit == 0) {
				divisibleDigits++
			}
		}
		fmt.Println(divisibleDigits)
	}
}
