# https://adventofcode.com/2020/day/12
with open('./day12.txt', 'r') as f:
    input_list = [(line[0], int(line[1:])) for line in f.read().splitlines()]


def solve_part1(input_list):
    facing_direction = "E"
    current_degree = 0
    horizontal_value = 0
    vertical_value = 0
    degree_direction_dict = {0: "E", 90: "S", 180: "W", 270: "N"}
    for i in input_list:
        action = i[0]
        if action == "F":
            action = facing_direction
        if action == "N":
            vertical_value += i[1]
        elif action == "S":
            vertical_value -= i[1]
        elif action == "E":
            horizontal_value += i[1]
        elif action == "W":
            horizontal_value -= i[1]
        elif action == "L":
            current_degree = (current_degree - i[1]) % 360
            facing_direction = degree_direction_dict[current_degree]
        elif action == "R":
            current_degree = (current_degree + i[1]) % 360
            facing_direction = degree_direction_dict[current_degree]
    return abs(horizontal_value) + abs(vertical_value)


if __name__ == "__main__":
    print(solve_part1(input_list))
