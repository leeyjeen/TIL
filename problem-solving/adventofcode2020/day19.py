# https://adventofcode.com/2020/day/19
import re

with open('./day19.txt', 'r') as f:
    rules, values = f.read().split("\n\n")
    rules = rules.split("\n")
    values = values.split("\n")
    rules_dict = dict(rule.split(": ") for rule in rules)

rule_to_regex_dict = {}


def convert_rules_to_regex(rule_value, rule_to_regex_dict):
    if rule_value in rule_to_regex_dict:
        return rule_to_regex_dict[rule_value]

    regex = ""
    if "\"" in rule_value:
        regex = rule_value.replace("\"", "")
    elif "|" in rule_value:
        opt1, opt2 = rule_value.split(" | ")
        regex = "("+convert_rules_to_regex(opt1, rule_to_regex_dict) + \
            "|"+convert_rules_to_regex(opt2, rule_to_regex_dict)+")"
    else:
        for v in rule_value.split(" "):
            regex += convert_rules_to_regex(rules_dict[v], rule_to_regex_dict)
    rule_to_regex_dict[rule_value] = regex

    return regex


def count_matched_values(reg_pattern, values):
    s = 0
    for v in values:
        if re.match(reg_pattern, v):
            s += 1
    return s


def solve_part1():
    reg_pattern = "^" + \
        convert_rules_to_regex(rules_dict["0"], rule_to_regex_dict)+"$"
    return count_matched_values(reg_pattern, values)


if __name__ == "__main__":
    print(solve_part1())
