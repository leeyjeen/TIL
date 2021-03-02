# https://www.acmicpc.net/problem/5883
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    b = []
    keys = {}
    for _ in range(n):
        input = int(sys.stdin.readline())
        b.append(input)
        keys[input] = True
    max_continuous_count = 1
    for key_to_delete in keys:
        new_line = []
        for i in range(n):
            if b[i] != key_to_delete:
                new_line.append(b[i])
        if len(new_line) == 0:
            max_continuous_count = 0
            break
        prev = new_line[0]
        continuous_count = 1
        for i in range(1, len(new_line)):
            cur = new_line[i]
            if cur == prev:
                continuous_count += 1
                max_continuous_count = max(max_continuous_count, continuous_count)
            else:
                continuous_count = 1
            prev = cur
    print(max_continuous_count)