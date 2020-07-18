import sys

def greedy_algorithm(values, k):
    coin_count = 0
    for i in reversed(range(0, n)):
        coin_count = coin_count + int(k / values[i])
        k %= values[i]
    return coin_count

if __name__ == "__main__":
    n, k = map(int, sys.stdin.readline().split()) 
    values = []
    for i in range(0, n):
        values.append(int(sys.stdin.readline()))
    
    min_count = greedy_algorithm(values, k)
    print(min_count)