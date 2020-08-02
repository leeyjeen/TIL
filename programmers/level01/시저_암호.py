# https://programmers.co.kr/learn/courses/30/lessons/12926

def solution(s, n):
    s = list(s)
    for i, v in enumerate(s):
        ascii_code = ord(v)
        convert = ascii_code + n
        if v.istitle() and convert > 90 or not v.istitle() and convert > 122:
            convert -= 26
        elif ascii_code == 32:
            convert = ascii_code
        s[i] = chr(convert)
    return ''.join(s)
    
# if __name__ == "__main__":
#     print(solution("AB", 1))
#     print(solution("z", 1))
#     print(solution("a B z", 4))