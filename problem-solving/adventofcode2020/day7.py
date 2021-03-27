# https://adventofcode.com/2020/day/7
with open('./day7.txt', 'r') as f:
    input_list = [line for line in f.read().splitlines()]

"""
ex.
['shiny gold bags contain 1 faded blue bag, 2 muted yellow bags.', 'faded blue bags contain no other bags.', 'muted yellow bags contain no other bags.']
-> 
{'shiny gold': [{'faded blue': 1}, {'muted yellow': 2}], 'faded blue': [], 'muted yellow': []}
"""
def parse_input_list(input_list):
    bag_dict = {}
    for line in input_list:
        cur_bag = " ".join(line.split(" ")[:2])
        _, inner_bags = line.split("bags contain ")
        inner_bags = inner_bags.split(", ")
        
        bag_dict[cur_bag] = []
        for inner_bag in inner_bags:
            if inner_bag.split(" ")[0] == "no":
                continue
            count, bag_type = int(inner_bag.split(" ")[0]), " ".join(inner_bag.split(" ")[1:3])
            bag_dict[cur_bag].append({bag_type: count})
        contains = []
        for _, val in bag_dict.items():
            if cur_bag in val:
                contains.append(val)
    return bag_dict

def solve_first():
    return find_bag_count_containing_target_bag("shiny gold")

# target_bad을 포함하는 bag의 집합 개수를 구한다.
def find_bag_count_containing_target_bag(target_bag):
    contains_set = set()
    find_contains_set(target_bag, contains_set)
    return len(contains_set)

# target_bag을 포함하는 bag을 찾아서 contain_set(집합)에 추가. 해당 bag을 포함하는 bag도 재귀적으로 찾아서 추가함.
def find_contains_set(target_bag, contain_set):
    for key, val in bag_dict.items():
        for v in val:
            if target_bag in v:
                contain_set.add(key)
                find_contains_set(key, contain_set)

def solve_second():
    return find_inner_bag_count("shiny gold")

# bag의 내부에 있는 가방 수를 구한다.
def find_inner_bag_count(bag):
    if len(bag_dict[bag]) == 0:
        return 0
    count = 0
    for inner_bags in bag_dict[bag]:
        count += sum([c for c in inner_bags.values()]) # 한 단계 아래의 inner_bags의 개수를 더한다. 
        for key, val in inner_bags.items(): # 각 inner_bag마다 재귀적으로 
            inside_bag_count = find_inner_bag_count(key) # 해당 inner_bag의 inner_bag개수를 구한 후
            count += val*inside_bag_count # 해당 inner_bag의 개수를 곱해서 더한다.
    return count

if __name__ == "__main__":
    bag_dict = parse_input_list(input_list)
    print(solve_first())
    print(solve_second())