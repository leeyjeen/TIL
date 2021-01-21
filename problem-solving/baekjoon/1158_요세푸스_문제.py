# https://www.acmicpc.net/problem/1158
import sys

def josephus(n, k):
    queue = [i for i in range(1, n+1)]
    index = 0
    result = []
    while len(queue) > 0:
        index = (index + k-1) % len(queue)
        result.append(queue[index])
        queue.pop(index)
    return result

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    result = josephus(n, k)
    for i in range(len(result)):
        if i == 0:
            print("<{}".format(result[i]), end='')
        else:
            print(", {}".format(result[i]), end='')
    print(">")