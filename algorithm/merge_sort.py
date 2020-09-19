def merge(arr, left, mid, right):
    fIdx = left
    rIdx = mid+1
    sIdx = left
    sortArr = [0]*(right+1)

    while fIdx <= mid and rIdx <= right:
        if arr[fIdx] <= arr[rIdx]:
            sortArr[sIdx] = arr[fIdx]
            fIdx += 1
        else:
            sortArr[sIdx] = arr[rIdx]
            rIdx += 1
        sIdx += 1

    if fIdx > mid:
        for i in range(rIdx, right+1):
            sortArr[sIdx] = arr[i]
            sIdx += 1
    else:
        for i in range(fIdx, mid+1):
            sortArr[sIdx] = arr[i]
            sIdx += 1
    
    for i in range(left, right+1):
        arr[i] = sortArr[i]

def mergeSort(arr, left, right):
    if left < right:
        mid = (left+right)//2

        mergeSort(arr, left, mid)
        mergeSort(arr, mid+1, right)

        merge(arr, left, mid, right)
       

def printList(array):
    for i in range(len(array)):
        print(array[i], end=" ")
    print()

if __name__ == '__main__':
    array = [6, 5, 12, 10, 9, 1]

    mergeSort(array, 0, len(array)-1)

    print("Sorted array is: ")
    printList(array)