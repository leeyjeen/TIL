# https://programmers.co.kr/learn/courses/30/lessons/12901

def solution(a, b):
    days = [31,29,31,30,31,30,31,31,30,31,30,31]
    day_of_week = ['THU','FRI','SAT','SUN','MON','TUE','WED']
    total_days = b
    for i in range(0, a-1):
        total_days += days[i]
    answer = day_of_week[total_days%7]
    return answer
    
# if __name__ == "__main__":
#     print(solution(5, 24))