# https://programmers.co.kr/learn/courses/30/lessons/12916

def solution(s):
    return s.lower().count('p') == s.lower().count('y')

# def solution(s):
#     count_p, count_y = 0, 0
#     for letter in s:
#         if letter == "p" or letter =="P":
#             count_p += 1
#         elif letter == "y" or letter == "Y":
#             count_y += 1
#     return count_p == count_y

# if __name__ == "__main__":
#     print(solution("pPoooyY"))
#     print(solution("Pyy"))