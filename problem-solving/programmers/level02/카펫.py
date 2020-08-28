# https://programmers.co.kr/learn/courses/30/lessons/42842
def solution(brown, yellow):
    answer = []
    for i in range(0, yellow):
        portion = yellow // (i+1)
        if portion * (i+1) == yellow and portion*2 + (i+1)*2 + 4 == brown:
            answer.append(portion+2)
            answer.append(i+3)
            break
    return sorted(answer, reverse=True)

# if __name__ == "__main__":
#     print(solution(10, 2))
#     print(solution(8, 1))
#     print(solution(24, 24))