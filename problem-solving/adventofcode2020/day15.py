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


"""
[0,3,6]
0 3 6 
**Turn 4**
    last: 6
    number: 0 (because 6 is first..)
**Turn 5**
    last: 0
    number: 4-1 = 3 (0's turn...)
...
"""


def solve_part2(numbers):
    spoken = {}  # 각 number가 등장한 마지막 turn값을 저장

    for i, num in enumerate(numbers):
        spoken[num] = i+1

    last_num = numbers[-1]
    appeared_first = True

    for cur_turn in range(len(numbers)+1, 30000001):
        if appeared_first:  # 처음 나타난 경우 speak할 number는 0이다
            last_num = 0
        else:               # 이전에 나타난 적이 있을 경우 speak할 number는 turn의 차이값이다
            last_num = turn_diff

        appeared_first = last_num not in spoken  # speak하는 number가 이전에 나타난 적이 없는지 확인
        if not appeared_first:
            # 다음 turn에서 저장할 값 계산 (현재 speak하는 number의 diff)
            turn_diff = cur_turn - spoken[last_num]
        spoken[last_num] = cur_turn  # speak하는 number의 turn 업데이트

    return last_num


if __name__ == "__main__":
    print(solve_part1([0, 14, 1, 3, 7, 9]))
    print(solve_part2([0, 14, 1, 3, 7, 9]))
