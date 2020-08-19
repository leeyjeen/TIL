# https://programmers.co.kr/learn/courses/30/lessons/17681
# 2018 KAKAO BLIND RECRUITMENT

def solution(n, arr1, arr2):
    answer = []
    for i, j in zip(arr1, arr2):
        row = str(bin(i|j)[2:].zfill(n)) # 비트연산자 사용
        row = row.replace('1', '#')
        row = row.replace('0', ' ')
        answer.append(row)
    return answer

# def solution(n, arr1, arr2):
#     answer = []
#     for i, j in zip(arr1, arr2):
#         new_i = bin(i)[2:].zfill(n)
#         new_j = bin(j)[2:].zfill(n)
#         elem = ''
#         for ii, jj in zip(new_i, new_j):
#             if ii == '0' and jj == '0':
#                 elem = elem + ' '
#             else:
#                 elem = elem + '#'
#         answer.append(elem)
#     return answer

# if __name__ == "__main__":
#     print(solution(5, [9,20,28,18,11], [30,1,21,17,28]))
#     print(solution(6, [46,33,33,22,31,50], [27,56,19,14,14,10]))