// https://www.hackerrank.com/challenges/sherlock-and-the-beast

package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%v", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%v", &n)
		fmt.Println(largestDecentNumber(n))
	}
}

func largestDecentNumber(n int) string {
	for i := n; i >= 0; i-- {
		if ((i % 3 == 0) && ((n - i) % 5 == 0)) {
			return strings.Repeat("5", i) + strings.Repeat("3", n - i)
		}
	}
	return "-1"
}
