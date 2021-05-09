# https://adventofcode.com/2020/day/15

def solve_part1(numbers):
    while len(numbers) < 2020:
        if numbers.count(numbers[-1]) == 1:
            numbers.append(0)
        else:
            spoken_turns = [i+1 for i,
                            num in enumerate(numbers) if num == numbers[-1]]
            numbers.append(spoken_turns[-1] - spoken_turns[-2])
    return numbers[-1]


if __name__ == "__main__":
    print(solve_part1([0, 14, 1, 3, 7, 9]))
