n=input()
a=[]
sum=0

for i in range(int(n)):
   s=input()
   s=s.split()
   a.append(s)
   sum+=int(s[1])

winner=sum%int(n)

a=sorted(a)
print(a[winner][0])

