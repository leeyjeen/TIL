# https://programmers.co.kr/learn/courses/30/lessons/12947

def solution(x):
    if x % sum(map(int, list(str(x)))) == 0:
        return True
    return False

# if __name__ == "__main__":
#     print(solution(10))
#     print(solution(12))
#     print(solution(11))
#     print(solution(13))