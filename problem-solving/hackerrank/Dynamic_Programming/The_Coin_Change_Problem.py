import sys

def get_ways(n, coins):
    ways = [0]*(n+1)
    ways[0] = 1
    for _, coin in enumerate(coins):
        for i in range(1, n+1):
            if coin <= i:
                ways[i] += ways[i-coin]
    return ways[n]

if __name__ == "__main__":
    n, m = map(int, sys.stdin.readline().split())
    coins = list(map(int, sys.stdin.readline().split()))
    print(get_ways(n, coins))