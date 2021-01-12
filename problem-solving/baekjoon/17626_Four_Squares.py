# https://www.acmicpc.net/problem/17626
import sys

def get_square_count(n):
    min_count = 4
    for i in range(int(n**0.5), 0, -1):
        if i*i == n:
            min_count = min(min_count, 1)
            break
        else:
            new_n = n - i*i
            for j in range(int(new_n**0.5), 0, -1):
                if new_n == j*j:
                    min_count = min(min_count, 2)
                    break
                else:
                    new_n2 = n - i*i - j*j
                    for k in range(int(new_n2**0.5), 0, -1):
                        if new_n2 == k*k:
                            min_count = min(min_count, 3)
    return min_count

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    print(get_square_count(n))

# 시간 초과 발생
# def get_square_count(n): 
#     dp = [0 for i in range(n+1)]
#     dp[1] = 1
    
#     for i in range(2, n+1):
#         min_count = n
#         j = 1
#         while j*j <= i:
#             if min_count > dp[i-j*j]:
#                 min_count = dp[i-j*j]
#             j += 1 
#         dp[i] = min_count + 1
    
#     return dp[n]