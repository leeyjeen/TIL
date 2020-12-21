# https://www.acmicpc.net/problem/5635
import sys
import datetime

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    students = []
    for i in range(n):
        name, day, month, year = sys.stdin.readline().split()
        birth = datetime.datetime(int(year), int(month), int(day))
        students.append((birth, name))
    students.sort()
    print(students[-1][1])
    print(students[0][1])