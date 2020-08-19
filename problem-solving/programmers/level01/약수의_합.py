# https://programmers.co.kr/learn/courses/30/lessons/12928

def solution(n):
    sum = n
    for i in range(1, int(n/2)+1):
        if n%i == 0:
            sum += i
    return sum

# if __name__ == "__main__":
#     print(solution(12))
#     print(solution(5))