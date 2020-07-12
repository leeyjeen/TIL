import sys

a = int(sys.stdin.readline())
b = int(sys.stdin.readline())

print(a*(b%10))
print(a*(int(b/10)%10))
print(a*(int(b/100)%10))
print(a*b)
