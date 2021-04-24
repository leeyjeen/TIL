# https://adventofcode.com/2020/day/11
with open('./day11.txt', 'r') as f:
    seats = [list(line) for line in f.read().splitlines()]


def get_adjacent_seats_count(seats, row, col):
    count = 0
    eight_directions = [(-1, -1), (-1, 0), (0, -1), (1, 1),
                        (1, 0), (0, 1), (-1, 1), (1, -1)]
    for d in eight_directions:
        diff_row, diff_col = d[0], d[1]
        r = row+diff_row
        c = col+diff_col
        if r >= 0 and r < len(seats) and c >= 0 and c < len(seats[0]):
            seat = seats[r][c]
            if seat == '#':
                count += 1
    return count


def get_eight_direction_seats_count(seats, row, col):
    count = 0
    eight_directions = [(-1, -1), (-1, 0), (0, -1), (1, 1),
                        (1, 0), (0, 1), (-1, 1), (1, -1)]
    for d in eight_directions:
        diff_row, diff_col = d[0], d[1]
        r = row+diff_row
        c = col+diff_col
        while r >= 0 and r < len(seats) and c >= 0 and c < len(seats[0]):
            seat = seats[r][c]
            if seat == '#':
                count += 1
                break
            elif seat == 'L':
                break
            else:
                r += diff_row
                c += diff_col

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


def get_seats_after_one_round2(seats):
    new_seats = []
    for r in range(0, len(seats)):
        new_row = []
        for c in range(0, len(seats[r])):
            seat = seats[r][c]
            if seat == "L":
                if get_eight_direction_seats_count(seats, r, c) == 0:
                    new_row.append("#")
                else:
                    new_row.append("L")
            elif seat == "#":
                if get_eight_direction_seats_count(seats, r, c) >= 5:
                    new_row.append("L")
                else:
                    new_row.append("#")
            else:
                new_row.append(".")
        new_seats.append(new_row)
    return new_seats


def solve_part1(seats):
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


def solve_part2(seats):
    prev_seats = seats
    new_seats = []
    occupied_count = 0
    while True:
        new_seats = get_seats_after_one_round2(prev_seats)
        if prev_seats == new_seats:
            occupied_count = sum(s.count('#') for s in new_seats)
            break
        prev_seats = new_seats
    return occupied_count


if __name__ == "__main__":
    print(solve_part1(seats))
    print(solve_part2(seats))
