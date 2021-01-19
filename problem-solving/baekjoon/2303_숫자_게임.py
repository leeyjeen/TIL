# https://www.acmicpc.net/problem/2303
import sys

def get_winnder(cards):
    max_sum, winner = 0, 0
    sum_one, sum_two, sum_three = 0, 0, 0
    for i in range(len(cards)):
        for j in range(5):
            sum_one = cards[i][j]
            for k in range(j+1, 5):
                sum_two = sum_one + cards[i][k]
                for l in range(k+1, 5):
                    sum_three = sum_two + cards[i][l]
                if max_sum <= sum_three % 10:
                    max_sum = sum_three % 10
                    winner = i + 1
    return winner

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    cards = []
    for i in range(n):
        cards.append(list(map(int, sys.stdin.readline().split())))
    print(get_winnder(cards))