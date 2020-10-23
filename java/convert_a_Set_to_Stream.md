# Convert a Set to Stream in Java
Set 인터페이스는 Collection 인터페이스를 상속하고 Collection는 컬렉션의 순차적 스트림을 반환하는 stream() 메서드를 갖는다.

다음은 구현을 더 잘 이해하기 위한 몇 가지 예시이다.

예시1: Integer HashSet을 Integer Stream으로 변환.
```java
import java.util.*;
import java.util.stream.Stream;

class IntegerSetToStream {
    public static void main(String[] args) {
        // create an Integer HashSet
        Set<Integer> set = new HashSet<>();

        // add elements in set
        set.add(2);
        set.add(4);
        set.add(6);
        set.add(8);
        set.add(10);
        set.add(12);

        // convert Set to Stream
        Stream<Integer> stream = set.stream();
        
        // display elements of Stream using lambda expression
        stream.forEach(elem->System.out.print(elem+" "));
    }
}
```

예시2: String HashSet 변환.
```java
import java.util.*;
import java.util.stream.Stream;

class StringSetToStream {
    public static void main(String[] args) {
        // create an String HashSet
        Set<String> set = new HashSet<>();

        // add elements in set
        set.add("add");
        set.add("elements");
        set.add("in");
        set.add("set");

        // convert Set to Stream
        Stream<String> stream = set.stream();

        // display elements of Stream
        stream.forEach(elem->System.out.print(elem+" "));
    }
}
```
참고: HashSet에 삽입하는 객체는 동일한 순서로 삽입되지 않을 수 있다. 객체는 해시 코드를 기반으로 삽입된다.
 
## Reference
* https://www.geeksforgeeks.org/convert-set-stream-java