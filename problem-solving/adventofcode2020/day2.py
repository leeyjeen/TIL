# https://adventofcode.com/2020/day/2
input_file = open('./day2.txt', 'r')
input_list = input_file.read().splitlines()

def solve_first():
    valid_count = 0
    for input in input_list:
        times, letter, password = input.split()
        lowest_time, highest_time = list(map(int, times.split("-")))
        letter = letter[0]
        if check_if_password_has_valid_letter_count(password, letter, lowest_time, highest_time):
            valid_count += 1
    return valid_count

def check_if_password_has_valid_letter_count(password, letter, lowest_time, highest_time):
    count = password.count(letter)
    if count >= lowest_time and count <= highest_time:
        return True
    return False

def solve_second():
    valid_count = 0
    for input in input_list:
        position, letter, password = input.split()
        first_position, second_position = list(map(int, position.split("-")))
        letter = letter[0]
        if check_if_password_has_one_right_letter_position(password, letter, first_position, second_position):
            valid_count += 1
    return valid_count

def check_if_password_has_one_right_letter_position(password, letter, first_position, second_position):
    first_contains = password[first_position-1] == letter
    second_contains = password[second_position-1] == letter
    if first_contains and not second_contains or not first_contains and second_contains:
        return True
    return False

if __name__ == "__main__":
    print(solve_first())
    print(solve_second())