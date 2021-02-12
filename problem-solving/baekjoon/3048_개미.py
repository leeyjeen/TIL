# https://www.acmicpc.net/problem/3048
import sys

def ants(first, second, t):
    first = first[::-1]
    r = [ord(i) for i in first]
    r2 = [-ord(i) for i in second]
    r.extend(r2)
    for i in range(t):
        swapped = False
        for j in range(len(r)-1):
            if r[j] > 0 and r[j+1] < 0:
                if not swapped:
                    r[j], r[j+1] = r[j+1], r[j]
                    swapped = True
            else:
                swapped = False
    for i in range(len(r)):
        if r[i] < 0:
            r[i] = -r[i]
    ans = [chr(i) for i in r]
    return "".join(ans)

if __name__ == "__main__":
    n1, n2 = map(int, sys.stdin.readline().split())
    first = sys.stdin.readline().rstrip()
    second = sys.stdin.readline().rstrip()
    t = int(sys.stdin.readline())
    print(ants(first, second, t))