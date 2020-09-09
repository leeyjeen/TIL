# https://programmers.co.kr/learn/courses/30/lessons/42577
def solution(phone_book):
    phone_book.sort()
    for i in range(0, len(phone_book)-1):
        if phone_book[i] in phone_book[i+1]:
            if phone_book[i+1][0:len(phone_book[i])] == phone_book[i]:
                return False
    return True

# if __name__ == "__main__":
#     print(solution(["119", "97674223", "1195524421"]))
#     print(solution(["123", "456", "789"]))
#     print(solution(["12", "123", "1235", "567", "88"]))