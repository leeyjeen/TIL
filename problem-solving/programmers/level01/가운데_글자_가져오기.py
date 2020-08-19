# https://programmers.co.kr/learn/courses/30/lessons/12903
    
def solution(s):
    answer = ''
    mid_idx = int((len(s)-1)/2)
    if len(s) % 2 != 0:
        answer = s[mid_idx]
    else:
        answer = s[mid_idx] + s[mid_idx+1]
    return answer
        
# if __name__ == "__main__":
#     print(solution('abcde'))
#     print(solution('qwer'))