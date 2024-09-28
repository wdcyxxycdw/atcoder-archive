package main

import (
	"fmt"
	"sort"
)

type pos struct {
	x, y, z int
}

func (p pos) reset(x1, y1, z1 int) {
	p.x -= x1
	p.y -= y1
	p.z -= z1
}

func main() {
	q3()
}

func q2() {
	var a1, a2, b1, b2 pos
	fmt.Scan(&a1.x, &a1.y, &a1.z)
	fmt.Scan(&a2.x, &a2.y, &a2.z)
	fmt.Scan(&b1.x, &b1.y, &b1.z)
	fmt.Scan(&b2.x, &b2.y, &b2.z)

	a1.reset(a1.x, a1.y, a1.z)
	a2.reset(a1.x, a1.y, a1.z)
	b1.reset(a1.x, a1.y, a1.z)
	b2.reset(a1.x, a1.y, a1.z)

	xl, xr, yl, yr, zb, zt := b1.x, b2.x, b1.y, b2.y, b1.z, b2.z

	b3 := pos{
		x: xl,
		y: yl,
		z: zt,
	}
	b4 := pos{
		x: xl,
		y: yr,
		z: zb,
	}
	b5 := pos{
		x: xl,
		y: yr,
		z: zt,
	}
	b6 := pos{
		x: xr,
		y: yl,
		z: zb,
	}
	b7 := pos{
		x: xr,
		y: yl,
		z: zt,
	}
	b8 := pos{
		x: xr,
		y: yr,
		z: zb,
	}

	isInCube := func(a, target pos) bool {
		min := func(a, b int) int {
			if a > b {
				return b
			} else {
				return a
			}
		}
		max := func(a, b int) int {
			if a < b {
				return b
			} else {
				return a
			}
		}
		if a.x > min(0, target.x) && a.x < max(0, target.x) &&
			a.y > min(0, target.y) && a.y < max(0, target.y) &&
			a.z > min(0, target.z) && a.z < max(0, target.z) {
			return true
		} else {
			return false
		}
	}

	ans := isInCube(b1, a2) || isInCube(b2, a2) || isInCube(b3, a2) || isInCube(b4, a2) || isInCube(b5, a2) || isInCube(b6, a2) || isInCube(b7, a2) || isInCube(b8, a2)

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func q3() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	minimum := a[len(a)-1] - a[0]
	for i := 0; i <= k; i++ {
		if a[i+n-k-1]-a[i] < minimum {
			minimum = a[i+n-k-1] - a[i]
		}
	}
	fmt.Println(minimum)
}
