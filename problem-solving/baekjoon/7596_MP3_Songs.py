# https://www.acmicpc.net/problem/7596
import sys

if __name__ == "__main__":
    count = 1
    while True:
        n = int(sys.stdin.readline())
        if n == 0:
            break
        playlist = []
        for i in range(n):
            playlist.append(sys.stdin.readline().rstrip())
        playlist.sort()
        print(count)
        for song in playlist:
            print(song)
        count += 1



