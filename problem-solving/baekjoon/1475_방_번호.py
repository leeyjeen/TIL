# https://www.acmicpc.net/problem/1475
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    number_count = {}
    if n == 0:
        print(1)
        sys.exit()
    number_count = {}
    while n > 0:
        if number_count.get(n%10):
            number_count[n%10] += 1
        else:
            number_count[n%10] = 1
        n //= 10
    six_count = number_count.get(6) if number_count.get(6) else 0
    nine_count = number_count.get(9) if number_count.get(9) else 0
    if (six_count+nine_count)%2 == 0:
        number_count[6] = (six_count+nine_count)//2
        number_count[9] = number_count[6]
    else:
        number_count[6] = (six_count+nine_count)//2
        number_count[9] = number_count[6] + 1
    set_count = max(number_count.values())
    print(set_count)