# https://programmers.co.kr/learn/courses/30/lessons/12918

def solution(s):
    return (len(s) == 4 or len(s) == 6) and s.isnumeric()

# if __name__ == "__main__":
#     print(solution('a234'))
#     print(solution('1234'))