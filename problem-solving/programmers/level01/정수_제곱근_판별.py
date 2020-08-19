# https://programmers.co.kr/learn/courses/30/lessons/12934

import math

def solution(n):
    x = math.sqrt(n)
    if (x*10)%10 == 0:
        return int((x+1)**2)
    return -1

# if __name__ == "__main__":
#     print(solution(121))
#     print(solution(3))