import sys

def greedy_algorithm(n, p):
    p.sort()
    min_times = [p[0]]
    for i in range(1, n):
        min_times.append(min_times[i-1]+p[i])
    return sum(min_times)
    

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    p = list(map(int, sys.stdin.readline().split()))
    print(greedy_algorithm(n, p))