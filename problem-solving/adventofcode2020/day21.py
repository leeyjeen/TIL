# https://adventofcode.com/2020/day/21
from collections import defaultdict
from copy import deepcopy
from more_itertools import flatten

with open("./day21.txt") as f:
    lines = f.read().splitlines()


def solve_part1():
    ingredients = [line.split(" (contains ")[0].split() for line in lines]
    allergens = [line.split(" (contains ")[1][:-1].split(", ")
                 for line in lines]
    foods = list(zip(ingredients, allergens))
    possible_allergens = sorted(set(list(flatten(allergens))))
    possible_ingredients = []
    for allergen in possible_allergens:
        pi = []
        for ingres, algs in foods:
            if allergen in algs:
                pi.append(set(ingres))
        possible_ingredients.append(pi)

    common_possible_ingredients = [
        set.intersection(*ps) for ps in possible_ingredients]
    bad_ingredients = sorted(set(flatten(common_possible_ingredients)))
    return len([f for f in flatten(ingredients)
                if f not in bad_ingredients])


def solve_part2():
    ingredients = [line.split(" (contains ")[0].split() for line in lines]
    allergens = [line.split(" (contains ")[1][:-1].split(", ")
                 for line in lines]
    foods = list(zip(ingredients, allergens))
    possible_allergens = sorted(set(list(flatten(allergens))))
    possible_ingredients = []
    for allergen in possible_allergens:
        pi = []
        for ingres, algs in foods:
            if allergen in algs:
                pi.append(set(ingres))
        possible_ingredients.append((allergen, pi))
    common_possible_ingredients = [
        (ps[0], set.intersection(*ps[1])) for ps in possible_ingredients]
    copied_common_possible_ingredients = deepcopy(common_possible_ingredients)
    result_dict = defaultdict(str)
    idx = 0
    while idx < 5:
        for i, c in enumerate(common_possible_ingredients):
            if len(c[1]) == 0:
                continue
            if len(c[1]) == 1:
                result_dict[c[0]
                            ] = copied_common_possible_ingredients[i][1].pop()
            else:
                for val in c[1]:
                    if val in result_dict.values():
                        copied_common_possible_ingredients[i][1].discard(val)
                if len(copied_common_possible_ingredients[i][1]) == 1:
                    result_dict[c[0]
                                ] = copied_common_possible_ingredients[i][1].pop()
        common_possible_ingredients = deepcopy(
            copied_common_possible_ingredients)
        idx += 1

    return ",".join([i[1] for i in sorted(result_dict.items())])


if __name__ == "__main__":
    print(solve_part1())
    print(solve_part2())
