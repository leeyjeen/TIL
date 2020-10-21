# Java - Loop Control
## while Loop in Java
```java
while (Boolean_expression) {
    // statements
}
```

## for loop in Java
```java
for (initialization; Boolean_expression; update) {
    // statements
}
```

```java
public class Test {
    public static void main(String args[]) {
        for (int x = 10; x < 20; x = x + 1) {
            System.out.println("value of x : " + x);
        }
    }
}
```
Output:
```
value of x : 10
value of x : 11
value of x : 12
value of x : 13
value of x : 14
value of x : 15
value of x : 16
value of x : 17
value of x : 18
value of x : 19
```

## do...while loop in Java
```java
do {
    // statements
} while (Boolean_expression);
```

```java
public class Test {
    public static void main(String args[]) {
        int x = 10;

        do {
            System.out.println("value of x : " + x);
        } while (x < 20);
    }
}
```
Output:
```
value of x : 10
value of x : 11
value of x : 12
value of x : 13
value of x : 14
value of x : 15
value of x : 16
value of x : 17
value of x : 18
value of x : 19
```

## Loop Control Statements
### break statement
* 반복문 내에서 `break`문이 발견되면 반복문이 즉시 종료되고 프로그램 제어는 반복문 다음의 구문에서 다시 시작된다.
* `switch`문에서 case를 종료하기 위해 사용할 수 있다.

```java
break;
```

```java
public class Test {
    public static void main(String args[]) {
        int [] numbers = {10, 20, 30, 40, 50};

        for (int x : numbers) {
            if (x == 30) {
                break;
            }
            System.out.println(x);
        }
    }
}
```
Output:
```
10
20
```

### continue statement
`continue` 키워드는 모든 반복문 제어 구조에서 사용할 수 있다. 반복문이 즉시 다음 반복으로 넘어가도록 한다.

* for 반복문에서 `continue` 키워드를 사용하면 제어가 즉시 update문으로 이동한다.
* while 반복문 또는 do/while 반복문에서 제어는 즉시 Boolean 표현식으로 이동한다.

```java
continue;
```

```java
public class Test {
    public static void main(String args[]) {
        int [] numbers = {10, 20, 30, 40, 50};

        for (int x : numbers) {
            if (x == 30) {
                continue;
            }

            System.out.println(x);
        }
    }
}
```
Output:
```
10
20
40
50
```

## Enhanced for loop in Java
Java 5부터 향상된 for 반복문이 도입되었다. 주로 배열을 포함한 요소 컬렉션을 탐색하기 위해 사용된다.

```java
for (declaration : expression) {
    // statements
}
```
* Declaration
    * 새로 선언된 블록 변수는 접근하는 배열의 요소와 호환되는 유형이다. 변수는 for 블록 내에서 사용할 수 있으며 해당 값은 현재 배열 요소와 동일하다.
* Expression
    * 반복해야 하는 배열. 표현식은 배열 변수 또는 배열을 반환하는 메서드 호출일 수 있다.

```java
public class Test {
    public static void main(String args[]) {
        int [] numbers = {10, 20, 30, 40, 50};

        for (int x : numbers) {
            System.out.print(x + ",");
        }
        System.out.println();

        String [] names = {"James", "Larry", "Tom"};

        for (String name : names) {
            System.out.print(name + ",");
        }
    }
}
```
Output:
```
10, 20, 30, 40, 50,
James, Larry, Tom,
```

## Reference
* https://www.tutorialspoint.com/java/java_loop_control.htm