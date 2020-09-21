# Counting Sort | 계수 정렬
`계수 정렬`은 배열의 각 고유 요소의 발생 횟수를 계산하여 배열의 요소를 정렬하는 정렬 알고리즘이다. 카운트는 보조 배열에 저장되며, 카운트를 보조 배열의 인덱스로 매핑하여 정렬이 이루어진다.

# 계수 정렬의 작동 방식
1. 주어진 배열에서 최대값 요소(`max`로 할당)를 찾는다.

2. 길이가 `max+1`인 모든 요소가 0인 배열을 초기화한다. 이 배열은 배열의 요소의 카운트를 저장하기 위해 사용된다.

3. 카운트 배열의 각 인덱스에 각 요소의 카운트를 저장한다. 예를 들어, 요소 "3"의 카운트가 2인 경우, 2는 카운트 배열의 3번째 위치에 저장된다. 요소 "5"가 배열에 존재하지 않는 경우 0이 5번째 위치에 저장된다.

4. 카운트 배열의 요소의 **누적합**을 저장한다. 이는 정렬된 배열의 올바른 인덱스로 요소를 배치하는 데 도움이 된다.

5. 카운트 배열에서 원래 배열의 각 요소의 인덱스를 찾는다. 이것은 누적 카운트를 제공한다.

6. 각 요소를 올바른 위치에 배치한 후, 카운트를 1씩 감소시킨다.

## 코드
```python
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
```
Output:
```
Sorted Array in Ascending Order: 
[1, 2, 2, 3, 3, 4, 8]
```

## 복잡도
### 시간 복잡도
* 최악의 경우: `O(n+k)`
* 최선의 경우: `O(n+k)`
* 평균의 경우: `O(n+k)`

### 공간 복잡도
* `O(max)`

## Reference
- https://www.programiz.com/dsa/counting-sort
- 