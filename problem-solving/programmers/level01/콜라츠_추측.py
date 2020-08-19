# https://programmers.co.kr/learn/courses/30/lessons/12943

def solution(num):
    count = 0
    while num != 1:
        if count >= 500:
            return -1
        if num %2 == 0:
            num /= 2
        else:
            num = num*3+ 1
        count += 1
    return count

# if __name__ == "__main__":
#     print(solution(6))
#     print(solution(16))
#     print(solution(626331))