import sys

def count_binary_sequence(n):
    sequence = [1, 2]
    for i in range(3, n+1):
        sequence.append((sequence[i-2]+sequence[i-3])%15746)
    return sequence

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    sequence = count_binary_sequence(n)
    print(sequence[n-1])