# https://programmers.co.kr/learn/courses/30/lessons/42888
def solution(record):
    history = []
    users = {}
    for r in record:
        sentence = r.split()
        cmd = sentence[0]
        uid = sentence[1]
        if cmd == "Enter":
            name = sentence[2]
            users[uid] = name
            history.append((uid, "님이 들어왔습니다."))
        elif cmd == "Leave":
            history.append((uid, "님이 나갔습니다."))
        elif cmd == "Change":
            users[uid] = sentence[2]
    answer = []
    for h in history:
        answer.append(users[h[0]]+h[1])
    return answer

# if __name__ == "__main__":
#     print(solution(["Enter uid1234 Muzi", "Enter uid4567 Prodo","Leave uid1234","Enter uid1234 Prodo","Change uid4567 Ryan"]))