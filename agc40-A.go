// 10 November 2021
// AtCoder Grand Contest 040 - A
// link : https://atcoder.jp/contests/agc040/tasks/agc040_a

package main

import (
	"fmt"
)

func max(args ...int) int {
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func add(args ...int) (sum int) {
	for _, v := range args {
		sum += v
	}
	return sum
}

func main() {
	var S string
	fmt.Scan(&S)
	N := len(S) + 1
	a := make([]int, N)
	for i := 1; i < N; i++ {
		if S[i-1] == '<' {
			a[i] = a[i-1] + 1
		}
	}

	for i := N - 2; i >= 0; i-- {
		if S[i] == '>' {
			a[i] = max(a[i], a[i+1]+1)
		}
	}

	ans := add(a...)
	fmt.Println(ans)

}
