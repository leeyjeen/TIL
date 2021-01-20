# https://www.acmicpc.net/problem/5568
import sys
from itertools import permutations

def get_count_of_numbers(cards, k):
    result_dict = {}
    for i in range(len(cards)):
        result_dict = recursion(i, cards, [cards[i]], k, result_dict)
    return len(result_dict)

def recursion(i, cards, cur_list, k, result_dict):
    if len(cur_list) == k:
        permutes = list(permutations(cur_list))
        for i in range(len(permutes)):
            key = ""
            for j in range(len(permutes[i])):
                key += permutes[i][j]
            result_dict[key] = 1
        return result_dict
    for j in range(i+1, len(cards)):
        new_list = cur_list.copy()
        new_list.append(cards[j])
        result_dict = recursion(j, cards, new_list, k, result_dict)
    return result_dict

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    k = int(sys.stdin.readline())
    cards = []
    for i in range(n):
        cards.append(str(sys.stdin.readline().rstrip()))
    print(get_count_of_numbers(cards, k))