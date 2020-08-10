# https://programmers.co.kr/learn/courses/30/lessons/67256
# 2020 카카오 인턴십

def solution(numbers, hand):
    answer = ''
    lefts = [1,4,7]
    rights = [3,6,9]
    pos = {
        1: (0,0),2: (0,1),3: (0,2),
        4: (1,0),5: (1,1),6: (1,2),
        7: (2,0),8: (2,1),9: (2,2),
        '*': (3,0),0: (3,1),'#': (3,2)
    }
    left_prev = '*'
    right_prev = '#'
    for v in numbers:
        if v in lefts:
            answer += 'L'
            left_prev = v
        elif v in rights:
            answer += 'R'
            right_prev = v
        else:
            left_dist = get_distance(pos[left_prev], pos[v])
            right_dist = get_distance(pos[right_prev], pos[v])
            if left_dist < right_dist:
                answer += 'L'
                left_prev = v
            elif right_dist < left_dist:
                answer += 'R'
                right_prev = v
            else:
                if hand == 'left':
                    answer += 'L'
                    left_prev = v
                else:
                    answer += 'R'
                    right_prev = v
    return answer

def get_distance(p1, p2):
    return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
    
# if __name__ == "__main__":
#     print(solution([1, 3, 4, 5, 8, 2, 1, 4, 5, 9, 5], "right"))
#     print(solution([7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2], "left"))
#     print(solution([1, 2, 3, 4, 5, 6, 7, 8, 9, 0], "right"))