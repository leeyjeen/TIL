# Convert Stream to Set in Java
## 방법1: Collectors 사용하기
Stream `collect()` 메서드는 스트림에서 요소를 가져와서 컬렉션에 저장한다. `collect(Collector.toSet())`는 Stream에서 Set으로 요소를 컬렉트한다.

```java
import java.util.*;
import java.util.stream.Stream;
import java.util.stream.Collectors;

class StreamToSet {
    public static void main(String[] args) {
        // create a Stream of Integers
        Stream<Integer> stream = Stream.of(-2, -1, 0, 1, 2);

        // use Stream.collect() to collect the
        // elements of stream in a container
        Set<Integer> streamSet = stream.collect(Collectors.toSet());

        // display the elements
        streamSet.forEach(num->System.out.println(num));
    }
}
```
Output:
```
-1
0
-2
1
2
```

## 방법2: Stream -> Array -> Set
Stream을 Set으로 변환하는 문제는 두 부분으로 나뉘어 질 수 있다.
1) Stream을 Array로 변환한다.
2) Array를 Set으로 변환한다.

```java
import java.util.*;
import java.util.stream.Stream;
import java.util.stream.Collectors;

class StreamToArrayToSet {
    public static void main(String[] args) {
        // create a Stream of Strings
        Stream<String> stream = Stream.of("A", "B", "C", "D");

        // convert Stream into an Array
        String[] arr = stream.toArray(String[] :: new);

        // create a HashSet
        Set<String> set = new HashSet<>();

        // convert Array to set
        Collections.addAll(set, arr);

        // display the elements
        set.forEach(str->System.out.println(str));
    }
}
```
Output:
```
A
B
C
D
```
참고: HashSet은 생성된 해시 값으로 임의의 순서로 입력을 받기 때문에 출력은 랜덤이다.

## 방법3: forEach 사용
`forEach()`를 사용하여 Stream을 Set으로 변환할 수 있다. `forEach()` 메서드를 사용하여 Stream의 모든 요소를 반복한 다음 `set.add()`를 사용하여 각 요소를 빈 Set에 추가한다.

```java
import java.util.*;
import java.util.stream.Stream;
import java.util.stream.Collectors;

class StreamToSet {
    public static void main(String[] args) {
        // create a Stream of Integers
        Stream<Integer> stream = Stream.of(5, 10, 15, 20, 25);

        // create a HashSet
        Set<Integer> set = new HashSet<>();

        // use set.add() to add elements
        // of stream into empty set
        stream.forEach(set::add);

        // display the elements
        set.forEach(num->System.out.println(num));
    }
}
```
Output:
```
20
50
25
10
15
```
참고: Stream이 병렬이면 forEach() 메서드를 사용하여 요소가 원래 순서대로 처리되지 않을 수 있다. 원래 순서를 유지하기 위해 forEachOrdered() 메서드를 사용한다.




## Reference
* https://www.geeksforgeeks.org/convert-stream-set-java/