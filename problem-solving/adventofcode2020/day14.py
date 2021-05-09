# https://adventofcode.com/2020/day/14
with open('./day14.txt', 'r') as f:
    lines = f.read().splitlines()

"""
PART 1:

mask 0 -> change to 0
mask 1 -> change to 1
mask X -> do not change
value update..
"""


def apply_masking(mask, val):
    val_after_masking = ""
    for m, v in zip(mask, val):
        if m == '0' or m == '1':
            val_after_masking += m
        elif m == 'X':
            val_after_masking += v
    return int(val_after_masking, 2)


def solve_part1(lines):
    memory = {}
    mask = ""
    for line in lines:
        splited = line.split(" = ")
        if splited[0] == "mask":
            mask = splited[1]
        else:
            pos = splited[0].split("[")[1].split("]")[0]
            val = bin(int(splited[1]))[2:].zfill(36)
            memory[pos] = apply_masking(mask, val)
    return sum(memory.values())


"""
PART 2:

[modifies memory address]
mask 0 -> do not change
mask 1 -> overwrite with 1
mask X -> take on all possible values...!
주소를 먼저 변환한 후, 변환한 주소들에 값을 매핑.

answer: sum of all values left in memory after it completes
"""


def apply_masking2(mask, addr, val):
    bin_addr = bin(addr)[2:].zfill(36)
    values_after_masking = dict({"": val})  # 빈 문자열인 key로 초기화
    for m, a in zip(mask, bin_addr):
        if m == "0":    # 각 key 값에 addr 값 더해서 업데이트
            new_values_after_masking = dict()
            for v in values_after_masking:
                new_values_after_masking[v+a] = val
            values_after_masking = new_values_after_masking
        elif m == "1":  # 각 key 값에 "1" 값 더해서 업데이트
            new_values_after_masking = dict()
            for v in values_after_masking:
                new_values_after_masking[v+"1"] = val
            values_after_masking = new_values_after_masking
        elif m == "X":  # 각 key 값에 모든 경우 ("0", "1") 값 더해서 업데이트
            new_values_after_masking = dict()
            for v in values_after_masking:
                new_values_after_masking[v+"0"] = val
                new_values_after_masking[v+"1"] = val
            values_after_masking = new_values_after_masking
    return values_after_masking


def solve_part2(lines):
    addresses = dict()
    mask = ""
    for line in lines:
        splited = line.split(" = ")
        if splited[0] == "mask":
            mask = splited[1]
        else:
            addr = int(splited[0].split("[")[1].split("]")[0])
            val = int(splited[1])
            addresses.update(apply_masking2(mask, addr, val))
    return sum(addresses.values())


if __name__ == "__main__":
    print(solve_part1(lines))
    print(solve_part2(lines))
