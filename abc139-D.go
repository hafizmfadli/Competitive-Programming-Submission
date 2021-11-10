// 10 November 2021
// AtCoder Beginner Contest 139 - D
// link : https://atcoder.jp/contests/abc139/tasks/abc139_d

package main

import "fmt"

func main() {
	var N int64
	fmt.Scan(&N)
	ans := (N - 1) * N / 2
	fmt.Println(ans)
}
