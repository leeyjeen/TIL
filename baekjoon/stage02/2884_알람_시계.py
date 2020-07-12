import sys

h, m = sys.stdin.readline().split()
h, m = int(h), int(m)

if m - 45 < 0:
    m = m + 15
    if h - 1 < 0:
        h = h + 23
    else:
        h = h - 1
else:
    m = m - 45

print(h, m)