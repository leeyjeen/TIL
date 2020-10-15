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
To be continued..



## Reference
* https://www.tutorialspoint.com/java/java_basic_syntax.htm