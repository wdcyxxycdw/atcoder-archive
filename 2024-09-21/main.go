package main

import (
	"fmt"
	"strings"
	//"sort"
)

func main(){
	test()
}

func test(){
	fmt.Println("hello")
}

func q1(){
	s:=""
	fmt.Scanln(&s)
	a:=strings.Split(s, ".")
	ans:=""
	for i:=range a{
		ans+=a[i]
	}
	fmt.Println(ans)
}

func q2(){
	m:=100
	fmt.Scanln(&m)

	power:=[]int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049}

	for n:=1;n<=20;n++{
		ans:=make([]int,n)
		temp:=m
		for j:=0;j<n;j++{
			for i:=10;i>=0;i--{
				if power[i]<=temp{
					ans[j]=i
					temp-=power[i]
					if temp==0{
						fmt.Println(n)
						res:=""
						for k:=range ans{
							res+=fmt.Sprintf("%d ",ans[k])
						}
						fmt.Println(res)
						return
					}
					break
				}
			}
		}
	}
}

func q3() {
	n, q := 0, 0
	fmt.Scan(&n, &q)
	
	var s string
	fmt.Scan(&s)
	
	nums := make([]int, q)
	replace := make([]rune, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&nums[i])
		fmt.Scanf("%c",&replace[i])
	}
	sBytes := []byte(s)
	amount:=strings.Count(s, "ABC")


	isInABC:=func(i int)bool{
		switch sBytes[i]{
			case 'A':{
				if i+2<len(s)&&sBytes[i+1]=='B'&&sBytes[i+2]=='C'{
					return true
				}else{
					return false
				}
			}
			case 'B':{
				if i-1>=0&&i+1<len(s)&&sBytes[i-1]=='A'&&sBytes[i+1]=='C'{
					return true
				}else{
					return false
				}
			}
			case 'C':{
				if i-2>=0&&sBytes[i-1]=='B'&&sBytes[i-2]=='A'{
					return true
				}else{
					return false
				}
			}
			default:{
				return false
			}
		}
	}


	for i := 0; i < q; i++ {
		if isInABC(nums[i]-1){
			sBytes[nums[i]-1] = byte(replace[i])
			// s = string(sBytes)
			if !isInABC(nums[i]-1){
				amount--
			}
			fmt.Println(amount)
		}else{
			sBytes[nums[i]-1] = byte(replace[i])
			// s = string(sBytes)
			if isInABC(nums[i]-1){
				amount++
			}
			fmt.Println(amount)
		}
	}
}

func q4(){
	n:=0
	fmt.Scanln(&n)
	
	buildings:=make([]int,n)
	bmap:=make(map[int]int)
	for i:=0;i<n;i++{
		fmt.Scan(&buildings[i])
		bmap[buildings[i]]=i
	}
	ans:=make([]int,n)
	

	var cal func(l, r int)

	cal = func(l, r int){
		if l>=r{
			return
		}
		m:=buildings[r]
		mi:=r
		for i:=l+1;i<=r;i++{
			if buildings[i]>m{
				m=buildings[i]
				mi=i
			}
		}
		for i:=l;i<mi;i++{
			ans[i]++
		}
		cal(l,mi-1)
		cal(mi,r)
	}

	cal(0,n-1)
	res:=""
	for k:=range ans{
		res+=fmt.Sprintf("%d ",ans[k])
	}
	fmt.Println(res)
}