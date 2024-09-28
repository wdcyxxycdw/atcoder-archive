s=input().split(" ")

n=int(s[0])
l=int(s[1])
r=int(s[2])

a=[x for x in range(1, n+1)]

left=a[:l-1]
mid=a[l-1:r]
mid.reverse()
right=a[r:]

ans=[]
ans.extend(left)
ans.extend(mid)
ans.extend(right)

for i in ans:
    print(i, end=" ")