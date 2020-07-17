import sys

def calculate(n):
    lengths = [1, 1, 1]
    if n <= 3:
        return lengths
    for i in range(3, n+1):
        lengths.append(lengths[i-2]+lengths[i-3])
    return lengths
        
if __name__ == "__main__":
    t = int(sys.stdin.readline())
    for i in range(t):
        n = int(sys.stdin.readline())
        lengths = calculate(n)
        print(lengths[n-1])