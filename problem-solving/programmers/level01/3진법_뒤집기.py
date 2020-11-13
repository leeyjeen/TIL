# https://programmers.co.kr/learn/courses/30/lessons/68935
def solution(n):
    answer = 0
    reversed_tenary = ''
    while n != 0:
        reversed_tenary += str(n%3)
        n = n//3
    for i in range(len(reversed_tenary)):
        answer += int(reversed_tenary[i])*(3**(len(reversed_tenary)-i-1))
    return answer

# if __name__ == "__main__":
#     print(solution(45))
#     print(solution(125))