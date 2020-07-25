import sys

def fibonacci(t1,t2,n):
    f = [0]*n
    f[0], f[1] = t1, t2
    for i in range(2, n):
        f[i] = f[i-2] + f[i-1]*f[i-1]
    return f[n-1]

if __name__ == "__main__":
    t1, t2, n = map(int, sys.stdin.readline().split())
    print(fibonacci(t1,t2,n))