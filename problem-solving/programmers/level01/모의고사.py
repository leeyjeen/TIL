# https://programmers.co.kr/learn/courses/30/lessons/42840

def solution(answers):
    winners = []
    first = [1,2,3,4,5]
    second = [2,1,2,3,2,4,2,5]
    third = [3,3,1,1,2,2,4,4,5,5]
    first_score, second_score, third_score = 0, 0, 0
    for i in range(0,len(answers)):
        answer = answers[i]
        if answer == first[i%len(first)]:
            first_score += 1
        if answer == second[i%len(second)]:
            second_score += 1
        if answer == third[i%len(third)]:
            third_score += 1
    max_val = max(first_score, second_score, third_score)    
    if max_val == first_score:
        winners.append(1)
    if max_val == second_score:
        winners.append(2)
    if max_val == third_score:
        winners.append(3)
    return winners

def solution2(answers):
    winners = []
    first = [1,2,3,4,5]
    second = [2,1,2,3,2,4,2,5]
    third = [3,3,1,1,2,2,4,4,5,5]
    scores = [0, 0, 0]
    # enumerate() ==> return enumerate object which has index with value of list, set, tuple, dictionary, string
    for i, answer in enumerate(answers):
        if answer == first[i%len(first)]:
            scores[0] += 1
        if answer == second[i%len(second)]:
            scores[1] += 1
        if answer == third[i%len(third)]:
            scores[2] += 1
    for i, score in enumerate(scores):
        max_val = max(scores)
        if score == max_val:
            winners.append(i+1)
    return winners
    
# if __name__ == "__main__":
#     print(solution([1,3,2,4,2]))