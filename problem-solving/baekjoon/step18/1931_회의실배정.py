import sys

def greedy_algorithm(times, n):
    meeting_count = 0
    times.sort(key=lambda x: (x[1], x[0]))
    previous_time = []
    for i in range(0, n):
        if i == 0:
            meeting_count += 1
            previous_time = times[i]
            continue
        if times[i][0] >= previous_time[1]:
            meeting_count += 1
            previous_time = times[i]
    return meeting_count

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    times = []
    for i in range(0, n):
        start, end = map(int, sys.stdin.readline().split())
        times.append([start,end])
        
    max_count = greedy_algorithm(times, n)
    print(max_count)