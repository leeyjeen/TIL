import sys

a, b, c = sys.stdin.readline().split()
a, b, c = int(a), int(b), int(c)

print((a+b)%c)
print(((a%c)+(b%c))%c)
print((a*b)%c)
print((a%c)*(b%c)%c)