# https://programmers.co.kr/learn/courses/30/lessons/12910

def solution(arr, divisor):
    answer = sorted([val for val in arr if val % divisor == 0])
    if len(answer) == 0:
        answer = [-1]
    return answer

# if __name__ == "__main__":
#     print(solution([5,9,7,10],5))
#     print(solution([2,36,1,3],1))
#     print(solution([3,2,6],10))