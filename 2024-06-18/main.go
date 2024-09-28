package main

import(
	"fmt"
)

func main(){
	var n, m int
	fmt.Scan(&n, &m)
	h:=make([]int, n)
	for i:=range h{
		fmt.Scan(&h[i])
	}

	for i, v:=range h{
		if m<v{
			fmt.Println(i)
			break
		}
		m-=v
	}
}