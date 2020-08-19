import sys

def calculate(n):
    if n == 1:
        return values[0][0]
    max_values = [[values[0][0]]]
    idx = 0
    for i in range(1, n):
        val = []
        for j in range(0, i+1):
            if j-1<0:
                max_val = max_values[i-1][j]
            elif j>i-1:
                max_val = max_values[i-1][j-1]
            else:
                max_val = max(max_values[i-1][j], max_values[i-1][j-1])
                
            val.append(max_val+values[i][j])
        max_values.append(val)
    return max(max_values[n-1])


if __name__ == "__main__":
    n = int(sys.stdin.readline())
    values = []
    for i in range(n):
        values.append(list(map(int, sys.stdin.readline().split())))
    max_value = calculate(n)
    print(max_value)
