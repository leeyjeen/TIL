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
+ | 연산자의 양쪽 값을 더한다 | A+B = 30
- | 왼쪽 피연산자에서 오른쪽 피연산자를 뺀다 | A-B = -10
* | 연산자의 양쪽 값을 곱한다 | A*B = 200
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
> | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 큰 경우 조건이 참이 된다 | (A > B) is not true
< | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 작은 경우 조건이 참이 된다 | (A < B) is true
>= | 왼쪽 피연산자의 값이 오른쪽 피연산자의 값보다 크거나 같은 경우 조건이 참이 된다 | (A >= B) is not true
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
`|` | 바이너리 OR 연산자는 피연산자 중 하나라도 존재하는 경우 비트를 복사한다 | `(A | B) = 61` (0011 1101)
^ | 바이너리 XOR 연산자는 하나의 피연산자에 설정된 경우 비트를 복사한다 | (A ^ B) = 49 (0011 0001)
~ | 단항적이며 비트 플립핑 효과 | (~A) = -61 (1100 0011)
<< | 바이너리 왼쪽 시프트 연산자. 왼쪽 피연산자 값을 오른쪽 피연산자에 지정된 비트수만큼 왼쪽으로 이동시킨다 | (A << 2) = 240 (1111 0000)
>> | 바이너리 오른쪽 시프트 연산자. 왼쪽 피연산자 값을 오른쪽 피연산자에 지정된 비트 수만큼 오른쪽으로 이동시킨다 | (A >> 2) = 15 (1111)
>>> | Shift right zero fill operator. 왼쪽 피연산자 값을 오른쪽 피연산자에 지정된 비트 수만큼 오른쪽으로 이동시키고 이동된 값은 0으로 채운다 | (A >>> 2) = 15 (0000 1111)

## The Logical Operators
To be continued..

## The Assignment Operators


## Miscellaneous Operators


## Precedence of Java Operators




## Reference
* https://www.tutorialspoint.com/java/java_basic_operators.htm