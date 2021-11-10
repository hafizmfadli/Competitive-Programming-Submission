// 10 November 2021
// AtCoder Beginner COntest 066
// link : https://atcoder.jp/contests/abc066/tasks/abc066_b

package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	N := len(S) - 2
	ans := 0
	if N > 0 {
		for i := N/2 - 1; i >= 0; i-- {
			ok := false
			for j := 0; j <= i; j++ {
				if S[j] != S[j+i+1] {
					break
				}
				if j == i {
					ok = true
					ans = (i + 1) * 2
					break
				}
			}
			if ok {
				break
			}
		}
	}
	fmt.Println(ans)
}
