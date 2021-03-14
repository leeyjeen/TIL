# https://adventofcode.com/2020/day/4
import re

input_file = open('./day4.txt', 'r')
input_list = input_file.read().splitlines()

def solve_first():
    return get_valid_count(input_list, has_all_required_fields)

def solve_second():
    return get_valid_count(input_list, has_all_required_fields_and_valid)

def get_valid_count(input_list, check_valid_func):
    passport_data = {}
    valid_count = 0
    for input in input_list:
        if input == "":
            if check_valid_func(passport_data):
                valid_count += 1
            passport_data = {}
            continue
        key_values = input.split(' ')
        for key_value in key_values:
            key, value = key_value.split(':')
            passport_data[key] = value
    if check_valid_func(passport_data):
        valid_count += 1
    return valid_count

def has_all_required_fields(passport_data):
    return all(field in passport_data for field in ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"])

def has_all_required_fields_and_valid(p):
    return check_byr(p) and check_iyr(p) and check_eyr(p) and check_hgt(p) and check_hcl(p) and check_ecl(p) and check_pid(p)

def check_byr(passport_data):
    if "byr" not in passport_data:
        return False
    byr = passport_data["byr"]
    if not byr.isdigit() or len(byr) != 4:
        return False
    return 1920 <= int(byr) <= 2002

def check_iyr(passport_data):
    if "iyr" not in passport_data:
        return False
    iyr = passport_data["iyr"]
    if not iyr.isdigit() or len(iyr) != 4:
        return False
    return 2010 <= int(iyr) <= 2020

def check_eyr(passport_data):
    if "eyr" not in passport_data:
        return False
    eyr = passport_data["eyr"]
    if not eyr.isdigit() or len(eyr) != 4:
        return False
    return 2020 <= int(eyr) <= 2030

def check_hgt(passport_data):
    if "hgt" not in passport_data:
        return False
    hgt = passport_data["hgt"]
    followed = hgt[-2:]
    if followed == "cm":
        if not 150 <= int(hgt[:-2]) <= 193:
            return False
    elif followed == "in":
        if not 59 <= int(hgt[:-2]) <= 76:
            return False
    else:
        return False
    return True

def check_hcl(passport_data):
    if "hcl" not in passport_data:
        return False
    hcl = passport_data["hcl"]
    if len(hcl) != 7:
        return False
    if hcl[0] != "#":
        return False
    p = re.compile("[0-9a-f]+")
    return p.match(hcl[1:])

def check_ecl(passport_data):
    if "ecl" not in passport_data:
        return False
    return passport_data["ecl"] in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]

def check_pid(passport_data):
    if "pid" not in passport_data:
        return False
    pid = passport_data["pid"]
    return pid.isdigit() and len(pid) == 9

if __name__ == "__main__":
    print(solve_first())
    print(solve_second())