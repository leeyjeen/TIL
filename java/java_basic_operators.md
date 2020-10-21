# Java - Basic Operators
Java는 변수를 조작할 수 있는 풍부한 연산자를 제공한다. 모든 연산자를 다음 그룹으로 분류할 수 있다.
* Arithmetic Operators
* Relational Operators
* Bitwise Operators
* Logical Operators
* Assignment Operators
* Misc Operators

## The Arithmetic Operators
산술 연산자는 대수에서 사용되는 것과 같은 방식으로 수학 표현식에서 사용된다.
예시: A=10, B=20인 경우
Operator | Description | Example
------- | ------ | ------
`+` | 연산자의 양쪽 값을 더한다 | A+B = 30
`-` | 왼쪽 피연산자에서 오른쪽 피연산자를 뺀다 | A-B = -10
`*` | 연산자의 양쪽 값을 곱한다 | A*B = 200
/ | 왼쪽 피연산자를 오른쪽 피연산자로 나눈다 | B/A = 2
% | 왼쪽 피연산자를 오른쪽 피연산자로 나눈 나머지를 반환한다 | B%A = 0
++ | 피연산자의 값을 1 증가시킨다 | B++ = 21
-- | 피연산자의 값을 1 감소시킨다 | B-- = 19 

## The Relational Operators
Java에서 다음과 같은 관계 연산자를 지원한다.

예시: A=10, B=20인 경우
Operator | Description | Example
------- | ------ | ------
== | 두 피연산자의 값이 같은지 확인하고 같은 경우 조건이 참이 된다 | (A == B) is not true
!= | 두 피연산자의 값이 같은지 확인하고 같지 않은 경우 조건이 참이 된다 | (A != B) is true
`>` | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 큰 경우 조건이 참이 된다 | (A > B) is not true
< | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 작은 경우 조건이 참이 된다 | (A < B) is true
`>=` | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 크거나 같은 경우 조건이 참이 된다 | (A >= B) is not true
<= | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 작거나 같은 경우 조건이 참이 된다 | (A <= B) is true

## The Bitwise Operators
Java는 정수 유형인 long, int, short, char 및 byte에 적용할 수 있는 여러 비트 연산자를 정의한다.

비트 연산자는 비트에 작동하고 비트 단위 연산을 수행한다. a=60, b=13이라고 가정하면 바이너리 형식으로는 다음과 같다.

a = 0011 1100  
b = 0000 1101

a&b = 0000 1100  
a|b = 0011 1101  
a^b = 0011 0001  
~a = 1100 0011  

예시: A=60, B=13인 경우
Operator | Description | Example
------- | ------ | ------
& | 바이너리 AND 연산자는 두 피연산자가 모두 존재하는 경우 결과에 비트를 복사한다 | (A & B) = 12 (0000 1100)
```|``` | 바이너리 OR 연산자는 피연산자 중 하나라도 존재하는 경우 비트를 복사한다 | `(A | B) = 61` (0011 1101)
^ | 바이너리 XOR 연산자는 하나의 피연산자에 설정된 경우 비트를 복사한다 | (A ^ B) = 49 (0011 0001)
~ | 단항적이며 비트 플립핑 효과 | (~A) = -61 (1100 0011)
<< | 바이너리 왼쪽 시프트 연산자. 왼쪽 피연산자 값을 오른쪽 피연산자에 지정된 비트수만큼 왼쪽으로 이동시킨다 | (A << 2) = 240 (1111 0000)
`>>` | 바이너리 오른쪽 시프트 연산자. 왼쪽 피연산자 값을 오른쪽 피연산자에 지정된 비트 수만큼 오른쪽으로 이동시킨다 | (A >> 2) = 15 (1111)
`>>>` | Shift right zero fill operator. 왼쪽 피연산자 값을 오른쪽 피연산자에 지정된 비트 수만큼 오른쪽으로 이동시키고 이동된 값은 0으로 채운다 | (A >>> 2) = 15 (0000 1111)

## The Logical Operators
예시: A=true, B=false인 경우
Operator | Description | Example
------- | ------ | ------
&& | 논리 AND 연산자. 두 피연산자가 모두 0이 아니면 조건이 참이 된다 | (A && B) = false
`||` | 논리 OR 연산자. 두 피연산자 중 하나아 0이 아니면 조건이 참이 된다 | `(A || B)` = true
! | 논리 NOT 연산자. 피연산자의 논리 상태를 반전하기 위해 사용한다. | !(A && B) = true

## The Assignment Operators
Operator | Description | Example
------- | ------ | ------
= | 간단한 할당 연산자. 오른쪽 피연산자의 값을 왼쪽 피연산자에 할당한다 | (C = A + B)는 (A + B)값을 C에 넣는다
+= | 더하고 할당하는 연산자. 왼쪽 피연산자에 오른쪽 피연산자를 더한 결과를 왼쪽 피연산자에 할당한다 | (C += A)는 (C = C + A)와 같다
-= | 빼고 할당하는 연산자. 왼쪽 피연산자에서 오른쪽 피연산자를 뺀 결과를 왼쪽 피연산자에 할당한다
*= | 곱하고 할당하는 연산자. 오른쪽 피연산자와 왼쪽 피연산자를 곱한 결과를 왼쪽 피연산자에 할당한다
/= | 나누고 할당하는 연산자. 왼쪽 피연산자를 오른쪽 피연산자로 나눈 결과를 왼쪽 피연산자에 할당한다
%= | 나머지를 구하고 할당하는 연산자. 두 개의 피연산자를 사용하여 나머지를 구한 결과를 왼쪽 피연산자에 할당한다
`<<=` | 왼쪽 시프트 AND 할당 연산자. | `(C <<= 2)`는 `C=C<<2`와 같다
`>>=` | 오른쪽 시프트 AND 할당 연산자. | `(C >>= 2)`는 `C=C>>2`와 같다
&= | bitwise AND 대입 연산자. | (C &= 2)는 C=C&2와 같다
^= | bitwise exclusive OR 대입 연산자. | (C ^= 2)는 C=C^2와 같다
`|=` | bitwise inclusive OR 대입 연산자. | `(C |= 2)`는 `C = C|2`와 같다

## Miscellaneous Operators
### Conditional Operator ( ?: )
조건 연산자는 삼항 연산자라고도 한다. 이 연산자는 세 개의 피연산자로 구성되며 boolean 표현식을 평가하는 데 사용된다. 연산자의 목표는 변수에 할당할 값을 결정하는 것이다. 연산자는 다음과 같이 작성된다.
```
variable x = (expression) ? value if true : value if false
```
```java
public class Test {
    public static void main(String args[]) {
        int a, b;
        a = 10;
        b = (a == 1) ? 20 : 30;
        System.out.println("Value of b is : " + b);
        
        b = (a == 10) ? 20 : 30;
        System.out.println("Value of b is : " + b);
    }
}
```
Output:
```
Value of b is : 30
Value of b is : 20
```

### instanceof Operator
이 연산자는 객체 참조 변수에만 사용된다. 연산자는 객체가 특정 유형(클래스 또는 인터페이스 유형)인지 확인한다. instanceof 연산자는 다음과 같이 작성된다.

```
( Object reference variable ) instanceof  (class/interface type)
```

연산자 왼쪽의 변수가 참조하는 객체가 오른쪽의 클래스/인터페이스 유형에 대한 IS-A 검사를 통과하면 결과가 참이 된다. 

```java
public class Test {
    public static void main(String args[]) {
        String name = "James";
        // following will return true since name is type of String
        boolean result = name instanceof String;
        System.out.println(result);
    }
}
```
Output:
```
true
```

```java
class Vehicle {}

public class Car extends Vehicle {
    public static void main(String args[]) {
        Vehicle a = new Car();
        boolean result = a instanceof Car;
        System.out.println(result);
    }
}
```
Output:
```
true
```

## Precedence of Java Operators
연산자 우선 순위는 표현식에서 용어 그룹화를 결정한다. 이는 식이 평가되는 방식에 영향을 준다. 특정 연산자는 다른 연산자보다 우선 순위가 높다. 예를 들어 곱셈 연산자는 더하기 연산자보다 우선 순위가 높다.

아래 표는 우선 순위가 가장 높은 연산자부터 가장 낮은 연산자까지 표시한다. 표현식 내에서 우선 순위가 높은 연산자가 먼저 계산된다.

Category | Operator | Associativity
------- | ------ | ------
Postfix | expression++ expression-- | Left to right
Unary | ++expression --expression +expression -expression ~ ! | Right to left
Multiplicative | * / % | Left to right
Additive | + - | Left to right
Shift | `<<` `>>` `>>>` | Left to right
Relational | `<` `>` `<=` `>=` instanceof | Left to right
Equality | == != | Left to right
Bitwise AND | & | Left to right
Bitwise XOR | ^ | Left to right
Bitwise OR | `|` | Left to right
Logical AND | && | Left to right
Logical OR | `||` | Left to right
Conditional | ?: | Right to left
Assignment | = += -= *= /= %/ ^= `|=` `<<=` `>>=` `>>>=` | Right to left

## Reference
* https://www.tutorialspoint.com/java/java_basic_operators.htm