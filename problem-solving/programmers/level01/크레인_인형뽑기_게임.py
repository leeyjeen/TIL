# https://programmers.co.kr/learn/courses/30/lessons/64061
# 2019 카카오 개발자 겨울 인턴십

def solution(board, moves):
    count = 0
    result = []
    for i in moves:
        for row in board:
            if row[i-1] == 0:
                continue
            if len(result) > 0 and result[-1] == row[i-1]:
                result.pop()
                count += 2
                while len(result)>1 and result[-1] == result[-2]:
                    result.pop()
                    result.pop()
                    count += 2
                row[i-1] = 0
            else:
                result.append(row[i-1])
                row[i-1] = 0
            break
    return count

# if __name__ == "__main__":
#     print(solution([[0,0,0,0,0],[0,0,1,0,3],[0,2,5,0,1],[4,2,4,4,2],[3,5,1,3,1]], [1,5,3,5,1,2,1,4]))