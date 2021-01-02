# https://www.acmicpc.net/problem/18243
import sys

def floyd_warshall(friends):
    for i in range(len(friends)):
        for j in range(len(friends)):
            for k in range(len(friends)):
                if j == k:
                    continue
                if friends[j][i] != 0 and friends[i][k] != 0:
                    if friends[j][k] == 0 or friends[j][k] > friends[j][i]+friends[i][k]:
                        friends[j][k] = friends[j][i] + friends[i][k]
    return friends

def check_network(friends):
    for i in range(len(friends)):
        for j in range(len(friends[i])):
            if friends[i][j] == 0 and i != j or friends[i][j] > 6:
                return True
    return False
            

if __name__ == "__main__":
    n, k = list(map(int, sys.stdin.readline().split()))
    friends = [[0 for i in range(n)] for j in range(n)]
    for i in range(k):
        a, b = list(map(int, sys.stdin.readline().split()))
        friends[a-1][b-1] = 1
        friends[b-1][a-1] = 1

    friends = floyd_warshall(friends)
    is_big = check_network(friends)
    if is_big:
        print("Big World!")
    else:
        print("Small World!")