# https://programmers.co.kr/learn/courses/30/lessons/12922

def solution(n):
    answer = []
    for i in range(0,n):
        if i%2 == 0:
            answer.append('수')
        else:
            answer.append('박')
    return ''.join(answer)

# if __name__ == "__main__":
#     print(solution(3))
#     print(solution(4))