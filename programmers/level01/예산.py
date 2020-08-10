# https://programmers.co.kr/learn/courses/30/lessons/12982
# Summer/Winter Coding(~2018)

def solution(d, budget):
    answer = 0
    for i in sorted(d):
        budget -= i
        if budget < 0:
            break
        answer += 1
    return answer

# if __name__ == "__main__":
#     print(solution([1,3,2,5,4], 9))
#     print(solution([2,2,3,3], 10))
