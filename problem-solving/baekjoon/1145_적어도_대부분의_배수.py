# https://www.acmicpc.net/problem/1145
import sys

def get_gcd(first, second):
    if first < second:
        second, first = first, second
    
    while second != 0:
        first, second = second, first%second
    
    return first

def get_lcm(first, second):
    return first*second / get_gcd(first, second)

if __name__ == "__main__":
    numbers = list(map(int, sys.stdin.readline().split()))
    min_lcm = 10000000000
    for i in range(5):
        for j in range(i+1, 5):
            lcm_of_two = get_lcm(numbers[i], numbers[j])
            for k in range(j+1, 5):
                lcm_of_three = get_lcm(lcm_of_two, numbers[k])
                min_lcm = min(min_lcm, lcm_of_three)
    print(int(min_lcm))