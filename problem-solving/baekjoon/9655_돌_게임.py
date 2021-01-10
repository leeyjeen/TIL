# https://www.acmicpc.net/problem/9655
import sys

def get_winner(n):
    if n%2 == 0:
        return "CY"
    return "SK"

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    print(get_winner(n))