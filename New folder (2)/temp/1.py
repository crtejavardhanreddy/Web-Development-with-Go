n=int(input())
arr=list(map(int,input().split(",")))
L,R=input().split()
li=int(L)-1
ri=int(R)-1
c=0
x1=arr[li]
for i in range(li,ri):
    if arr[i]>=x1:c+=1
print(c)