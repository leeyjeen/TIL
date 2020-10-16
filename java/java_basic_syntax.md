# Java - Basic Syntax
Java 프로그램은 서로의 메소드를 호출하여 통신하는 객체의 모음으로 정의해볼 수 있다. 

* Object
    * 객체는 상태와 동작을 갖는다. 예를 들어 개의 경우, 색깔, 이름, 품종 등의 상태와 꼬리 흔들기, 짖기, 먹기 등의 동작을 갖는다. 객체는 클래스의 인스턴스이다.
* Class
    * 클래스는 해당 유형의 객체가 지원하는 동작/상태를 설명하는 템플릿, 청사진으로 정의할 수 있다.
* Methods
    * 메서드는 기본적으로 동작이다. 클래스에는 여러 메서드가 포함될 수 있다. 메서드에서 로직이 작성되고 데이터가 조작되고 모든 작업이 실행된다.
* Instance Variables
    * 각 객체는 고유한 인스턴스 변수 집합을 갖는다. 객체의 상태는 이러한 인스턴스 변수에 할당된 값에 의해 생성된다.

## First Java Program
'Hello World'를 출력하는 간단한 코드를 작성해보자.
```java
public class FirstJavaProgram {
    public static void main(String []args) {
        System.out.println("Hello World");
    }
}
```

* 파일을 저장, 컴파일, 프로그램 실행하는 방법:
    * 위 코드를 작성 후 `FirstJavaProgram.java` 파일로 저장한다.
    * 커맨드 입력 `javac FirstJavaProgram.java` => 코드를 컴파일한다. 
    * 커맨드 입력 `java FirstJavaProgram` => 프로그램 실행
    * 'Hello World'가 출력된다.

```bash
$ javac FirstJavaProgram.java
$ java FirstJavaProgram
Hello World
```

## Basic Syntax
* Case Sensitivity
    * Java는 대소문자를 구분한다. Java에서 Hello, hello는 다른 의미를 갖는다.
* Class Names
    * 모든 클래스는 첫 번째 글자가 대문자여야 한다. 클래스명에 여러 단어가 사용된다면 각 단어의 첫 글자는 대문자여야 한다. 
        ```java
        class FirstJavaClass
        ```
* Method Names
    * 모든 메서드는 소문자 글자로 시작한다. 메서드명에 여러 단어가 사용된다면 내부 각 단어의 첫 글자는 대문자여야 한다.
        ```java
        public void myMethodName()
        ```
* Program File Name
    * 프로그램 파일명은 클래스명과 정확히 일치해야 한다.
    * 파일 저장시, 클래스명을 사용하여 저장하고 이름 끝에 '.java'를 추가해야 한다. 파일 이름과 클래스 이름이 일치하지 않으면 프로그램이 컴파일되지 않는다.
    * 그러나 파일에 public class가 없는 경우 파일 이름이 클래스 이름과 다를 수 있다. 또한 파일에 public class가 반드시 있어야 하는 것은 아니다.
* public static void main(String args[])
    * Java 프로그램 처리는 모든 Java 프로그램의 필수 부분인 main() 메서드로부터 시작된다.

## Java Identifiers
Java의 모든 컴포넌트는 이름을 필요로 한다. 클래스, 변수, 메서드에 사용되는 이름을 `식별자(identifiers)`라고 한다.

식별자에 대한 규칙은 다음과 같다.

* 모든 식별자는 대문자, 소문자, 달러 기호($), 밑줄 문자(_)로 시작해야 한다.

* 첫 글자 뒤에는 문자들의 조합을 가질 수 있다.

* 예약어는 식별자로 사용할 수 없다.

* 대소문자를 구분한다.

* 가능한 식별자 예시: age, $salary, _value, __1_value.

* 불가능한 식별자 예시: 123abc, -salary.

## Java Variables
Java 변수 타입의 종류:

* Local Variables
* Class Variables (Static Variables)
* Instance Variables (Non-static Variables)

## Java Arrays
배열은 동일한 타입의 여러 변수를 저장하는 객체이다. 그러나 배열 자체는 힙의 객체이다.

## Java Enums
Enums는 Java 5.0에서 도입되었다. Enums는 미리 정의된 몇 가지 값 중 하나만 갖도록 변수를 제한한다. 이 열거된 목록의 값을 enums라고 한다.

enums을 사용하면 코드의 버그 수를 줄일 수 있다.

예를 들어, 신선한 주스 가게에 대한 응용 프로그램을 고려할 때, 컵의 크기를 소형, 중형, 및 대형으로 제한할 수 있다. 이렇게 하면 누구도 소형, 중형, 또는 대형 이외의 크기를 주문할 수 없다.

* 예시:
    ```java
    class FreshJuice {
        enum FreshJuiceSize{ SMALL, MEDIUM, LARGE }
        FreshJuiceSize size;
    }

    public class FreshJuiceTest {
        public static void main(String args[]) {
            FreshJuice juice = new FreshJuice();
            juice.size = FreshJuice.FreshJuiceSize.MEDIUM;
            System.out.println("Size: " + juice.size);
        }
    }
    ```
* Output:
    ```bash
    Size: MEDIUM
    ```

## Java Keywords
다음 목록은 Java의 예약어이다. 이러한 예약어는 상수, 변수 또는 기타 식별자 이름으로 사용할 수 없다.
* abstract
* assert
* boolean
* break
* byte
* case
* catch
* char
* class
* const
* continue
* default
* do
* double
* else
* enum
* extends
* final
* finally
* float
* for
* goto
* if
* implements
* import
* instanceof
* int
* interface
* long
* native
* new
* package
* private
* protected
* public
* return
* short
* static
* strictfp
* super
* switch
* synchronized
* this
* throw
* throws
* transient
* try
* void
* volatile
* while

## Comments in Java
Java는 C, C++와 매우 유사한 단일 및 다중 라인 주석을 지원한다. 주석 내부의 모든 문자는 Java 컴파일러가 무시한다.

    ```java
    public class FirstJavaProgram {
        /* 
        multi line 
        comment
        */
        public static void main(String []args) {
            // single line comment
            /* single line comment */
            System.out.println("Hello World");
        }
    }
    ```

## Inheritance
Java에서 클래스는 클래스로부터 파생될 수 있다. 기본적으로, 새 클래스를 만들어야 하고 필요로 하는 코드를 이미 가지고 있는 클래스가 있는 경우, 이미 존재하는 코드에서 새 클래스를 파생시킬 수 있다.

이 개념을 사용하면 새 클래스에서 코드를 다시 작성하지 않고도 기존 클래스의 필드와 메서드를 다시 사용할 수 있다. 이 시나리오에서 기존 클래스를 *superclass*라고 하고 파생된 클래스를 *subclass*라고 한다.

## Interfaces
Java에서 인터페이스는 객체들이 서로 통신하는 방법에 대한 객체 간의 약속이라고 할 수 있다. 인터페이스는 상속 개념에 관련하여 중요한 역할을 한다.

인터페이스는 파생 클래스(subclass)가 사용할 메서드를 정의한다. 그러나 메서드의 구현은 전적으로 하위 클래스에 달려 있다.

## Reference
* https://www.tutorialspoint.com/java/java_basic_syntax.htm