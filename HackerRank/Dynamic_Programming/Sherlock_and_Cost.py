import sys

def cost(b):
    max_values = [[0]*2 for _ in range(len(b))]
    for i in range(1, len(b)):
        max_values[i][0] = max(max_values[i-1][0], max_values[i-1][1]+b[i-1]-1)
        max_values[i][1] = max(max_values[i-1][0]+b[i]-1, max_values[i-1][1]+abs(b[i-1]-b[i]))
    return max(max_values[len(b)-1][0], max_values[len(b)-1][1])

if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(0, t):
        n = int(sys.stdin.readline())
        b = list(map(int, sys.stdin.readline().split()))
        print(cost(b))
