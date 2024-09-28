package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	q4()
}

func q1(){
	var r int
	fmt.Scan(&r)

	if r<100{
		fmt.Println(100-r)
	}else if r<200{
		fmt.Println(200-r)
	}else if r<300{
		fmt.Println(300-r)
	}
}

func q2(){
	var n, t, p int
	fmt.Scan(&n, &t, &p)

	l:=make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&l[i])
	}

	sort.Slice(l, func(i, j int)bool{
		if l[i]>l[j]{
			return true
		}else{
			return false
		}
	})
	if l[p-1]>t{
		fmt.Println(0)
	}else{
		fmt.Println(t-l[p-1])
	}
}

func q3(){
	var n, k int
	fmt.Scan(&n, &k)

	s:=""
	fmt.Scan(&s)

	isPalindrome:=func(str string)bool{
		l, r:=0, len(str)-1

		for l<=r{
			if str[l]!=str[r]{
				return false
			}else{
				l++
				r--
			}
		}
		return true
	}

	var permutation func(S string)[]string
	permutation=func(S string) []string {
		if len(S) == 1 {
			return []string{S}
		}
		// 与拼接得到的各个字符串再进行拼接
		ret := []string{}
		for i, s := range S {
			// 差了第i个字符的剩余字符串往下传，并将得到的结果进行合并
			tmp := fmt.Sprintf("%s%s", S[:i], S[i+1:])
			res := permutation(tmp)
			for _, r := range res {
				ret = append(ret, fmt.Sprintf("%c%s", s, r))
			}
		}
		return ret
	}
	
	c:=permutation(s)
	smap:=make(map[string]bool)
	for i:=range c{
		smap[c[i]]=true
	}

	cnt:=0
	for S:=range smap{
		l, r:=0, k
		cnt++
		for r<=len(s){
			if isPalindrome(S[l:r]){
				cnt--
				break
			}
			l++
			r++
		}
	}

	fmt.Println(cnt)
}

func q4(){
	ans:=make([]uint64, 2)
	ans[0]=10
	ans[1]=20

	search:=func(n uint64)int{
		for ans[len(ans)-1]<n{
			ans = append(ans, ans[len(ans)-1]*11)
		}
		return len(ans)
	}


	var re func (n uint64, k int)string
	re=func (n uint64, k int)string{
		if k==2{
			s:=strconv.FormatUint(n, 10)
			return s+s
		}
		if k==1{
			s:=strconv.FormatUint(n, 10)
			return s
		}

		a:=search(n)
		
	
	}


}