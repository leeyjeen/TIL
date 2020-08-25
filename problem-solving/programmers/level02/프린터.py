# https://programmers.co.kr/learn/courses/30/lessons/42587

def solution(priorities, location):
    queue =  [(i,p) for i,p in enumerate(priorities)]
    answer = 0
    while True:
        cur = queue.pop(0)
        if any(cur[1] < q[1] for q in queue):
            queue.append(cur)
        else:
            answer += 1
            if cur[0] == location:
                return answer

# def solution(priorities, location):
#     tuples = [(i,p) for i,p in enumerate(priorities)] # [(0, 2), (1, 1), (2, 3), (3, 2)]
#     index = 0
#     rest_count = len(priorities)
#     result = []
#     while True:
#         if index >= len(tuples) or rest_count == 0:
#             break
#         cur_tuple = tuples[index]
#         cur_val = cur_tuple[1]
#         if cur_val >= max(priorities):
#             priorities[cur_tuple[0]] = 0
#             result.append(cur_tuple)
#             rest_count -= 1
#         else:
#             tuples.append(cur_tuple)
#         index += 1
#     return result.index(tuples[location])+1

# if __name__ == "__main__":
#     print(solution([2,1,3,2],2))
#     print(solution([1,1,9,1,1,1],0))
    