from heapq import heappop, heappush, heapify

# 빈 heap 생성
heap = []
heapify(heap)

# heappush 함수 이용하여 항목 추가
heappush(heap, -1*10)
heappush(heap, -1*30)
heappush(heap, -1*20)
heappush(heap, -1*400)

# 최대 요소 값 출력
print("Head value of heap : " + str(-1*heap[0]))

# 힙의 요소 출력
print("The heap elements : ")
for i in heap:
    print(-1*i, end=' ')
print("\n")

element = heappop(heap)

# 힙의 요소 출력
print("The heap elements : ")
for i in heap:
    print(-1*i, end=' ')
