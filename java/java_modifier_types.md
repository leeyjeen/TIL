# Java - Modifier Types
제어자(Modifier)는 정의의 의미를 변경하기 위해 추가하는 예약어이다. Java 언어에는 다음을 포함한 다양한 제어자가 있다.
* 접근 제어자(Access Modifiers)
* 비접근 제어자(Non Access Modifiers)

제어자를 사용하려면, 클래스, 메서드 또는 변수 정의시 해당 키워드를 포함한다. 제어자는 다음 예에서와 같이 선언부 나머지 부분 앞에 온다. 

```java
public class className {
    // ...
}

private boolean myFlag;
static final double weeks = 9.5;
protected static final int BOXWIDTH = 42;

public static void main(String[] args) {
}
```

## Access Control Modifiers
Java는 클래스, 변수, 메서드 및 생성자에 대한 접근 레벨을 설정하기 위한 여러 접근 제어자를 제공한다.
* 4가지 접근 레벨
    * `default`: 같은 package 내부에서만 접근 가능. 기본값. 제어자가 필요하지 않다.
    * `private`: 같은 class 내부에서만 접근 가능.
    * `public`: 모두 접근 가능.
    * `protected`: 같은 package 내부와, 상속 관계에 있는 모든 서브 클래스에서 접근 가능.

## Non-Access Modifiers
Java는 다른 많은 기능을 수행하기 위해 여러 비접근 제어자를 제공한다.
* `static`
* `final`
* `abstract`
* `synchronized`, `volatile`

### The Static Modifier
#### Static Variables
`static` 키워드는 클래스에 대해 생성된 인스턴스와 독립적으로 존재하는 변수를 생성할 때 사용된다. 클래스의 인스턴스 수에 관계없이 정적 변수의 복사본이 하나만 존재한다.

정적 변수는 클래스 변수라고도 한다. 지역 변수는 정적으로 선언할 수 없다.

#### Static Methods
`static` 키워드는 클래스에 대해 생성된 인스턴스와 독립적으로 존재하는 메서드를 생성하기 위해 사용된다.

정적 메서드는 정의된 클래스 객체의 인스턴스 변수를 사용하지 않는다. 정적 메서드는 매개변수에서 모든 데이터를 가져와 변수에 대한 참조 없이 해당 매개변수로부터 무언가를 계산한다.

{클래스이름}.{변수} 또는 {클래스이름}.{메서드이름}을 사용하여 클래스 변수 및 메서드에 접근할 수 있다.

예시:
```java
public class InstanceCounter {
    private static int numInstances = 0;

    protected static int getCount() {
        return numInstances;
    }

    private static void addInstance() {
        numInstances++;
    }

    InstanceCounter() {
        InstanceCounter.addInstance();
    }

    public static void main(String[] args) {
        System.out.println("Starting with " + InstanceCounter.getCount() + " instances");

        for (int i = 0; i < 500; ++i) {
            new InstanceCounter();
        }
        System.out.println("Created " + InstanceCounter.getCount() + " instances");
    }
}
```
Output:
```
Starting with 0 instances
Created 500 instances
```

### The Final Modifier
#### Final Variables
`final` 변수는 명시적으로 한 번만 초기화할 수 있다. final로 선언된 참조 변수는 다른 객체를 참조하도록 다시 할당할 수 없다.

그러나 객체 내의 데이터는 변경할 수 있다. 따라서 객체의 상태는 변경할 수 있지만 참조는 변경할 수 없다.

변수와 함께 final 제어자는 종종 static과 함께 사용되어 상수를 클래스 변수로 만든다.

예시:
```java
public class Test {
    final int value = 10;

    // The following are examples of declaring constants:
    public static final int BOXWIDTH = 6;
    static final String TITLE = "Manager";

    public void changeValue() {
        value = 12; // will give an error
    }
}
```

#### Final Methods
`final` 메서드는 하위 클래스에 의해 재정의될 수 없다. final 제어자는 메서드가 하위 클래스에서 수정되는 것을 방지한다.

메서드를 final로 만드는 주된 의도는 메서드의 내용을 외부인이 변경해서는 안된다는 것이다.

```java
public class Test {
    public final void changeName() {
        // ...
    }   
}
```

#### Final Classes
클래스를 final로 선언하는 주된 목적은 클래스가 서브 클래스가 되는 것을 방지하는 것이다. 클래스가 final로 선언되면 어떤 클래스도 final 클래스의 기능을 상속할 수 없다.

```java
public final class Test {
    // ...
}
```

### The abstract Modifier
#### Abstract Class
추상 클래스는 인스턴스화할 수 없다. 클래스가 추상으로 선언된 경우 유일한 목적은 클래스를 확장하는 것이다. 

클래스는 추상적이면서 동시에 최종적일 수는 없다(최종 클래스는 확장될 수 없기 때문에). 클래스에 추상 메서드가 포함된 경우 클래스는 추상으로 선언되어야 한다. 그렇지 않으면 컴파일 오류가 발생한다.

추상 클래스에는 추상 메서드와 일반 메서드가 모두 포함될 수 있다.

```java
abstract class Caravan {
    private double price;
    private String model;
    private String year;
    public abstract void goFast(); // an abstract method
    public abstract void changeColor();
}
```

#### Abstract Methods
추상(abstract) 메서드는 구현없이 선언된 메서드이다. 메서드 본문(구현)은 서브클래스에서 제공한다. 추상 메서드는 final이거나 strict할 수 없다.

추상 클래스를 확장하는 모든 클래스는 이 또한 추상 클래스가 아닌 경우 상위 클래스의 모든 추상 메서드를 구현해야 한다.

클래스에 하나 이상의 추상 메서드가 포함된 경우 클래스는 추상으로 선언되어야 한다. 추상 클래스가 추상 메서드를 꼭 포함할 필요는 없다.

추상 메서드는 세미콜론으로 끝난다. 
```java
public abstract sample();
```

```java
public abstract class SuperClass {
    abstract void m();  // abstract method
}

class SubClass extends SuperClass {
    // implements the abstract method
    void m() {
        // ...
    }
}
```

### The Synchronized Modifier
To be continued..
### The Transient Modifier

### The Volatile Modifier

## Reference
* https://www.tutorialspoint.com/java/java_modifier_types.htm