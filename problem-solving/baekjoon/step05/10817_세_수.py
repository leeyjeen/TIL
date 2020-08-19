import sys

a, b, c = map(int, sys.stdin.readline().split())

if a >= b:
    if a >= c:
        if b >= c:
            second = b
        else:
            second = c
    else:
        second = a
elif b >= c:
    if a >= c:
        second = a
    else:
        second = c
else:
    second = b
    
print(second)