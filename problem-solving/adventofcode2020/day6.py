# https://adventofcode.com/2020/day/6
input_file = open('./day6.txt', 'r')
input_list = input_file.read().splitlines()

def solve_first():
    total_count = 0
    count_set = set()
    for i in input_list:
        if i == "":
            total_count += get_unique_question_count(count_set)
            count_set = set()
            continue
        update_count_set(i, count_set)
    total_count += get_unique_question_count(count_set)
    return total_count

def update_count_set(questions, count_set):
    count_set.update(set(questions))

def get_unique_question_count(count_set):
    return len(count_set)

def solve_second():
    total_count = 0
    count_dict = {}
    group_member_count = 0
    for i in input_list:
        if i == "":
            total_count += sum(value == group_member_count for value in count_dict.values())
            group_member_count = 0
            count_dict = {}
            continue
        group_member_count += 1
        update_count_dict(i, count_dict)
    total_count += sum(value == group_member_count for value in count_dict.values())
    return total_count

def update_count_dict(questions, count_dict):
    for q in list(questions):
        if q in count_dict:
            count_dict[q] += 1
        else:
            count_dict[q] = 1

if __name__ == "__main__":
    print(solve_first())
    print(solve_second())