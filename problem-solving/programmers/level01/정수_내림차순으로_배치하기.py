# https://programmers.co.kr/learn/courses/30/lessons/12933

def solution(n):
    return int(''.join(list(reversed(sorted(str(n))))))

# if __name__ == "__main__":
#     print(solution(118372))