# https://www.acmicpc.net/problem/15828
import sys
import queue

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    q = queue.Queue()
    while True:
        input = int(sys.stdin.readline())
        if input == -1:
            break
        elif input == 0:
            q.get()
        else:
            if q.qsize() < n:
                q.put(input)
    if q.qsize() == 0:
        print("empty")
    else:
        while not q.empty():
            print("{} ".format(q.get()), end='')