# https://adventofcode.com/2020/day/10
from collections import Counter

with open('./day10.txt', 'r') as f:
    input_list = [int(line) for line in f.read().splitlines()]
    input_list.sort()


def solve_part1():
    prev = 0
    one_jolt_count, three_jolt_count = 0, 0
    for i in input_list:
        one_jolt_count += int(i-prev == 1)
        three_jolt_count += int(i-prev == 3)
        prev = i
    three_jolt_count += 1
    return one_jolt_count*three_jolt_count


def solve_part2():
    dp = Counter()
    dp[0] = 1
    for num in input_list:
        dp[num] = dp[num - 1] + dp[num - 2] + dp[num - 3]
    return dp[input_list[-1]]


if __name__ == "__main__":
    print(solve_part1())
    print(solve_part2())
