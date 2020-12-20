# https://www.acmicpc.net/problem/3182
import sys

graph = []
visited = []

def dfs(start, num_of_contacts):
    global graph
    global visited
    visited[start] = True
    num_of_contacts += 1
    if graph[graph[start]] != 0 and not visited[graph[start]]:
        num_of_contacts = dfs(graph[start], num_of_contacts)
    return num_of_contacts

def init_visited():
    global visited
    for i in range(len(visited)):
        visited[i] = False

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    graph = [0 for i in range(n+1)]
    visited = [False for i in range(n+1)]
    for i in range(n):
        response = int(sys.stdin.readline())
        graph[i+1] = response

    max_contacts = 0
    who_to_contact = n+1
    for i in range(n+1):
        num_of_contacts = 0
        if graph[i] != 0:
            visited[i] = True
            num_of_contacts = dfs(graph[i], num_of_contacts)
        if num_of_contacts > max_contacts:
            max_contacts = num_of_contacts
            who_to_contact = i
        elif num_of_contacts == max_contacts:
            if who_to_contact > i:
                who_to_contact = i
        init_visited()
    print(who_to_contact)