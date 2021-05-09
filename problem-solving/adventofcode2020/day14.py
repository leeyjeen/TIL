# https://adventofcode.com/2020/day/14
with open('./day14.txt', 'r') as f:
    lines = f.read().splitlines()


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


if __name__ == "__main__":
    print(solve_part1(lines))
