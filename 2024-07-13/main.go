package main
import (
	"fmt"
	"math"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	q3()
}

func min(a, b int)int{
	if a>b{
		return b
	}else{
		return a
	}
}

func q1(){
	var r, g, b int
	fmt.Scan(&r, &g, &b)
	var s string
	fmt.Scan(&s)

	ans:=0
	switch s{
	case "Blue": ans=min(r, g)
	case "Red": ans=min(g, b)
	case "Green": ans=min(r, b)
	}
	fmt.Println(ans)
}

func q2(){
	var xa, ya, xb, yb, xc, yc int
	fmt.Scan(&xa, &ya, &xb, &yb, &xc, &yc)

	innerp:=func(x1, y1, x2, y2 int)int{
		return x1*x2+y1*y2
	}

	a:=innerp(xa-xb, ya-yb, xa-xc, ya-yc)
	b:=innerp(xa-xc, ya-yc, xb-xc, yb-yc)
	c:=innerp(xa-xb, ya-yb, xb-xc, yb-yc)

	if a*b*c==0{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}

func q3(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	
	intervals := make([][2]int, N)
	sumL, sumR := 0, 0
	
	for i := 0; i < N; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		L, _ := strconv.Atoi(parts[0])
		R, _ := strconv.Atoi(parts[1])
		intervals[i] = [2]int{L, R}
		sumL += L
		sumR += R
	}
	
	if sumL > 0 || sumR < 0 {
		fmt.Println("No")
		return
	}
	
	result := make([]int, N)
	currentSum := 0
	
	for i, interval := range intervals {
		L, R := interval[0], interval[1]
		if currentSum<0{
			if -currentSum<R{
				result[i]=-currentSum
				currentSum=0
			}else{
				currentSum+=R
				result[i]=R
			}
		}else if currentSum>0{
			if -currentSum>L{
				result[i]=-currentSum
				currentSum=0
			}else{
				currentSum+=L
				result[i]=L
			}
		}else if currentSum==0{
			if L>=0{
				result[i]=L
			}else if R<=0{
				result[i]=R
			}else{
				result[i]=0
			}
			currentSum+=result[i]
		}
	}
	
	fmt.Println("Yes")
	for _, x := range result {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}

type vertex struct{
	id uint
	w uint64
	d uint64
	prev int
}

func q4(){
	const inf uint64 = math.MaxInt64
	var n, m uint
	fmt.Scan(&n, &m)
	v:=make([]vertex, n)
	for i:=uint(0);i<n;i++{
		fmt.Scan(&v[i].w)
		v[i].id=i
		v[i].prev=-1
		v[i].d=inf
	}
	graph:=make([][][]uint, n)
	for i:=uint(0);i<m;i++{
		var t1, t2 uint
		var w uint
		fmt.Scan(&t1, &t2, &w)
		temp:=[]uint{t1, w}
		graph[i]=append(graph[i], temp)
	}

	dijkstra:=func(id int){
		visited:=make([]bool, n)
		visited[0]=true
		queue:=make([]vertex, 0)

		queue=append(queue, v[0])
		for queue!=nil{
			for i:=range graph[v[0].id]{
				queue = append(queue, v[graph[v[0].id][i][0]])
			}
		}
	}
	dijkstra(1)
}