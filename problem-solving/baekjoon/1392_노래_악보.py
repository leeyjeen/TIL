# https://www.acmicpc.net/problem/1392
import sys

if __name__ == "__main__":
    n, q = list(map(int, sys.stdin.readline().split()))
    scores = []
    for i in range(n):
        seconds = int(sys.stdin.readline())
        if i == 0:
            scores.append(seconds)
        else:
            scores.append(scores[i-1]+seconds)

    questions = []
    for i in range(q):
        seconds = int(sys.stdin.readline())
        questions.append(seconds)
        
    for i in range(q):
        question = questions[i]
        for j in range(n):
            if scores[j] > question:
                print(j+1)
                break