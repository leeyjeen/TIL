# https://adventofcode.com/2020/day/8
with open('./day8.txt', 'r') as f:
    input_list = [line for line in f.read().splitlines()]


class Command:
    def __init__(self, opr, arg):
        self.operation = opr
        self.argument = arg

    def convert_operation(self):
        if self.operation == "nop":
            self.operation = "jmp"
            return True
        elif self.operation == "jmp":
            self.operation = "nop"
            return True
        return False


def solve_first():
    visited = {}
    idx = 0
    accumulator = 0
    commands = parse_inputs_to_commands(input_list)
    while idx not in visited:
        visited[idx] = True
        opr, arg = commands[idx].operation, commands[idx].argument
        if opr == "nop":
            idx += 1
        elif opr == "acc":
            accumulator += arg
            idx += 1
        elif opr == "jmp":
            idx += arg
    return accumulator


def solve_second():
    commands = parse_inputs_to_commands(input_list)
    for cmd in commands:
        if not cmd.convert_operation():
            continue
        visited = {}
        idx = 0
        accumulator = 0
        is_continue = False
        while idx < len(commands):
            if idx not in visited:
                visited[idx] = True
            else:
                is_continue = True
                break
            visited[idx] = True
            opr, num = commands[idx].operation, commands[idx].argument
            if opr == "nop":
                idx += 1
            elif opr == "acc":
                accumulator += num
                idx += 1
            elif opr == "jmp":
                idx += num
        if is_continue:
            cmd.convert_operation()
            continue
        return accumulator


def parse_inputs_to_commands(input_list):
    commands = []
    for i in input_list:
        opr, arg = i.split()
        commands.append(Command(opr, int(arg)))
    return commands


if __name__ == "__main__":
    print(solve_first())
    print(solve_second())
