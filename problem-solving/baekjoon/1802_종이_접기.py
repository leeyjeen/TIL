# https://www.acmicpc.net/problem/1802
import sys

def check(rules):
    if len(rules) == 1:
        return True
    for i in range(len(rules)//2):
        if rules[i] == rules[len(rules)-1-i]:
            return False
    return check(rules[0:len(rules)//2]) and check(rules[len(rules)//2+1:len(rules)])

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        rules = list(sys.stdin.readline().rstrip())
        if check(rules):
            print("YES")
        else:
            print("NO")