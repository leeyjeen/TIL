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


def solve_part2(input_list):
    waypoint = [10, 1]
    ship_pos = [0, 0]
    for i in input_list:
        action = i[0]
        if action == "F":
            ship_pos = [ship_pos[0]+waypoint[0] *
                        i[1], ship_pos[1]+waypoint[1]*i[1]]
        elif action == "L":
            if i[1] == 90:
                waypoint = [-waypoint[1], waypoint[0]]
            elif i[1] == 180:
                waypoint = [-waypoint[0], -waypoint[1]]
            elif i[1] == 270:
                waypoint = [waypoint[1], -waypoint[0]]
        elif action == "R":
            if i[1] == 90:
                waypoint = [waypoint[1], -waypoint[0]]
            elif i[1] == 180:
                waypoint = [-waypoint[0], -waypoint[1]]
            elif i[1] == 270:
                waypoint = [-waypoint[1], waypoint[0]]
        elif action == "N":
            waypoint[1] += i[1]
        elif action == "S":
            waypoint[1] -= i[1]
        elif action == "E":
            waypoint[0] += i[1]
        elif action == "W":
            waypoint[0] -= i[1]
    return abs(ship_pos[0]) + abs(ship_pos[1])


if __name__ == "__main__":
    print(solve_part1(input_list))
    print(solve_part2(input_list))
