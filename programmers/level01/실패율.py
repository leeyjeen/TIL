# https://programmers.co.kr/learn/courses/30/lessons/42889
# 2019 KAKAO BLIND RECRUITMENT

def solution(N, stages):
    answer = []
    dic = {}
    for i in stages:
        if dic.get(i):
            dic[i] += 1
        else:
            dic[i] = 1
    fail = []
    reached = len(stages)
    for i in range(1, N+1):
        if not dic.get(i):
            fail.append(0)
            continue
        fail.append(dic[i]/reached)
        reached -= dic[i]
    return [x+1 for x in sorted(range(0, len(fail)), key=lambda k: fail[k], reverse=True)]

# if __name__ == "__main__":
#     print(solution(5, [2,1,2,6,2,4,3,3]))
#     print(solution(4, [4,4,4,4,4]))