# https://programmers.co.kr/learn/courses/30/lessons/60057
def solution(s):
    min_len = len(s)
    for chunk_size in range(1, len(s)+1):
        splited = [s[i:i+chunk_size] for i in range(0, len(s), chunk_size)]
        prev = ''
        count = 1
        answer = []
        for j in range(0, len(splited)+1):
            if j == 0:
                prev = splited[j]
                answer.append(prev)
            elif j <len(splited):
                cur = splited[j]
                if prev == cur:
                    count += 1
                else:
                    if count != 1:
                        answer[-1] = str(count)+answer[-1]
                    answer.append(cur)
                    count = 1
                prev = cur
            else:
                if count != 1:
                    answer[-1] = str(count)+answer[-1]
        if min_len > len(''.join(answer)):
            min_len = len(''.join(answer))
    return min_len

# if __name__ == "__main__":
#     print(solution("aabbaccc"))
#     print(solution("ababcdcdababcdcd"))
#     print(solution("abcabcdede"))
#     print(solution("abcabcabcabcdededededede"))
#     print(solution("xababcdcdababcdcd"))