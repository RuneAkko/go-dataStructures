package main

import "fmt"

func main() {
	test := func(a, b string) bool {
		n, m := len(a), len(b)
		for i, j := 0, 0; i < n && j < m; i, j = i+1, j+1 {
			if a[i] != b[j] {
				return a[i] > b[j]
			}
			if i+1 >= n {
				return true
			}
			if j+1 >= m {
				return false
			}
		}
		return true
	}
	fmt.Println(test("3", "30"))
}
