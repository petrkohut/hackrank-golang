// https://www.hackerrank.com/challenges/utopian-tree

package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
	cycles := make([]int, t)
    for i := 0; i < t; i++ {
        fmt.Scan(&cycles[i])
    }
	for i := 0; i < t; i++ {
		result := 1
		for j := 0; j < cycles[i]; j++ {
			if ((j + 2) % 2) == 0 {
				result = result * 2
			} else {
				result++
			}
		}
		fmt.Println(result)
	}
}
