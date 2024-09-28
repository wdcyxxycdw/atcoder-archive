package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, m)
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&x[i][j])
		}
	}

	for i := 0; i < m; i++ {
		temp := 0
		for j := 0; j < n; j++ {
			temp += x[j][i]
		}
		fmt.Println(temp)
		if temp < a[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
