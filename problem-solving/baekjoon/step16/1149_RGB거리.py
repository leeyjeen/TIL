import sys

def calculate(n):
    min_val = 0
    if n == 1:
        min_val = min(cost[0][0], cost[0][1], cost[0][2])
    min_costs = [[cost[0][0], cost[0][1], cost[0][2]]]
    for i in range(1, n):
        min_costs.append([min(min_costs[i-1][1], min_costs[i-1][2])+cost[i][0],
                          min(min_costs[i-1][0], min_costs[i-1][2])+cost[i][1],
                          min(min_costs[i-1][0], min_costs[i-1][1])+cost[i][2]])
    min_val = min(min_costs[n-1][0], min_costs[n-1][1], min_costs[n-1][2])
    return min_val

if __name__ == "__main__":
    n = int(sys.stdin.readline())

    cost = []
    for i in range(n):
        a, b, c = map(int, sys.stdin.readline().split())
        cost.append([a, b, c])

    min_val = calculate(n)
    print(min_val)