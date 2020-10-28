import sys

sum = 0
for i in range(5):
    score = int(sys.stdin.readline())
    if score < 40:
        score = 40
    sum += score

print(int(sum/5))