# https://www.acmicpc.net/problem/3036
import sys

def get_gcm(a, b):
    if b==0:
        return a
    a, b = b, a%b
    return get_gcm(a, b)

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    inputs = list(map(int, sys.stdin.readline().split()))
    for i in range(1, n):
        gcm = get_gcm(inputs[0], inputs[i])
        print("%d/%d" % (inputs[0]/gcm, inputs[i]/gcm))