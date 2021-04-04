# https://adventofcode.com/2020/day/9
with open('./day9.txt', 'r') as f:
    input_list = [int(line) for line in f.read().splitlines()]


def find_first_num(preamble_len):
    for i in range(preamble_len, len(input_list)):
        num_to_check = input_list[i]
        is_valid = False
        for j in range(i-preamble_len, i-1):
            num_to_check2 = input_list[j]
            if abs(num_to_check - num_to_check2) in input_list[i-preamble_len:j] or abs(num_to_check - num_to_check2) in input_list[j:i]:
                is_valid = True
                break
        if not is_valid:
            return num_to_check
    return 0


def solve_part1(preamble_len):
    return find_first_num(preamble_len)


def compute_sum_dp():
    dp = [[0]*len(input_list)
          for _ in range(len(input_list))]  # i~j까지의 합을 저장할 배열
    for i in range(len(input_list)):
        dp[i][i] = input_list[i]
    for i in range(len(input_list)):
        for j in range(i+1, len(input_list)):
            dp[i][j] = dp[i][j-1] + input_list[j]
    return dp


def find_encryption_weakness(first_number):
    dp = compute_sum_dp()
    for i in range(len(input_list)):
        for j in range(len(input_list)):
            if dp[i][j] == first_number:
                arr = input_list[i:j+1]
                return min(arr) + max(arr)


def solve_part2(first_number):
    return find_encryption_weakness(first_number)


if __name__ == "__main__":
    first_number = solve_part1(25)
    print(first_number)
    print(solve_part2(first_number))
