# https://programmers.co.kr/learn/courses/30/lessons/12921

def solution(n):
    primes = [True] * (n+1)
    primes[0], primes[1] = False, False
    for i, v in enumerate(primes):
        if i > 1:
            for j in range(i*2, len(primes), i):
                primes[j] = False
    return primes.count(True)

# if __name__ == "__main__":
#     print(solution(10))
#     print(solution(5))