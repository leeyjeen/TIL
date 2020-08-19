# https://programmers.co.kr/learn/courses/30/lessons/12912

def solution(a, b):
    answer, start, end = 0, 0, 0
    if a>b:
        start, end = b, a
    else:
        start, end = a, b
    for i in range(start, end+1):
        answer += i
    return answer

# if __name__ == "__main__":
#     print(solution(3,5))
#     print(solution(3,3))
#     print(solution(5,3))