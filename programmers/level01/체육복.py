# Greedy algorithm
# https://programmers.co.kr/learn/courses/30/lessons/42862

def solution(n, lost, reserve):
    lost_set, reserve_set = set(lost), set(reserve)
    lost_set, reserve_set = lost_set - reserve_set, reserve_set - lost_set
    for r in reserve_set:
        if r-1 in lost_set:
            lost_set.remove(r-1)
        elif r+1 in lost_set:
            lost_set.remove(r+1)
    return n - len(lost_set)

# # failed to pass 7th test case.. 
# def solution(n, lost, reserve):
#     lost_dict, reserve_dict = {}, {}
#     for l in lost:
#         lost_dict.setdefault(l, 1)
#     for r in reserve:
#         reserve_dict.setdefault(r, 1)
    
#     for l in lost_dict:
#         if l in reserve_dict and reserve_dict[l] == 1:
#             reserve_dict[l] = 0
#             lost_dict[l] = 0
#     for l in lost_dict:
#         can_borrow_prev = True if l-1 in reserve_dict and reserve_dict[l-1] == 1 else False
#         can_borrow_next = True if l+1 in reserve_dict and reserve_dict[l+1] == 1 else False
#         if can_borrow_prev and not can_borrow_next:
#             reserve_dict[l-1] = 0
#             lost_dict[l] = 0
#             continue
#         if can_borrow_next and not can_borrow_prev:
#             reserve_dict[l+1] = 0
#             lost_dict[l] = 0
#             continue
#         if can_borrow_prev and can_borrow_next:
#             reserve_dict[l-1] = 0
#             lost_dict[l] = 0
            
#     for _, v in lost_dict.items():
#         if v == 1:
#             n -= 1
#     return n
    
# if __name__ == "__main__":
#     print(solution(5, [2,3,4], [1,3]))