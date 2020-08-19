import sys

min_burger, min_beverage = 2000, 2000

for i in range(3):
    price = int(sys.stdin.readline())
    if price < min_burger:
        min_burger = price

for i in range(2):
    price = int(sys.stdin.readline())
    if price < min_beverage:
        min_beverage = price
        
print(min_burger + min_beverage - 50)