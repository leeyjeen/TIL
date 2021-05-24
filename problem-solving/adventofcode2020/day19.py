# https://adventofcode.com/2020/day/19
import re


def convert_rules_to_regex(rule_value, rules_dict, rule_to_regex_dict):
    if rule_value in rule_to_regex_dict:
        return rule_to_regex_dict[rule_value]

    regex = ""
    if '"' in rule_value:
        regex = rule_value.replace("\"", "")
    elif "|" in rule_value:
        opt1, opt2 = rule_value.split(" | ")
        regex = "("+convert_rules_to_regex(opt1, rules_dict, rule_to_regex_dict) + \
            "|"+convert_rules_to_regex(opt2, rules_dict,
                                       rule_to_regex_dict)+")"
    else:
        for v in rule_value.split(" "):
            regex += convert_rules_to_regex(
                rules_dict[v], rules_dict, rule_to_regex_dict)
    rule_to_regex_dict[rule_value] = regex

    return regex


def count_matched_values(reg_pattern, values):
    s = 0
    for v in values:
        if re.match(reg_pattern, v):
            s += 1
    return s


def count_matched_values2(reg_pattern, values):
    s = 0
    for v in values:
        match = re.match(reg_pattern, v)
        if match:
            if len(match.group(1)) > len(match.group(2)):
                s += 1
    return s


def solve_part1():
    rules, values = open('./day19.txt', 'r').read().split("\n\n")
    rules = rules.split("\n")
    values = values.split("\n")
    rules_dict = dict(rule.split(": ") for rule in rules)
    rule_to_regex_dict = {}
    reg_pattern = "^" + convert_rules_to_regex(
        rules_dict["0"], rules_dict, rule_to_regex_dict)+"$"
    return count_matched_values(reg_pattern, values)


def solve_part2():
    rules, values = open('./day19.txt', 'r').read().split("\n\n")
    rules = rules.split("\n")
    values = values.split("\n")
    rules_dict = dict(rule.split(": ") for rule in rules)
    rules_dict["8"] = "42 | 42 8"
    rules_dict["11"] = "42 31 | 42 11 31"
    rule_to_regex_dict = {}
    convert_rules_to_regex(  # regex계산하는 함수가 잘못된 것인지 답이 이상하게 나온다...모르겠음
        rules_dict["42"], rules_dict, rule_to_regex_dict)
    convert_rules_to_regex(
        rules_dict["31"], rules_dict, rule_to_regex_dict)
    return count_matched_values2(re.compile('^('+rule_to_regex_dict[rules_dict["42"]]+'+)('+rule_to_regex_dict[rules_dict["31"]]+'+)$'), values)


if __name__ == "__main__":
    print(solve_part1())
    print(solve_part2())
