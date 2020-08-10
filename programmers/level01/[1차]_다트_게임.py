# https://programmers.co.kr/learn/courses/30/lessons/17682
# 2018 KAKAO BLIND RECRUITMENT

def solution(dartResult):
    answer = 0
    bonus = ['S', 'D', 'T']
    option = ['*', '#']
    dart_list = str(dartResult)
    result = []
    for i, v in enumerate(dart_list):
        if v.isnumeric():
            if i>0 and dart_list[i-1].isnumeric():
                result[-1] = result[-1]*10 + int(v)
            else:
                result.append(int(v))
        elif v in bonus:
            result[-1] = result[-1]**(bonus.index(v)+1)
        elif v in option:
            if v == '*':
                result[-1] *= 2
                if len(result) > 1:
                    result[-2] *= 2
            elif v == '#':
                result[-1] *= -1
    return sum(result)

# if __name__ == "__main__":
#     print(solution("1S2D*3T"))
#     print(solution("1D2S#10S"))
#     print(solution("1D2S0T"))
#     print(solution("1S*2T*3S"))
#     print(solution("1D#2S*3S"))
#     print(solution("1T2D3D#"))
#     print(solution("1D2S3T*"))
#     print(solution("10D2S3T*"))