# https://www.acmicpc.net/problem/11055
import sys

def longest_ascending_sequence_sum(sequence):
    dp = [0 for _ in range(len(sequence))]
    max_sum = 0
    for i in range(len(sequence)):
        max_dp = 0
        for j in range(i):
            if sequence[j] < sequence[i]:
                max_dp = max(max_dp, dp[j])
        dp[i] = max_dp+sequence[i]
        max_sum = max(max_sum, dp[i])
    return max_sum

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    a = list(map(int, sys.stdin.readline().split()))
    print(longest_ascending_sequence_sum(a))