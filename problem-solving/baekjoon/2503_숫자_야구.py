# https://www.acmicpc.net/problem/2503
import sys

def check_valid(a, b, c):
    if b == 0 or c == 0:
        return False
    if a == b or b == c or c == a:
        return False
    return True

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    games = []
    for i in range(n):
        games.append(list(map(int, sys.stdin.readline().split())))
    
    count = 0
    for i in range(123, 988):
        is_answer = True
        a, b, c = i//100, (i//10)%10, i%10
        if not check_valid(a, b, c):
            continue

        for j in range(0, n):
            strike, ball, number = 0, 0, games[j][0]
            a2, b2, c2 = number//100, number//10%10, number%10
            if a == a2:
                strike += 1
            if a == b2 or a == c2:
                ball += 1
            if b == b2:
                strike += 1
            if b == a2 or b == c2:
                ball += 1
            if c == c2:
                strike += 1
            if c == a2 or c == b2:
                ball += 1
            if not (strike == games[j][1] and ball == games[j][2]):
                is_answer = False
                break
        if is_answer:
            count += 1
    print(count)