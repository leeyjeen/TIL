# https://programmers.co.kr/learn/courses/30/lessons/12915

def solution(strings, n):
    return sorted(strings, key=lambda x:(x[n], x))

# if __name__ == "__main__":
#     print(solution(["sun", "bed", "car"], 1))
#     print(solution(["abce", "abcd", "cdx"], 2))