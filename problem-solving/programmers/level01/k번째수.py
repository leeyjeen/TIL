# Sorting
# https://programmers.co.kr/learn/courses/30/lessons/42748

def solution(array, commands):
    answer = []
    for v in commands:
        start, end, aim = v[0]-1, v[1], v[2]-1
        answer.append(sorted(array[start:end])[aim])
    return answer

# if __name__ == "__main__":
#     array = [1,5,2,6,3,7,4]
#     commands = [
#         [2,5,3], [4,4,1], [1,7,3]
#     ]
#     print(solution(array, commands))