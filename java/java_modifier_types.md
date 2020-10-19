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

### Default Access Modifier - No Keyword
Default 접근 제어자는 클래스, 필드, 메서드 등에 대한 접근 제어자를 명시적으로 선언하지 않음을 의미한다.

접근 제어자 없이 선언된 변수 또는 메서드는 동일한 패키지의 다른 클래스에서 사용할 수 있다. 인터페이스의 필드는 암시적으로 `public static final`이고 인터페이스의 메서드는 기본적으로 `public`이다.

```java
String version = "1.5.1";

boolean processOrder() {
    return true;
}
```

### Private Access Modifier - Private
`private`으로 선언된 메서드, 변수 및 생성자는 선언된 클래스 자체 내에서만 액세스할 수 있다.

`private` 접근 제어자는 가장 제한적인 접근 레벨이다. 클래스와 인터페이스는 `private`일 수 없다.

클래스에 public getter 메서드가 존재하는 경우에는 클래스 외부에서 private으로 선언된 변수에 접근할 수 있다.

private 제어자를 사용하는 것은 객체 자체를 캡슐화하고 외부 세계로부터 데이터를 숨기는 주요 방법이다.

```java
public class Logger {
    private String format;

    public String getFormat() {
        return this.format;
    }

    public void setFormat(String format) {
        this.format = format;
    }
}
```

여기서 Logger 클래스의 format 변수는 private이므로 다른 클래스가 값을 직접 검색하거나 설정할 수 없다.

따라서 이 변수를 외부에서 사용할 수 있도록 하기 위해 format의 값을 반환하는 `getFormat()`과 값을 설정하는 `setFormat(String)` 두 개의 public 메서드를 정의했다.

### Public Access Modifier - Public
`public`으로 선언된 클래스, 메서드, 생성자, 인터페이스 등은 다른 모든 클래스에서 접근할 수 있다.

따라서 public 클래스 내에서 선언된 필드, 메서드, 블록은 Java 유니버스에 속한 모든 클래스에서 접근할 수 있다.

그러나 접근하려는 public 클래스가 다른 패키지에 있는 경우 public 클래스가 import되어야 한다. 클래스 상속으로 인해 클래스의 모든 public 메서드와 변수는 하위 클래스에 의해 상속된다.

```java
public static void main(String[] args) {
    // ...
}
```

애플리케이션의 `main()` 메서드는 `public`이어야 한다. 그렇지 않으면 Java 인터프리터(예: java)가 클래스를 실행하기 위해 호출할 수 없다.

### Protected Access Modifier - Protected
슈퍼 클래스에서 `protected`로 선언된 변수, 메서드 및 생성자는 다른 패키지의 하위 클래스 또는 protected 멤버 클래스의 패키지 내의 클래스에서만 접근할 수 있다.

`protected` 접근 제어자는 클래스와 인터페이스에 적용할 수 없다. 메서드, 필드는 `protected`로 선언될 수 있지만 인터페이스의 메서드 및 필드는 `protected`로 선언될 수 없다.

`protected` 접근은 하위 클래스가 helper 메서드 또는 변수를 사용할 수 있는 기회를 제공하는 동시에 관련 없는 클래스가 이를 사용하지 못하도록 한다.

```java
class AudioPlayer {
    protected boolean openSpeaker(Speaker sp) {
        // implementation details
    }
}

class StreamingAudioPlayer extends AudioPlayer {
    boolean openSpeaker(Speaker sp) {
        // implementation details
    }
}
```

여기서 `openSpeaker()` 메서드를 `private`으로 정의하면 `AudioPlayer` 이외의 다른 클래스에서 접근할 수 없다. `public`으로 정의하면 모든 외부 세계에서 접근할 수 있다. 그러나 우리의 의도는 이 메서드를 하위 클래스에만 노출하는 것이므로 `protected` 제어자를 사용했다.

### Access Control and Inheritance
상속된 메서드에 대해 다음 규칙이 적용된다.
* 상위 클래스에서 `public`으로 선언된 메서드는 모든 하위 클래스에서도 `public`이어야 한다.
* 상위 클래스에서 `protected`로 선언된 메서드는 하위 클래스에서 `protected` 또는 `public`이어야 한다. 그들은 `private`일 수 없다.
* `private`으로 선언된 메서드는 전혀 상속되지 않으므로 규칙이 없다.

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
`synchronized` 키워드는 한 번에 하나의 스레드에서만 메소드에 접근할 수 있음을 나타내는 데 사용된다. synchronized 제어자는 4개의 접근 레벨 제어자와 함께 적용할 수 있다.

```java
public synchronized void showDetails() {
    // ...
}
```

### The Transient Modifier
인스턴스 변수는 이를 포함하는 객체를 직렬화할 때 JVM이 특정 변수를 건너 뛰도록 표시하기 위해 `transient`로 표시된다.

이 제어자는 변수를 생성하는 구문에 포함되며, 클래스 또는 변수의 데이터 타입 앞에 위치한다.

```java
public transient int limit = 55;    // will not persist
public int b;   // will persist
```

### The Volatile Modifier
`volatile` 제어자는 변수에 접근하는 스레드가 항상 메모리의 master 사본과 변수의 private 사본을 병합해야 함을 JVM에 알리기 위해 사용된다.

volatile 변수에 접근하면 주 메모리에 있는 모든 캐시된 변수 사본이 동기화된다. volatile은 객체 또는 private 유형의 인스턴스 변수에만 적용할 수 있다. volatile 객체 참조는 `null`일 수 있다.

```java
public class MyRunnable implements Runnable {
    private volatile boolean active;

    public void run() {
        active = true;
        while (active) { // line 1
            // ...
        }
    }

    public void stop() {
        active = false; // line 2
    }
}
```

일반적으로 `run()`은 하나의 스레드(Runnable 사용을 시작한 스레드)에서 호출되고 `stop()`은 다른 스레드에서 호출된다. line 1에서 캐시된 `active`값이 사용되는 경우, line 2에서 `active`값을 `false`로 설정했을 때 루프가 중지되지 않을 수 있다. 이 때 `volatile`을 사용해야 한다.

## Reference
* https://www.tutorialspoint.com/java/java_modifier_types.htm