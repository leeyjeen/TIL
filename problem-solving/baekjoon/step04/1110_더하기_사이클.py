import sys

n = int(sys.stdin.readline())
cycle = 0

new_n = n
while True:
    new_n = (new_n%10)*10 + (int(new_n/10) + new_n%10)%10
    cycle = cycle + 1
    if new_n == n:
        break
print(cycle)