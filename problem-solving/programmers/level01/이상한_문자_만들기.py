# https://programmers.co.kr/learn/courses/30/lessons/12930

def solution(s):
    s = s.split(' ')
    v = []
    for i, word in enumerate(s):
        for j, letter in enumerate(word):
            if j%2 == 0:
                v.append(letter.upper())
            else:
                v.append(letter.lower())
        if i < len(s)-1:
            v.append(' ')
    return ''.join(v)
    
# if __name__ == "__main__":
#     print(solution("Where   is My Telephone"))