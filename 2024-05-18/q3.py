n=int(input())
card=[]

for i in range(n):
    a=input()
    a=a.split()
    a.append(str(i))
    card.append(a)

def convert_to_int(lst):
    if isinstance(lst, list):
        return [convert_to_int(item) for item in lst]
    else:
        return int(lst)

card=convert_to_int(card)
card.sort()

m={}

for i in range(len(card)):
    j=i+1
    while(j<len(card)):
        if card[i][1]>card[j][1]:
            m[card[i][2]]=True
            del card[i]
            i-=1
            break
        j+=1
    
ans=[]
for i in range(n):
    if i not in m:
        ans.append(i+1)

print(len(ans))
output=' '.join(map(str, ans))
print(output)