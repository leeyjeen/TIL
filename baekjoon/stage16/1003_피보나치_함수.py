import sys

def result(n):
    zero_count = [1, 0]
    one_count = [0, 1]
    if n <= 1:
        return zero_count, one_count
    for i in range(2, n+1):
        zero_count.append(zero_count[i-1]+zero_count[i-2])
        one_count.append(one_count[i-1]+one_count[i-2])
    return zero_count, one_count

t = int(sys.stdin.readline())

for i in range(0, t):
    n = int(sys.stdin.readline())
    zero_count, one_count = result(n)
    print('%d %d' %(zero_count[n], one_count[n]))
