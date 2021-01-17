# https://www.acmicpc.net/problem/14916
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    num_of_five, num_of_two = 0, 0
    is_available = False

    while n >= 0:
        if n%5 == 0:
            num_of_five = n // 5
            is_available = True
            break
        else:
            n -= 2
            num_of_two += 1
    
    if is_available:
        print(num_of_five + num_of_two)
    else:
        print(-1)