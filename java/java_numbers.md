# Java - Numbers Class
일반적으로 `Numbers`로 작업할 때 `byte`, `int`, `long`, `double` 등과 같은 원시 데이터 유형을 사용한다.

예시:
```java
int i = 5000;
float gpa = 13.65f;
double mask = 125;
```

그러나 개발 과정에서 원시 데이터 유형 대신 객체를 사용해야 하는 상황이 발생한다. 이를 위해 Java는 wrapper class를 제공한다.

모든 wrapper class(Integer, Long, Byte, Double, Float, Short)는 추상 클래스 Number의 하위 클래스이다.

wrapper class의 객체는 해당 원시 데이터 유형을 포함하거나 래핑한다. 원시 데이터 유형을 객체로 변환하는 것을 `boxing`이라고 하며 컴파일러에서 처리한다. 따라서 wrapper class를 사용하는 동안 원시 데이터 유형의 값을 wrapper class의 생성자에 전달하기만 하면 된다.

그리고 Wrapper 객체는 다시 원시 데이터 유형으로 변환되며 이 프로세스를 `unboxing`이라고 한다. Number 클래스는 `java.lang` 패키지의 한 부분이다.

unboxing 예시:
```java
public class Test {
    public static void main(String args[]) {
        Integer x = 5; // boxes int to an Integer object
        x = x + 10; // unboxes the Integer to a int
        System.out.println(x);
    }
}
```
Output:
```
15
```

x에 정수 값이 할당되면 x가 정수 객체이므로 컴파일러는 정수를 boxing한다. 나중에 x는 정수로 더해질 수 있도록 unboxing된다.

## Number Methods
Number 클래스의 모든 하위 클래스가 구현하는 인스턴스 메서드 리스트:
### `xxxValue()` method
메서드를 호출하는 Number 객체의 값을 원시 데이터 형식으로 변환하고 반환한다.

```java
byte byteValue()
short shortValue()
int intValue()
long longValue()
float floatValue()
double doubleValue()
```

```java
public class Test {
    public static void main(String args[]) {
        Integer x = 5;

        // Returns byte primitive data type
        System.out.println(x.byteValue());

        // Returns double primitive data type
        System.out.println(x.doubleValue());

        // Returns long primitive data type
        System.out.println(x.longValue());
    }
}
```
Output:
```
5
5.0
5
```

### `compareTo()` method
메서드를 호출한 Number 객체를 인수와 비교한다. Byte, Long, Integer 등을 비교할 수 있다.

그러나 다른 유형의 경우는 비교할 수 없다. 메서드를 호출하는 Number 객체와 인수는 모두 같은 유형이어야 한다.

```java
public int compareTo(NumberSubClass referenceName)
```
#### Parameters
* referenceName
    * Byte, Double, Integer, Float, Long, Short
#### Return Value
* Integer가 인수와 동일하다면 0을 반환한다.
* Integer가 인수보다 작다면 -1을 반환한다.
* Integer가 인수보다 크다면 1을 반환한다.

```java
public class Test {
    public static void main(String args[]) {
        Integer x = 5;

        System.out.println(x.compareTo(3));
        System.out.println(x.compareTo(5));
        System.out.println(x.compareTo(8));
    }
}
```
Output:
```
1
0
-1
```

### `equals()` method
메서드를 호출하는 Number 객체가 인수로 전달된 객체와 같은지 여부를 검사한다.

```java
public boolean equals(Object o)
```

#### Return Value
인수가 null이 아니고 같은 유형의 객체이며 값은 숫자 값을 갖는 경우 True를 반환한다.

```java
public class Test {
    public static void main(String args[]) {
        Integer x = 5;
        Integer y = 10;
        Integer z = 5;
        Short a = 5;

        System.out.println(x.equals(y));
        System.out.println(x.equals(z));
        System.out.println(x.equals(a));
    }
}
```
Output:
```
false
true
false
```

### `valueOf()` method
전달된 인수의 값을 포함하는 관련 Number 객체를 반환한다. 인수는 원시 데이터 유형, 문자열 등이 될 수 있다.

이 메서드는 정적 메서드이다. 두 개의 인수를 사용할 수 있다. 하나는 문자열이고 다른 하나는 기수이다.

```java
static Integer valueOf(int i)
static Integer valueOf(String s)
static Integer valueOf(String s, int radix)
```

#### Parameters
* **i** - Integer 표현이 반환될 int 값
* **s** - Integer 표현이 반환될 String 값
* **radix** - 전달된 문자열을 기반으로 반환된 정수의 값을 결정하기 위해 사용된다.

```java
public class Test {
    public static void main(String args[]) {
        Integer x = Integer.valueOf(9);
        Double c = Double.valueOf(5);
        Float a = Float.valueOf("80");
        Integer b = Integer.valueOf("444", 16);

        System.out.println(x); 
        System.out.println(c);
        System.out.println(a);
        System.out.println(b);
    }
}
```
Output:
```
9
5.0
80.0
1092
```

### `toString()` method
Number 객체의 값을 나타내는 String 객체를 가져오기 위해 사용된다.

메서드가 기본 데이터 형식을 인수로 사용하면 기본 데이터 형식 값을 나타내는 String 객체가 반환된다.

메서드가 두 개의 인수를 사용하는 경우 두 번째 인수에 지정된 기수로 첫 번째 인수의 문자열 표현이 리턴된다.

```java
String toString()
static String toString(int i)
```

```java
public class Test {
    public static void main(String args[]) {
        Integer x = 5;

        System.out.println(x.toString());
        System.out.println(Integer.toString(12));
    }
}
```
Output:
```
5
12
```

### `parseInt()` method
특정 문자열의 **기본 데이터 유형**을 가져오기 위해 사용된다. `parseXXX()`는 정적 메서드이며 하나 또는 두 개의 인수를 가질 수 있다.

```java
static int parseInt(String s)
static int parseInt(String s, int radix)
```

```java
public class Test {
    public static void main(String args[]) {
        int x = Integer.parseInt("9");
        double c = Double.parseDouble("5");
        int b = Integer.parseInt("444", 16);

        System.out.println(x);
        System.out.println(c);
        System.out.println(b);
    }
}
```
Output:
```
9
5.0
1092
```

### `abs()` method
인수의 절대값을 제공한다. 인수는 `int`, `float`, `double`, `short`, `byte`가 될 수 있다.

```java
double abs(double d)
float abs(float f)
int abs(int i)
long abs(long lng)
```

```java
public class Test {
    public static void main(String args[]) {
        Integer a = -8;
        double d = -100;
        float f = -90;

        System.out.println(Math.abs(a));
        System.out.println(Math.abs(d));     
        System.out.println(Math.abs(f)); 
    }
}
```
Output:
```
8
100.0
90.0
```

### `ceil()` method
인수보다 크거나 같은 가장 작은 정수를 제공한다.

```java
double ceil(double d)
double ceil(float f)
```

### `floor()` method
인수보다 작거나 같은 가장 큰 정수를 제공한다.

```java
double floor(double d)
double floor(float f)
```

```java
public class Test {
    public static void main(String args[]) {
        double d = -100.675;
        float f = -90;

        System.out.println(Math.ceil(d));
        System.out.println(Math.ceil(f)); 

        System.out.println(Math.floor(d));
        System.out.println(Math.floor(f)); 
    }
}
```
Output:
```
-100.0
-90.0
-101.0
-90.0
```

### `random()` method
0.0에서 1.0 사이의 난수를 생성하기 위해 사용한다. 범위는 0.0 <= Math.random < 1.0이다. 산술 연산을 이용하여 다른 범위를 구할 수 있다.

```java
static double random()
```

```java
public class Test {
    public static void main(String args[]) {
        System.out.println(Math.random());
        System.out.println(Math.random());
    }
}
```
Output:
```
0.16763945061451657
0.400551253762343
```

## Reference
* https://www.tutorialspoint.com/java/java_numbers.htm