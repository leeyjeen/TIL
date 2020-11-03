# https://www.acmicpc.net/problem/1676
# 소인수분해의 성질을 활용하여 N!의 끝에 0이 얼마나 많이 오는지 구하는 문제
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
	# 끝에 0이 오는 개수는 소인수분해 했을 때 2*5 쌍의 개수와 같다
	# 소인수분해 했을 때 2의 개수가 5의 개수보다 훨씬 많으므로, 5의 개수만 카운팅하면 된다
    five_count = 0
    for i in range(1, n+1):
        temp = i
        while temp%5 == 0 and temp > 0:
            five_count += 1
            temp /= 5
    print(five_count)