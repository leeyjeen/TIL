# https://programmers.co.kr/learn/courses/30/lessons/12951
def solution(s):
    words = s.split(' ')
    for i, v in enumerate(words):
        words[i] = v.lower()
        if len(v) == 0:
            continue
        if v[0].isalpha():
            if len(v) > 1:
                words[i] = words[i][0].upper() + words[i][1:]
            else:
                words[i] = words[i].upper()
    return ' '.join(words)

# if __name__ == "__main__":
#     print(solution("3people unFollowed me"))
#     print(solution("for the  last week"))