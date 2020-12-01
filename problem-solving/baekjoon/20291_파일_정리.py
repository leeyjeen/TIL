# https://www.acmicpc.net/problem/20291
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    ext_dict = {}
    for i in range(n):
        _, ext = sys.stdin.readline().rstrip().split(".")
        if ext_dict.get(ext):
            ext_dict[ext] += 1
        else:
            ext_dict[ext] = 1
    for key in sorted(ext_dict.keys()):
        print('{} {}'.format(key, ext_dict[key]))