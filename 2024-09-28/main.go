package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// q1()
	// q2()
	// q3()
	q4()
}

func q1() {
	s := make([]string, 12)
	for i := 0; i < 12; i++ {
		fmt.Scan(&s[i])
	}

	ans := 0
	for i := 0; i < 12; i++ {
		if i+1 == len(s[i]) {
			ans++
		}
	}

	fmt.Println(ans)
}

func q2() {
	var s string
	fmt.Scan(&s)

	location := make(map[rune]int)
	for i, r := range s {
		location[r] = i
	}

	target := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	cur := 'A'
	ans := 0.0
	for _, r := range target {
		if r == cur {
			continue
		} else {
			ans += math.Abs(float64(location[r] - location[cur]))
			cur = r
		}
	}

	fmt.Println(int(ans))
}

func q3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		return num
	}

	n := readInt()

	maxa, maxb := -1<<31, -1<<31
	for i := 0; i < n; i++ {
		a := readInt()
		if a > maxa {
			maxa = a
		}
	}

	for i := 0; i < n; i++ {
		b := readInt()
		if b > maxb {
			maxb = b
		}
	}

	fmt.Println(maxa + maxb)
}

func q4() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		return num
	}

	n, m := readInt(), readInt()

	dist := make([]int64, n)
	graph := make([][]struct{ v, w int }, n)
	for i := 0; i < m; i++ {
		u, v, w := readInt()-1, readInt()-1, readInt()
		graph[u] = append(graph[u], struct{ v, w int }{v, w})
		graph[v] = append(graph[v], struct{ v, w int }{u, -w})
	}

	isVisited := make([]bool, n)
	queue := make([]int, 0, n)

	bfs := func(s int) {
		if isVisited[s] {
			return
		}
		queue = queue[:0] // 重置队列
		queue = append(queue, s)
		isVisited[s] = true
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, edge := range graph[u] {
				if !isVisited[edge.v] {
					queue = append(queue, edge.v)
					isVisited[edge.v] = true
					dist[edge.v] = dist[u] + int64(edge.w)
				}
			}
		}
	}

	for s := 0; s < n; s++ {
		bfs(s)
	}

	for _, v := range dist {
		fmt.Print(v)
		fmt.Print(" ")
	}
}

func q5(){
	
}