# https://programmers.co.kr/learn/courses/30/lessons/12948

def solution(phone_number):
    return "*"*len(phone_number[:-4]) + phone_number[-4:]
    
# if __name__ == "__main__":
#     print(solution("01033334444"))
#     print(solution("027778888"))