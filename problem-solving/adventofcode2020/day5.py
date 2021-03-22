# https://adventofcode.com/2020/day/5

input_file = open('./day5.txt', 'r')
input_list = input_file.read().splitlines()

def get_highest_seat_id(input_list):
    highest_seat_id = 0
    for input in input_list:
        row_string, col_string = input[0:7], input[7:10]
        row = compute_row(row_string)
        col = compute_col(col_string)
        unique_seat_id = get_unique_seat_id(row, col)
        highest_seat_id = max(highest_seat_id, unique_seat_id)
    return highest_seat_id

def find_my_missing_seat(input_list):
    seat_ids = []
    for input in input_list:
        row_string, col_string = input[0:7], input[7:10]
        row = compute_row(row_string)
        col = compute_col(col_string)
        unique_seat_id = get_unique_seat_id(row, col)
        seat_ids.append(unique_seat_id)
    seat_ids.sort()
    my_id = 0
    for i in range(1, len(seat_ids)):
        if seat_ids[i] - seat_ids[i-1] != 1:
            my_id = seat_ids[i]-1
            break
    return my_id

def compute_row(row_string):
    start = 0
    end = 127
    for v in list(row_string):
        diff = (end-start)//2+1
        if v == "F":
            end -= diff
        elif v == "B":
            start += diff
    return start

def compute_col(col_string):
    start = 0
    end = 7
    for v in list(col_string):
        diff = (end-start)//2+1
        if v == "L":
            end -= diff
        elif v == "R":
            start += diff
    return start

def get_unique_seat_id(row, col):
    return row * 8 + col

if __name__ == "__main__":
    print(get_highest_seat_id(input_list))
    print(find_my_missing_seat(input_list))
