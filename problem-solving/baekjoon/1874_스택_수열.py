# https://www.acmicpc.net/problem/1874
# 스택을 활용하는 문제
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())

    stack_num = 1
    stack_seq = []
    result = []
    for i in range(n):
        val = int(sys.stdin.readline())
        if len(stack_seq) > 0 and val > stack_seq[-1] or val > stack_num:
            for j in range(stack_num, val+1):
                stack_seq.append(j)
                result.append("+")
                stack_num += 1
            stack_seq = stack_seq[:-1]
            result.append("-")
        elif val == stack_num:
            stack_seq.append(val)
            result.append("+")
            stack_num += 1
            stack_seq = stack_seq[:-1]
            result.append("-")
        elif len(stack_seq) > 0 and val == stack_seq[-1]:
            stack_seq = stack_seq[:-1]
            result.append("-")

    if len(stack_seq) > 0:
        print("NO")
    else:
        for v in result:
            print(v)