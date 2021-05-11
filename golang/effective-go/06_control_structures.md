# 제어 구조(Control structures)

Go언어의 제어 구조는 C언어와 연관성이 있으나 중요한 점에서 차이가 있다. Go언어에서는 `do`와 `while` 반복문이 없으며, 조금 일반화된 `for` 반복문만이 존재한다. `switch`는 더 유연하다. `if`와 `switch`는 `for`와 같이 선택적으로 초기화 구문을 받을 수 있다. `break`와 `continue`구문은 빠져 나가거나 계속할지 식별하기 위한 선택적인 레이블을 갖는다. 그리고 새로운 제어 구조가 존재한다. `type switch`와, 다방향 통신 멀티플렉서로서 사용되는 `select`이다. 문법 또한 다소 다르다. 괄호가 필요하지 않으며 body는 항상 중괄호로 구분되어야 한다.

## If

Go언어에서 간단한 if문은 다음과 같다.
```go
if x > 0 {
    return y
}
```

의무적인 중괄호는 여러 라인에 간단한 if문을 작성할 수 있도록 해준다. 어차피 그렇게 하는 것이 좋은 스타일이며, body가 return 또는 break와 같은 제어 구조를 포함할 때 특히 더 그러하다.

`if`와 `switch`는 초기화 구문을 허용하므로, 지역 변수를 설정하기 위해 초기화 구문을 사용하는 것을 흔히 볼 수 있다.
```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

Go 라이브러리에서, `if`문이 다음 구문으로 실행되지 않을 경우, 즉 body가 `break`, `continue`, `goto`, `return` 으로 끝나는 경우, 불필요한 `else`는 생략되는 것을 볼 수 있을 것이다.
```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

다음은 코드가 일련의 에러 조건들에 대비해 조심해야 하는 흔한 상황들에 대한 예시이다. 제어의 흐름이 성공적이라면 에러가 발생하는 케이스를 제거하면서 코드가 잘 작동할 것이다. 에러 케이스는 return 구문에서 종료하려 하기 때문에 결과적으로 코드는 else 구문을 필요로 하지 않는다.
```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

## 재선언과 재할당(Redelaration and reassignment)

이전 섹션에서 마지막 예시는 짧은 선언 `:=`가 어떻게 동작하는지 살펴 보았다. `os.Open`를 호출하는 선언부를 보자.
```go
f, err := os.Open(name)
```
`f`, `err` 두 변수를 선언하고 있다. 
`f.Stat`을 호출하는 부분을 보자.
```go
d, err := f.Stat()
```
`d`, `err`를 선언하는 것처럼 보인다. 그러나, `err`는 두 구문 모두에서 발견할 수 있다. 이 중복은 합리적이다. `err`가 첫 번째 구문에서 선언되지만, 두 번째 구문에서는 **재할당**된다. 즉, `f.Stat` 호출시 이미 위에서 선언되어 존재하는 `err` 변수를 사용하며, 새로운 값을 할당하는 것을 의미한다.

`:=` 선언에서 다음과 같은 경우 이미 선언된 경우에도 변수 v가 나타날 수 있다.
- v의 선언과 동일한 스코프 안에서 선언할 때 (v가 바깥 스코프에서 이미 선언되었다면, 선언은 새로운 변수를 생성한다.)
- 초기화에 일치하는 값이 v에 할당가능할 때
- 선언에 의해 생성된 다른 변수가 적어도 하나 이상 존재할 때

이 특이한 속성은 완전히 실용적이며, 예를 들어 긴 if-else 체인에서 단일 에러 값을 사용하기 쉽게 해준다. 

여기서 주목할 점은, Go언어에서 function 파라미터와 return 값의 범위가 function body를 감싸고 있는 중괄호 외부에 위치하지만, 그 스코프는 function body의 스코프와 동일하다는 점이다!


## For

Go에서 for 반복문은 C의 for 반복문과 비슷하지만 동일하지는 않다. for와 while을 통합하였으며, do-while이 존재하지 않는다. 세 가지 형태가 있으며, 한 가지 형태만 세미콜론을 사용한다.

```go
// C의 for문과 비슷한 경우
for init; condition; post { }

// C의 while문과 비슷한 경우
for condition { }

// C의 for(;;)와 비슷한 경우
for { }
```

짧은 선언은 반복문에서 index 변수를 선언하기 쉽게 해준다.
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

array, slice, map을 반복 또는 channel로부터 읽어오는 경우, `range` 구문을 활용할 수 있다.
```go
for key, value := range oldMap {
    newMap[key] = vale
}
```

range에서 첫 번째 항목만 필요한 경우 (key 또는 index) 두 번째는 생략할 수 있다.
```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

range에서 두 번째 항목만 필요한 경우 (value) 첫 번째 항목을 버리기 위해 blank 식별자, 즉 언더바를 사용할 수 있다.
```go
sum := 0
for _, value := range array {
    sum += value
}
```

blank 식별자는 나중에 나올 섹션에서 더 자세히 나올 예정이다.

문자열의 경우, range는 UTF-8을 파싱에 의한 개별적인 유니코드 코드를 처리하는 데 더 많은 역할을 할 수 있다. 잘못된 인코딩은 1바이트를 소비하고 U+FFFD rune 문자로 대체할 것이다. (rune이라는 이름은 단일 유니코드 코드 포인트를 위한 Go 용어이다.)

다음 반복문은
```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
다음을 출력한다.
```bash
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

마지막으로, Go는 콤마 연산자가 없으며, ++와 --는 표현식이 아니다. 그러므로 for문에서 여러 변수를 실행하고 싶다면, 병렬 할당을 사용해야 한다. (++, --를 배제하더라도..)
```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

## Switch

Go의 switch는 C보다 더 일반적이다. 표현식이 상수이거나 정수일 필요가 없다. case는 true값이 아닌 경우 match될 때까지 위에서 아래로 내려가면서 비교한다. 따라서 switch로 if-else-if-else 체인을 작성하는 것이 가능하며, 더 Go 스럽다.

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

자동으로 다음으로 통과하는 동작이 없지만, 콤마로 구분된 목록을 사용해 case를 표현할 수 있다.
```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

다른 C와 같은 언어에서처럼 Go에서 일반적이지는 않지만, break 구문으로 switch를 일찍 종료할 수 있다. 그러나, switch 외에 반복문을 빠져나오기 위해 break를 사용해야 하는 경우가 있다. Go에서는 반복문에 레이블을 붙이고, 해당 레이블을 탈출하도록 할 수 있다.
다음 예시는 두 가지를 모두 사용한 예이다.
```go
Loop:
    for n:=0; n<len(src); n+=size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])
        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift) 
        }
    }
```

물론 continue 구문도 선택적 레이블을 사용할 수 있으나, 반복문에만 적용된다.

다음은 두 개의 switch 구문을 사용하는 byte slices 비교 루틴이다.
```go
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i:=0; i<len(a) && i<len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```

## Type switch

switch는 interface 변수의 동적 타입을 발견하기 위해 사용되기도 한다. type switch는 괄호 안의 type 키워드와 함께 type 단언의 문법을 사용한다. switch가 표현식에서 변수를 선언하는 경우, 변수는 각 구문에서 상응하는 타입을 가질 것이다. 각 case에서 같은 이름이지만 다른 타입인 새 변수를 선언하면서 case에서 이름을 재사용하는 것이 관례적이다. 

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```