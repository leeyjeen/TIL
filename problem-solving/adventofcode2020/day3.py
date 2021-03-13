# https://adventofcode.com/2020/day/3
input_file = open('./day3.txt', 'r')
input_list = input_file.read().splitlines()

def solve_first():
    return get_tree_count(3, 1)

def solve_second():
    return get_tree_count(1, 1)*get_tree_count(3, 1)*get_tree_count(5,1)*get_tree_count(7, 1)*get_tree_count(1, 2)

def get_tree_count(right_interval, down_interval):
    row, col = 0, 0
    tree_count = 0
    while row+down_interval < len(input_list):
        col += right_interval
        col %= len(input_list[0])
        row += down_interval
        if input_list[row][col] == "#":
            tree_count += 1
    return tree_count

if __name__ == "__main__":
    print(solve_first())
    print(solve_second())
            
