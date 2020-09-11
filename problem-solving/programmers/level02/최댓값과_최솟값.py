# https://programmers.co.kr/learn/courses/30/lessons/12939
def solution(s):
    numbers = sorted(list(map(int, s.split())))
    return str(numbers[0])+" "+str(numbers[-1])

# if __name__ == "__main__":
#     print(solution("1 2 3 4"))
#     print(solution("-1 -2 -3 -4"))