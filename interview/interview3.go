package main

import "fmt"

func interview3(s string) string {
	ss := []rune(s)

	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	fmt.Println(string(ss))
	return string(ss)
}
