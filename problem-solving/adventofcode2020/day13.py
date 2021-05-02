# https://adventofcode.com/2020/day/13
import sys
from collections import defaultdict

with open('./day13.txt', 'r') as f:
    depart_time_stamp, input_buses = f.read().splitlines()
    depart_time_stamp = int(depart_time_stamp)
    bus_set = set(input_buses.split(","))
    bus_set.discard("x")
    bus_list = list(map(int, bus_set))
    bus_list.sort()


def solve_part1():
    min_time_stamp = sys.maxsize
    bus_number = 0
    for bus in bus_list:
        remain = depart_time_stamp % bus
        if remain > 0:
            if min_time_stamp > depart_time_stamp + (bus-remain):
                min_time_stamp = depart_time_stamp + (bus-remain)
                bus_number = bus
    return (min_time_stamp-depart_time_stamp) * bus_number


def solve_part2():
    bus_numbers = [(i, int(bus))
                   for i, bus in enumerate(input_buses.split(",")) if bus != "x"]
    earliest_time_stamp = 0
    while True:
        is_valid = True
        for i in range(1, len(bus_numbers)):
            bus_tuple = bus_numbers[i]
            if (earliest_time_stamp + bus_tuple[0]) % bus_tuple[1] != 0:
                earliest_time_stamp += bus_numbers[0][1]
                is_valid = False
                break
        if is_valid:
            break

    return earliest_time_stamp


if __name__ == "__main__":
    print(solve_part1())
    print(solve_part2())
