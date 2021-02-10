# https://www.acmicpc.net/problem/1049
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    min_package, min_single = 1000, 1000
    for i in range(m):
        package_price, single_price = list(map(int, sys.stdin.readline().split()))
        min_package = min(min_package, package_price)
        min_single = min(min_single, single_price)
    money = 0
    if min_single*6 > min_package:
        money += (n // 6) * min_package
        n %= 6
        if n > 0:
            if min_single * n > min_package:
                money += min_package
            else:
                money += min_single*n
    else:
        money += n * min_single
    print(money)