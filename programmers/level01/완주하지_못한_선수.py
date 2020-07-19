import collections

# https://programmers.co.kr/learn/courses/30/lessons/42576
 
def solution(participant, completion):
    # collections.Counter ==> dict subclass for counting hashable objects
    answer = collections.Counter(participant) - collections.Counter(completion)
    return list(answer)[0]

def solution2(participant, completion):
    participant.sort()
    completion.sort()
    # zip(*iterable) ==> list multiple slicing
    for p, c in zip(participant, completion):
        print(p, c)
        if p != c:
            return p
    
    return participant[-1]

# if __name__ == "__main__":
#     print(solution2(["a", "b", "c", "c"], ["a", "c", "c"]))