# https://programmers.co.kr/learn/courses/30/lessons/12931

def solution(n):
    sum = n%10
    while n != 0:
        n = n//10
        sum += n%10
    return sum

# if __name__ == "__main__":
#     print(solution(123))
#     print(solution(987))