// https://www.hackerrank.com/challenges/simple-array-sum

package main
import "fmt"

func main() {
    var n, currentNumber, sum int
    fmt.Scan(&n)
    for i:=0; i<n; i++ {
        fmt.Scan(&currentNumber)
        sum += currentNumber
    }
    fmt.Print(sum)
}
