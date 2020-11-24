# https://www.acmicpc.net/problem/10815
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    cards = {}
    cards_inputs = list(map(int, sys.stdin.readline().split()))
    for i in cards_inputs:
        cards[i] = 1

    m = int(sys.stdin.readline())
    num_inputs = list(map(int, sys.stdin.readline().split()))
    results = []
    for i, v in enumerate(num_inputs):
        if cards.get(v):
            print(1, end=' ')
        else:
            print(0, end=' ')
        if i == m-1:
            print()