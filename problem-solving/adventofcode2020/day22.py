# https://adventofcode.com/2020/day/22
with open('./day22.txt', 'r') as f:
    p1, p2 = f.read().split("\n\n")
    cards1 = list(map(int, p1.splitlines()[1:]))
    cards2 = list(map(int, p2.splitlines()[1:]))


def solve_part1():
    ccards1, ccards2 = cards1.copy(), cards2.copy()
    while len(ccards1) > 0 and len(ccards2) > 0:
        num1 = ccards1.pop(0)
        num2 = ccards2.pop(0)
        if num1 > num2:
            ccards1.append(num1)
            ccards1.append(num2)
        elif num1 < num2:
            ccards2.append(num2)
            ccards2.append(num1)
    result = 0
    if len(ccards1) > 0:
        for i in range(len(ccards1)):
            result += ccards1[i]*(len(ccards1)-i)
    else:
        for i in range(len(ccards2)):
            result += ccards2[i]*(len(ccards2)-i)
    return result


def do_subgame(ccards1, ccards2):
    history = set()
    while len(ccards1) > 0 and len(ccards2) > 0:
        deck = (tuple(ccards1), tuple(ccards2))
        if deck in history:
            return 1
        history.add(deck)

        num1 = ccards1.pop(0)
        num2 = ccards2.pop(0)

        if num1 <= len(ccards1) and num2 <= len(ccards2):
            winner = do_subgame(ccards1[:num1], ccards2[:num2])
        else:
            winner = 1 if num1 > num2 else 2
        if winner == 1:
            ccards1.append(num1)
            ccards1.append(num2)
        elif winner == 2:
            ccards2.append(num2)
            ccards2.append(num1)
    return 1 if len(ccards1) > 0 else 2


def solve_part2():
    ccards1, ccards2 = cards1.copy(), cards2.copy()
    do_subgame(ccards1, ccards2)
    result = 0
    if len(ccards1) > 0:
        for i in range(len(ccards1)):
            result += ccards1[i]*(len(ccards1)-i)
    else:
        for i in range(len(ccards2)):
            result += ccards2[i]*(len(ccards2)-i)
    return result


if __name__ == "__main__":
    print(solve_part1())
    print(solve_part2())
