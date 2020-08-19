# https://programmers.co.kr/learn/courses/30/lessons/12906

def solution(arr):
    answer = []
    last = -1
    for i in arr:
        if last == i:
            continue
        answer.append(i)
        last = i
    return answer

# if __name__ == "__main__":
#     print(solution([1,1,3,3,0,1,1]))
#     print(solution([4,4,4,3,3]))