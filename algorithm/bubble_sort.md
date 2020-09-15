# Bubble Sort | 버블 정렬
`버블 정렬`은 인접한 두 개의 요소를 비교하여 순서가 잘못된 경우 반복적으로 교환하는 방식의 가장 간단한 정렬 알고리즘이다.

## 예시
1. 첫 번째 과정

    (`5 1` 4 2 8) -> (`1 5` 4 2 8) : 첫 번째 두 요소 비교하여 5 > 1 이므로 교환.

    (1 `5 4` 2 8) -> (1 `4 5` 2 8) : 5 > 4 이므로 교환.

    (1 4 `5 2` 8) -> (1 4 `2 5` 8) : 5 > 2 이므로 교환

    (1 4 2 `5 8`) -> (1 4 2 `5 8`) : 이미 정렬되어 있으므로 교환하지 않음.

2. 두 번째 과정

    (`1 4` 2 5 8) -> (`1 4` 2 5 8)

    (1 `4 2` 5 8) -> (1 `2 4` 5 8) : 4 > 2 이므로 교환.

    (1 2 `4 5` 8) -> (1 2 `4 5` 8)

    (1 2 4 `5 8`) -> (1 2 4 `5 8`)

    이제 배열이 이미 정렬되어 있지만 알고리즘이 완료되었는지 알지 못한다. 알고리즘은 정렬이 되었는지 알기 위해 교환이 일어나지 않더라도 전체 과정이 필요하다.

3. 세 번째 과정

    ...(반복)

## 코드
```python
def bubbleSort(arr):
    n = len(arr)
    for i in range(0, n-1):
        for j in range(0, n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]

arr = [64, 34, 25, 12, 22, 11, 90]
bubbleSort(arr)

print("Sorted array is:")
for i in range(len(arr)):
    print("%d" % arr[i], end=" ")
```
Output:
```bash
Sorted array is:
11 12 22 25 34 64 90
```
위의 코드는 항상 `O(n^2)`의 시간 복잡도를 갖는다. 
내부 반복문에서 아무 교환이 일어나지 않을 경우 알고리즘을 멈추게 하여 최적화할 수 있다.

```python
def bubbleSort(arr):
    n = len(arr)
    for i in range(0, n-1):
        swapped = False
        for j in range(0, n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True

        if swapped == False:
            break
arr = [64, 34, 25, 12, 22, 11, 90]
bubbleSort(arr)

print("Sorted array is:")
for i in range(len(arr)):
    print("%d" % arr[i], end=" ")
```
Output:
```bash
Sorted array is:
11 12 22 25 34 64 90
```

## Reference
- https://www.geeksforgeeks.org/bubble-sort/