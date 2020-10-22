# Stream In Java
Java 8에 도입된 Stream API는 객체 컬렉션을 처리하기 위해 사용된다. stream은 원하는 결과를 생성하기 위해 파이프라인화될 수 있는 다양한 메서드를 지원하는 일련의 객체이다.

Java stream의 특징:
* 스트림은 데이터 구조가 아니라 컬렉션, 배열 또는 I/O 채널로부터 입력을 받는다.
* 스트림은 원래 데이터 구조를 변경하지 않고 파이프라인된 메서드에 따라 결과만 제공한다.
* 각 중간 작업은 느리게 실행되고 그 결과 스트림을 반환하므로 다양한 중간 작업을 파이프라인화 할 수 있다. 터미널 작업은 스트림의 끝을 표시하고 결과를 반환한다.

stream에서의 다른 작업:
## Intermediate Operations
* `map`
    * map 메서드는 주어진 함수를 이 스트림의 요소에 적용한 결과로 구성된 스트림을 반환하는 데 사용된다.
    ```java
    List number = Arrays.asList(2,3,4,5);
    List square = number.stream().map(x->x*x).collect(Collectors.toList());
    ```
* `filter`
    * filter 메서드는 인수로 전달된 조건자에 따라 요소를 선택하는 데 사용된다.
    ```java
    List names = Arrays.asList("Reflection", "Collection", "Stream");
    List result = names.stream().filter(s->s.startWith("S")).collect(Collectors.toList());
    ```
* `sorted`
    * sorted 메서드는 스트림을 정렬하는 데 사용된다.
    ```java
    List names = Arrays.asList("Refelection", "Collection", "Stream");
    List result = names.stream().sorted().collect(Collectors.toList());
    ```

## Terminal Operations
* `collect`
    * collect 메서드는 stream 중간 연산의 결과를 반환하는 데 사용된다.
    ```java
    List number = Arrays.asList(2,3,4,5,3);
    Set square = number.stream().map(x->x*x).collect(Collectors.toSet());
    ```
* `forEach`
    * forEach 메서드는 스트림의 모든 요소를 반복하는 데 사용된다.
    ```java
    List number = Arrays.asList(2,3,4,5);
    number.stream().map(x->x*x).forEach(y->System.out.println(y));
    ```
* `reduce`
    * reduce 메서드는 스트림의 요소들을 단일 값으로 줄이는 데 사용된다. BinaryOperator를 매개변수로 취한다.
    ```java
    List number = Arrays.asList(2,3,4,5);
    int even = number.stream().filter(x->x%2==0).reduce(0,(ans,i)->ans+i);
    ```
    여기서 ans 변수에는 초기값으로 0이 할당되고 i가 더해진다.

```java
import java.util.*;
import java.util.stream.*;

class Demo {
    public static void main(String args[]) {
        // Create a list of integers
        List<Integer> number = Arrays.asList(2,3,4,5);

        // map method
        List<Integer> square = number.stream().map(x->x*x).collect(Collectors.toList());
        System.out.println(square);

        // create a list of String
        List<String> names = Arrays.asList("Reflection", "Collection", "Stream");

        // filter method
        List<String> result = names.stream().filter(s->s.startsWith("S")).collect(Collectors.toList());
        System.out.println(result);

        // sorted method
        List<String> show = names.stream().sorted().collect(Collectors.toList());
        System.out.println(show);

        // create a list of integers
        List<Integer> numbers = Arrays.asList(2,3,4,5,2);

        // collect method returns a set
        Set<Integer> squareSet = numbers.stream().map(x->x*x).collect(Collectors.toSet());
        System.out.println(squareSet);

        // forEach method
        number.stream().map(x->x*x).forEach(y->System.out.println(y));

        // reduce method
        int even = number.stream().filter(x->x%2==0).reduce(0,(ans,i)->ans+i);
        System.out.println(even);
    }
}
```
Output:
```
[4,9,16,25]
[Stream]
[Collection, Reflection, Stream]
[16,4,9,25]
4
9
16
25
6
```

## Important Points/Observations
* 스트림은 소스 다음에 함께 결합된(파이프라인) 0개 이상의 중간 메서드와 묘사된 메서드에 따라 소스에서 얻은 객체를 처리하는 터미널 메서드로 구성된다.
* 스트림은 객체의 원래 값을 변경하지 않고 파이프라인된 메서드에 따라 요소를 계산하는 데 사용된다.

## Reference
* https://www.geeksforgeeks.org/stream-in-java/