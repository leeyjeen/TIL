# https://www.acmicpc.net/problem/4573
import sys
import math

if __name__ == "__main__":
    menu = 0
    while True:
        menu += 1
        n = int(sys.stdin.readline())
        if n == 0:
            break
        selected_diamter = 0
        low_price_per_square = 10000
        for i in range(n):
            diameter, price = list(map(int, sys.stdin.readline().split()))
            radius = diameter/2.0
            price_per_square = price / (radius*radius*math.pi)
            if price_per_square < low_price_per_square:
                low_price_per_square = price_per_square
                selected_diamter = diameter
        print("Menu {}: {}".format(menu, selected_diamter))