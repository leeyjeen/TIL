# https://www.acmicpc.net/problem/16165
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    group_infos = {}
    member_infos = {}
    for i in range(n):
        group_name = sys.stdin.readline().rstrip()
        member_count = int(sys.stdin.readline())
        members = []
        for j in range(member_count):
            member_name = sys.stdin.readline().rstrip()
            members.append(member_name)
            member_infos[member_name] = group_name
        group_infos[group_name] = (group_name, members)
    for i in range(m):
        quiz = sys.stdin.readline().rstrip()
        quiz_type = int(sys.stdin.readline())
        if quiz_type == 0:
            for j in sorted(group_infos[quiz][1]):
                print(j)
        elif quiz_type == 1:
            print(member_infos[quiz])
