# https://programmers.co.kr/learn/courses/30/lessons/12940

def solution(n, m):
    answer = []
    smaller, bigger = min(n,m), max(n,m)
    for i in reversed(range(1, smaller+1)):
        if bigger % i == 0 and smaller % i == 0:
            answer.append(i)
            break
    for i in range(bigger, smaller*bigger+1, bigger):
        if i % smaller == 0:
            answer.append(i)
            break
    return answer

# if __name__ == "__main__":
#     print(solution(3, 12))
#     print(solution(2, 5))