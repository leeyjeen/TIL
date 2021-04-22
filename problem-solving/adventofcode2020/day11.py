# https://adventofcode.com/2020/day/11
with open('./day11.txt', 'r') as f:
    seats = [list(line) for line in f.read().splitlines()]


def get_adjacent_seats_count(seats, row, col):
    count = 0
    for i in [-1, 0, 1]:
        if row+i < 0 or row+i >= len(seats):
            continue
        for j in [-1, 0, 1]:
            if i == 0 and j == 0:
                continue
            if col+j < 0 or col+j >= len(seats[i]):
                continue
            seat = seats[row+i][col+j]
            if seat == '#':
                count += 1
    return count


def get_seats_after_one_round(seats):
    new_seats = []
    for r in range(0, len(seats)):
        new_row = []
        for c in range(0, len(seats[r])):
            seat = seats[r][c]
            if seat == "L":
                if get_adjacent_seats_count(seats, r, c) == 0:
                    new_row.append("#")
                else:
                    new_row.append("L")
            elif seat == "#":
                if get_adjacent_seats_count(seats, r, c) >= 4:
                    new_row.append("L")
                else:
                    new_row.append("#")
            else:
                new_row.append(".")
        new_seats.append(new_row)
    return new_seats


def solvePart1(seats):
    prev_seats = seats
    new_seats = []
    occupied_count = 0
    while True:
        new_seats = get_seats_after_one_round(prev_seats)
        if prev_seats == new_seats:
            occupied_count = sum(s.count('#') for s in new_seats)
            break
        prev_seats = new_seats
    return occupied_count


if __name__ == "__main__":
    print(solvePart1(seats))
