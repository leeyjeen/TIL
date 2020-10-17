# Java - Constructors
생성자는 생성될 때 객체를 초기화한다. 클래스와 동일한 이름을 갖고 구문적으로 메서드와 유사하다. 그러나 생성자는 명시적인 반환 타입을 갖지 않는다.

일반적으로 생성자를 사용하여 클래스에 의해 정의된 인스턴스 변수에 초기 값을 제공하거나 완전한 형식의 객체를 생성하기 위해 필요한 그 외 초기 절차를 수행한다.

Java는 모든 멤버 변수를 0으로 초기화하는 기본 생성자를 자동으로 제공하기 때문에 모든 클래스에는 생성자가 있다. 그러나 생성자를 정의하면 기본 생성자가 더 이상 사용되지 않는다.

## Syntax
```java
class ClassName {
    ClassName() {
    }
}
```

Java는 생성자의 두 가지 유형의 생성자를 허용한다.
* 매개변수가 없는 생성자
* 매개변수화된 생성자

## No argument Constructors
매개변수가 없는 생성자를 사용하면 메서드의 인스턴스 변수가 모든 객체에 대한 고정 값으로 초기화된다.

```java
public class MyClass {
    int num;
    MyClass() {
        num = 100;
    }
}
```
```java
public class ConsDemo {
    public static void main(String args[]) {
        MyClass t1 = new MyClass();
        MyClass t2 = new MyClass();
        System.out.println(t1.num + " " + t2.num);
    }
}
```
Output:
```
100 100
```

## Parameterized Constructors
대부분의 경우 하나 이상의 매개변수를 허용하는 생성자가 필요하다. 매개변수는 메서드에 추가되는 것과 동일한 방식으로 생성자에 추가된다. 생성자 이름 뒤의 괄호 안에 선언하기만 하면 된다.

```java
// A simple constructor
class MyClass {
    int x;

    // Following is the constructor
    MyClass(int i) {
        x = i;
    }
}
```
```java
public class ConsDemo {
    public static void main(String args[]) {
        MyClass t1 = new MyClass(10);
        MyClass t2 = new MyClass(20);
        System.out.println(t1.x + " " + t2.x);
    }
}
```
Output:
```
10 20
```

## Reference
* https://www.tutorialspoint.com/java/java_constructors.htm