# https://adventofcode.com/2020/day/18
with open('./day18.txt', 'r') as f:
    lines = [line.replace(" ", "") for line in f.read().splitlines()]


def evaluate(line):
    numbers = []
    operators = []
    in_the_brackets = []
    while len(line) > 0:
        popped = line[0]
        line = line[1:]
        if popped.isdigit():
            numbers.append(popped)
        else:
            if popped == ')':
                while True:
                    in_the_brackets.insert(0, numbers.pop())
                    if operators[-1] == '(':
                        operators.pop()
                        break
                    else:
                        in_the_brackets.insert(0, operators.pop())
                while len(in_the_brackets) >= 3:
                    num1 = in_the_brackets.pop(0)
                    opr = in_the_brackets.pop(0)
                    num2 = in_the_brackets.pop(0)
                    in_the_brackets.insert(0, eval(str(num1)+opr+str(num2)))
                numbers.append(in_the_brackets[0])
                in_the_brackets.clear()
            else:
                operators.append(popped)
    while len(numbers) > 1:
        num1 = numbers.pop(0)
        num2 = numbers.pop(0)
        opr = operators.pop(0)
        numbers.insert(0, eval(str(num1)+opr+str(num2)))
    return numbers[0]


def evaluate2(line):
    numbers = []
    operators = []
    in_the_brackets = []
    while len(line) > 0:
        popped = line[0]
        line = line[1:]
        if popped.isdigit():
            numbers.append(popped)
        else:
            if popped == ')':
                while True:
                    in_the_brackets.insert(0, numbers.pop())
                    if operators[-1] == '(':
                        operators.pop()
                        break
                    else:
                        in_the_brackets.insert(0, operators.pop())
                while len(in_the_brackets) >= 3:
                    if "+" in in_the_brackets:
                        idx = in_the_brackets.index("+")
                        num1 = in_the_brackets.pop(idx-1)
                        opr = in_the_brackets.pop(idx-1)
                        num2 = in_the_brackets.pop(idx-1)
                        in_the_brackets.insert(
                            idx-1, eval(str(num1)+opr+str(num2)))
                    else:
                        num1 = in_the_brackets.pop(0)
                        opr = in_the_brackets.pop(0)
                        num2 = in_the_brackets.pop(0)
                        in_the_brackets.insert(
                            0, eval(str(num1)+opr+str(num2)))
                numbers.append(in_the_brackets[0])
                in_the_brackets.clear()
            else:
                operators.append(popped)
    while len(numbers) > 1:
        if "+" in operators:
            idx = operators.index("+")
            num1 = numbers.pop(idx)
            num2 = numbers.pop(idx)
            opr = operators.pop(idx)
            numbers.insert(
                idx, eval(str(num1)+opr+str(num2)))
        else:
            num1 = numbers.pop(0)
            opr = operators.pop(0)
            num2 = numbers.pop(0)
            numbers.insert(
                0, eval(str(num1)+opr+str(num2)))
    return numbers[0]


def solve_part1(lines):
    sum = 0
    for line in lines:
        val = evaluate(line)
        sum += val
    return sum


def solve_part2(lines):
    sum = 0
    for line in lines:
        val = evaluate2(line)
        sum += val
    return sum


if __name__ == "__main__":
    print(solve_part1(lines))
    print(solve_part2(lines))
