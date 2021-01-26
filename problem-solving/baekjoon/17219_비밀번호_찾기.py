# https://www.acmicpc.net/problem/17219
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    passwords = {}
    
    for i in range(n):
        site, password = sys.stdin.readline().rstrip().split(" ")
        passwords[site] = password

    for i in range(m):
        site = sys.stdin.readline().rstrip()
        print(passwords[site])