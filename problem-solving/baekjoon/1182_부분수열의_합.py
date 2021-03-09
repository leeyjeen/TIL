# https://www.acmicpc.net/problem/1182
import sys

count = 0
seq = []

def compute_subsequence_count(index, sum, aim):
    global count
    if index > len(seq)-1:
        return
    sum += seq[index]
    if sum == aim:
        count += 1
    compute_subsequence_count(index+1, sum, aim)
    compute_subsequence_count(index+1, sum-seq[index], aim)

if __name__ == "__main__":
    n, s = list(map(int, sys.stdin.readline().split()))
    seq = list(map(int, sys.stdin.readline().split()))
    compute_subsequence_count(0, 0, s)
    print(count)
