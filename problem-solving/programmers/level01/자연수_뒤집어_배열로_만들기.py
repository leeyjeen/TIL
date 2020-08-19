# https://programmers.co.kr/learn/courses/30/lessons/12932

def solution(n):
    return list(map(int,(reversed(str(n)))))

# if __name__ == "__main__":
#     print(solution(12345))