# https://programmers.co.kr/learn/courses/30/lessons/12954

def solution(x, n):
    arr = [x]
    for i in range(2, n+1):
        arr.append(x*i)
    return arr

# if __name__ == "__main__":
#     print(solution(2, 5))
#     print(solution(4, 3))
#     print(solution(-4, 2))