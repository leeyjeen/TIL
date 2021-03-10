# https://adventofcode.com/2020/day/1
input_file = open('./day1.txt', 'r')
input_list = list(map(int, input_file.read().splitlines()))

def solve_first():
    for i in input_list:
        if 2020-i in input_list:
            return i*(2020-i)

def solve_second():
    for i in range(len(input_list)):
        for j in range(i+1, len(input_list)):
            for k in range(j+1, len(input_list)):
                if input_list[i] + input_list[j] + input_list[k] == 2020:
                    return input_list[i]*input_list[j]*input_list[k]

if __name__ == "__main__":
    print(solve_first())
    print(solve_second())