# https://programmers.co.kr/learn/courses/30/lessons/12919

def solution(seoul):
    return "김서방은 {}에 있다".format(seoul.index('Kim'))

# def solution(seoul):
#     for i, v in enumerate(seoul):
#         if v == "Kim":
#             return "김서방은 {}에 있다".format(i)

# if __name__ == "__main__":
#     print(solution(["Jane", "Kim"]))