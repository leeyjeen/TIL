# https://www.acmicpc.net/problem/11722
import sys

def longest_descreasing_sequence_count(sequence):
    dp = [0 for _ in range(len(sequence))]
    longest_count = 0
    for i in range(len(sequence)-1, -1, -1):
        dp[i] = 1
        length = 0
        for j in range(i+1, len(sequence)):
            if sequence[i] > sequence[j]:
                length = max(length, dp[j])
        dp[i] += length
        longest_count = max(longest_count, dp[i])
    return longest_count

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    sequence = list(map(int, sys.stdin.readline().split()))
    print(longest_descreasing_sequence_count(sequence))