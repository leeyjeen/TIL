# Java - Object and Classes
Java는 객체 지향 언어이다. 객체 지향 특성을 가진 언어로서 Java는 다음과 같은 기본 개념을 지원한다.

* Polymorphism
* Inheritance
* Encapsulation
* Abstraction
* Classes
* Objects
* Instance
* Method
* Message Passing

이 장에서는 클래스와 객체라는 개념을 살펴볼 것이다.

* Object
    * 객체는 상태와 동작을 갖는다. 예를 들어 개의 경우, 색깔, 이름, 품종 등의 상태와 꼬리 흔들기, 짖기, 먹기 등의 동작을 갖는다. 객체는 클래스의 인스턴스이다.
* Class
    * 클래스는 해당 유형의 객체가 지원하는 동작/상태를 설명하는 템플릿, 청사진으로 정의할 수 있다.

## Objects in Java
이제 객체가 무엇인지 자세히 살펴 보겠다. 현실 세계를 고려해보면 우리 주변에서 많은 물체, 자동차, 개, 인간 등을 찾을 수 있다. 이 모든 물체는 상태와 행동을 가지고 있다.

개의 경우 상태는 이름, 색, 품종이며 행동은 짖고, 꼬리를 흔들고, 달리는 것이다.

소프트웨어 객체를 실제 객체와 비교하면 매우 유사한 특성이 있다.

소프트웨어 객체에는 상테와 동작도 있다. 소프트웨어 객체의 상태는 필드에 저장되고 동작은 메서드를 통해 표시된다. 

따라서 소프트웨어 개발에서 메서드는 객체의 내부 상태에서 작동하고 객체 간 통신은 메서드를 통해 수행된다.

## Classes in Java
클래스는 개별 객체가 생성되는 청사진이다.
다음은 클래스의 한 예시이다.

    ```java
    public class Dog {
        String breed;
        int age;
        String color;

        void barking() {

        }

        void hungry() {

        }

        void sleeping() {
            
        }
    }
    ```

클래스는 다음 변수 타입을 포함한다.
* 지역 변수
    * 메서드, 생성자 또는 블록 내에서 정의된 변수를 지역 변수라고 한다. 
    * 메서드 내에서 선언 및 초기회되고 메서드가 완료될 때 소멸된다.
* 인스턴스 변수
    * 클래스 내의 변수이지만 메서드 외부에 있는 변수이다.
    * 클래스가 인스턴스화될 때 초기화된다.
    * 특정 클래스의 모든 메서드, 생성자 또는 블록 내에서 액세스할 수 있다.
* 클래스 변수
    * 클래스 내의 변수이지만 메서드 외부에서 static 키워드를 사용하여 선언된 변수이다.

## Constructors
클래스에 대해 논의할 때 가장 중요한 하위 주제 중 하나는 생성자이다. 모든 클래스는 생성자를 갖는다. 클래스에 대한 생성자를 명시적으로 작성하지 않으면 Java 컴파일러는 해당 클래스에 대한 기본 생성자를 빌드한다.

새 객체가 생성될 때마다 적어도 하나의 생성자가 호출된다. 생성자의 주요 규칙은 클래스와 동일한 이름을 가져야 한다는 것이다. 클래스에는 둘 이상의 생성자가 있을 수 있다.

다음은 생성자의 예시이다.

    ```java
    public class Puppy {
        public Puppy() {

        }

        public Puppy(String name) {
            // This constructor has one parameter, name.
        }
    }
    ```

Java는 클래스의 인스턴스를 하나만 생성할 수 있는 Singleton 클래스 또한 지원한다.

## Creating an Object
기본적으로 객체는 클래스로부터 생성된다. Java에서 **new** 키워드는 새 객체를 생성하기 위해 사용된다.

클래스로부터 객체를 생성할 때 다음 3단계를 따른다.
1. 선언
    * 객체 타입, 변수명과 함께 변수 선언
2. 인스턴스화
    * **new** 키워드는 객체를 생성하는 데 사용된다.
3. 초기화
    * **new** 키워드 뒤에 생성자 호출이 이어진다. 이 호출은 새 객체를 초기화 한다.

다음은 객체 생성의 예시이다.
    
```java
public class Puppy {
    public Puppy(String name) {
        // This constructor has one parameter, name.
        System.out.println("Passed Name is :" + name);
    }

    public static void main(String []args) {
        // Following statement would create an object myPuppy
        Puppy myPuppy = new Puppy("tommy");
    }
}
```
Output
```bash
Passed Name is :tommy
```

## Accessing Instance Variables and Methods
인스턴스 변수와 메서드는 생성된 객체를 통해 접근된다. 인스턴스 변수에 접근하려면 다음과 같은 경로를 거쳐야 한다.
```java
/* create an object */
ObjectReference = new Constructor();

/* call a variable as follows */
ObjectReference.variableName;

/* call a class method as follows */
ObjectReference.MethodName();
```

```java
public class Puppy {
    int puppyAge;

    public Puppy(String name) {
        System.out.println("Name chosen is : " + name);
    }

    public void setAge(int age) {
        puppyAge = age;
    }

    public int getAge() {
        System.out.println("Puppy's age is : " + puppyAge);
        return puppyAge;
    }

    public static void main(String []args) {
        Puppy myPuppy = new Puppy("tommy");
        myPuppy.setAge(2);
        myPuppy.getAge();
        System.out.println("Variable Value : " + myPuppy.puppyAge);
    }
}
```
Output:
```
Name chosen is : tommy
Puppy's age is : 2
Variable Value : 2
```

## Source File Declaration Rules
* 소스파일 당 하나의 public 클래스만 있을 수 있다.
* 소스파일은 여러 개의 non-public 클래스를 가질 수 있다.
* 소스파일 이름은 public 클래스 이름의 끝에 .java가 추가된 것이어야 한다. 예: 클래스 이름이 public class Employee{}인 경우 소스파일은 Employee.java이어야 한다.
* 클래스가 패키지 내부에 정의된 경우 package 구문은 소스파일의 첫 번째 구문이어야 한다.
* import 구문이 있는 경우 package 구문과 클래스 선언 사이에 작성되어야 한다. package 구문이 없는 경우 import 구문은 소스파일의 첫 번째 줄에 있어야 한다.
* import 및 package 구문은 소스파일에 있는 모든 클래스를 암시한다. 다른 import 또는 package 구문을 소스파일의 다른 클래스에 선언할 수 없다.

## Java Package
간단히 말해서 package는 클래스와 인터페이스를 분류하는 방법이다. Java로 애플리케이션 개발시 수백 개의 클래스와 인터페이스가 작성되므로, 이러한 클래스를 분류하는 것은 필수이다.

## Import Statements
To be continued..


## Reference
* https://www.tutorialspoint.com/java/java_object_classes.htm