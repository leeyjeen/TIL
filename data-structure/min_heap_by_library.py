from heapq import heapify, heappush, heappop

# 빈 heap 생성
heap = []
heapify(heap)

# heappush 함수 이용하여 항목 추가
heappush(heap, 10)
heappush(heap, 30)
heappush(heap, 20)
heappush(heap, 400)

# 최소 요소 값 출력
print("Head value of heap : " + str(heap[0]))

# 힙의 요소 출력
print("The heap elements : ")
for i in heap:
    print(i, end=' ')
print("\n")

element = heappop(heap)

# 힙의 요소 출력
print("The heap elements : ")
for i in heap:
    print(i, end=' ')