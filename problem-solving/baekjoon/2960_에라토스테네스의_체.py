# https://www.acmicpc.net/problem/2960
import sys

def eratos(n, k):
    prime_numbers = [True for _ in range(n+1)]
    kth = 0
    for i in range(2, n+1):
        if prime_numbers[i]:
            for j in range(i, n+1, i):
                if prime_numbers[j]:
                    prime_numbers[j] = False
                    k -= 1
                    if k == 0:
                        kth = j
                        return kth
    return kth

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    print(eratos(n, k))