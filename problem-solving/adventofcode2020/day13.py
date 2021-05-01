# https://adventofcode.com/2020/day/13
import sys

with open('./day13.txt', 'r') as f:
    depart_time_stamp, buses = f.read().splitlines()
    depart_time_stamp = int(depart_time_stamp)
    bus_set = set(buses.split(","))
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


if __name__ == "__main__":
    print(solve_part1())
