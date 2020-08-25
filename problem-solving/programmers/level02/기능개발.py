# https://programmers.co.kr/learn/courses/30/lessons/42586
import math
def solution(progresses, speeds):
    wait = []
    deploy = []
    
    for i, (a,b) in enumerate(zip(progresses, speeds)):
        if i > 0 and math.ceil((100-a)/b) < wait[i-1]:
            wait.append(wait[i-1])
        else:
            wait.append(math.ceil((100-a)/b))

    for i in range(0, len(wait)):
        if i == 0:
            deploy.append(1)
        else:
            if wait[i-1] == wait[i]:
                deploy[-1] += 1
            else:
                deploy.append(1)
    return deploy

# if __name__ == "__main__":
#     print(solution([93,30,55], [1,30,5]))