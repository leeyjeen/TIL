# https://programmers.co.kr/learn/courses/30/lessons/42578
def solution(clothes):
    answer = 1
    clothes_dict = {}
    for v in clothes:
        if clothes_dict.get(v[1]):
            clothes_dict[v[1]] += 1
        else:
            clothes_dict[v[1]] = 1
    for v in clothes_dict:
        answer *= (clothes_dict[v]+1)
    answer -= 1
    return answer

# if __name__ == "__main__":
#     print(solution([["yellow_hat", "headgear"], ["blue_sunglasses", "eyewear"], ["green_turban", "headgear"]]))