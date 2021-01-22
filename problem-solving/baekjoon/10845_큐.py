# https://www.acmicpc.net/problem/10845
import sys

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    queue = []
    for i in range(n):
        inputs = sys.stdin.readline().split()
        command = inputs[0]
        number = 0
        if len(inputs) == 2:
            number = inputs[1]
        
        if command == "push":
            queue.append(number)
        elif command == "pop":
            output = -1
            if len(queue) > 0:
                output = queue[0]
                queue.pop(0)
            print(output)
        elif command == "size":
            print(len(queue))
        elif command == "empty":
            if len(queue) == 0:
                print(1)
            else:
                print(0)
        elif command == "front":
            output = -1
            if len(queue) > 0:
                output = queue[0]
            print(output)
        elif command == "back":
            output = -1
            if len(queue) > 0:
                output = queue[len(queue)-1]
            print(output)