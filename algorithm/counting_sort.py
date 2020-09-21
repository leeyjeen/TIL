def countingSort(array):
    size = len(array)
    output = [0]*size

    # 카운트 배열 초기화
    count = [0]*10

    # 카운트 배열에 각 요소의 카운트 저장
    for i in range(0, size):
        count[array[i]] += 1

    # 누적 카운트 저장
    for i in range(1, 10):
        count[i] += count[i-1]

    # 카운트 배열에서 원래 배열의 각 요소의 인덱스를 찾는다
    # output 배열에 각 요소를 배치한다
    i = size - 1
    while i >= 0:
        output[count[array[i]] - 1] = array[i]
        count[array[i]] -= 1
        i -= 1

    # 원본 배열에 정렬된 요소를 복사한다
    for i in range(0, size):
        array[i] = output[i]

data = [4, 2, 2, 8, 3, 3, 1]
countingSort(data)
print("Sorted Array in Ascending Order: ")
print(data)