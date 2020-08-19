# https://programmers.co.kr/learn/courses/30/lessons/12899

def solution(n):
    mod = n%3
    result = ""
    if mod == 1:
        result = "1"
    elif mod == 2:
        result = "2"
    elif mod == 0:
        result = "4"
    portion = (n-1)//3
    multiple = 10
    while portion >= 1:
        if (portion)%3 == 1:
            result = "1" + result
        elif (portion)%3 == 2:
            result = "2" + result
        else:
            result = "4" + result
        portion = (portion-1) // 3
    return result

# if __name__ == "__main__":
#     print(solution(8))
#     print(solution(9))
#     print(solution(15))