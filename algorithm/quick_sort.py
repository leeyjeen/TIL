# 피벗 요소에 근거한 배열 분할 함수
def partition(array, low, high):
    # 피벗 요소 선택
    pivot = array[high]
    i = low - 1

    # 왼쪽보다 작은 요소를 피벗의 왼쪽에, 
    # 피벗보다 큰 요소를 피벗의 오른쪽에 배치한다.
    for j in range(low, high):
        if array[j] <= pivot:
            i = i + 1
            array[i], array[j] = array[j], array[i]

    array[i + 1], array[high] = array[high], array[i + 1]

    return i + 1


def quickSort(array, low, high):
    if low < high:
        # 피벗 위치를 선택하고 피벗보다 작은 모든 요소를 왼쪽에, 피벗보다 큰 모든 요소를 오른쪽에 배치한다
        pi = partition(array, low, high)

        # 피벗의 왼쪽 요소들을 정렬한다
        quickSort(array, low, pi - 1)

        # 피벗의 오른쪽 요소들을 정렬한다
        quickSort(array, pi + 1, high)


data = [8, 7, 2, 1, 0, 9, 6]
size = len(data)
quickSort(data, 0, size - 1)
print('Sorted Array in Ascending Order:')
print(data)