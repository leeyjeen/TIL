# https://www.acmicpc.net/problem/2422
import sys

def get_count_of_good_combination(bad_combination):
    count = 0
    for i in range(len(bad_combination)-2):
        for j in range(i+1, len(bad_combination)-1):
            if bad_combination[i][j]:
                continue
            for k in range(j+1, len(bad_combination)):
                if bad_combination[i][k] or bad_combination[j][k]:
                    continue
                count += 1
    return count

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    bad_combination = [[False for j in range(n)] for i in range(n)]

    for i in range(m):
        ice_one, ice_two = list(map(int, sys.stdin.readline().split()))
        bad_combination[ice_one-1][ice_two-1] = True
        bad_combination[ice_two-1][ice_one-1] = True

    print(get_count_of_good_combination(bad_combination))