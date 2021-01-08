# https://www.acmicpc.net/problem/14501
import sys

result = 0
n = 0
infos = []

def get_max(index, profit):
    global result

    if index == n:
        if result < profit:
            result = profit
        return
    
    if index > n:
        return

    get_max(index+infos[index][0], profit+infos[index][1])
    get_max(index+1, profit)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    for i in range(0, n):
        time, profit = list(map(int, sys.stdin.readline().split()))
        infos.append((time, profit))
    get_max(0, 0)
    print(result)