import sys

def greedy_algorithm(formula):
    result = 0
    for i in range(len(formula)):
        formula[i] = sum(list(map(int, formula[i].split('+'))))
        if i == 0:
            result = formula[i]
        else:
            result = result - formula[i]
    return result

if __name__ == "__main__":
    formula = sys.stdin.readline().split('-')
    result = greedy_algorithm(formula)
    print(result)