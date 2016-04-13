// https://www.hackerrank.com/challenges/angry-professor

package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%v", &t)
	for i := 0; i < t; i++ {
		var n, k int
		fmt.Scanf("%v %v", &n, &k)
		onTimeCount := 0
		for i := 0; i < n; i++ {
			var a int
			fmt.Scanf("%v", &a)
			if a <= 0 {
				onTimeCount++
			}
		}
		if onTimeCount >= k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
