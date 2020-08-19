# https://programmers.co.kr/learn/courses/30/lessons/12950

def solution(arr1, arr2):
    answer = arr1
    for i, a2 in enumerate(arr2):
        for j, v in enumerate(a2):
            answer[i][j] += a2[j]
    return answer

# if __name__ == "__main__":
#     print(solution([[1,2],[2,3]], [[3,4],[5,6]]))
#     print(solution([[1],[2]],[[3],[4]]))