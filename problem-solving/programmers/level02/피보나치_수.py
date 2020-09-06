# https://programmers.co.kr/learn/courses/30/lessons/12945
def solution(n):
    fib = {}
    fib[0] = 0
    fib[1] = 1
    for i in range(2, n+1):
        fib[i] = fib[i-1] + fib[i-2]
    return fib[n]%1234567

# if __name__ == "__main__":
#     print(solution(3))
#     print(solution(5))