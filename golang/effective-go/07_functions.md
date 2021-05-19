# 함수(Functions)

## 다중 반환 값(Multiple return values)

Go언어의 특징 중 하나는 function과 method가 여러 값을 반환할 수 있다는 점이다. 이는 C 프로그램에서 여러 골치아팠던 문제를 향상시킬 수 있도록 해준다. 예를 들어, EOF를 표현하기 위해 -1와 같은 에러를 반환하고, 주소로 전달한 매개변수를 수정할 수 있는 등.

C언어에서는 휘발성의 위치에 비밀스럽게 음수와 에러 코드로 쓰기 에러가 발생한다. Go언어에서 Write는 count와 error를 반환할 수 있다. ex: “Yes, you wrote some bytes but not all of them because you filled the device”. `os` package의 파일에 있는 Write 메서드의 시그니처는 다음과 같다.

```go
func (file *File) Write(b []byte) (n int, err error)
```

문서에서는 쓰인 바이트 개수와 n != len(b)인 경우 non-nil 에러를 반환한다고 설명하고 있다. 이는 일반적인 스타일이다. 에러 처리 섹션에서 더 많은 예시를 보자.

유사한 접근법으로 참조 파라미터를 가장하여 반환값에 포인터를 넘겨줄 필요가 없다. 다음은 숫자와 다음 위치를 반환하는 byte slice의 위치로부터 숫자를 가져오는 간단한 function이다.

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```

다음과 같이 입력 slice b에서 숫자를 스캔할 때 사용할 수 있다.

```go
for i := 0; i < len(b); {
    x, i = nextInt(b, i)
    fmt.Println(x)
}
```

## Named result parameters

Go언어의 함수는 반환 또는 결과 파라미터에 이름을 부여하고 인풋 파라미터처럼 일반 변수로 사용할 수 있다. 이름을 부여하면, 함수가 시작될 때 해당 형식에 해당하는 zero값으로 초기화된다. 인자 없이 return 구문을 수행할 때는 결과 파라미터의 현재 값이 반환 값으로 사용된다.

이름을 부여하는 것이 필수적이지는 않지만 코드를 짧고 간결하게 만들어주며, 문서화가 된다. `nextInt`의 결과에 이름을 부여한다면, 반환되는 int값이 어떤 int값인지 명확해진다.

```go
func nextInt(b []byte, pos int) (value, nextPos int) {
```

이름을 부여한 결과값은 초기화되고 아무 내용 없이 반환되기 때문에, 명확하고 단순해진다. 이를 잘 사용하는 버전의 `io.ReadFull`을 살펴 보자.

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

## Defer

Go언어의 `defer` 구문은 defer를 실행하는 함수가 반환되기 직전에 defer된 함수를 실행하도록 예약한다. 이는 색다르지만 함수가 어떤 경로를 통하여 반환을 하는지 상관없이 해지되어야 하는 리소스와 같은 상황을 해결해야 할 때 효과적인 방법이다. 대표적인 예시로 mutex의 잠금을 풀거나 파일을 닫는 것이 있다.

```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

`Close`와 같은 함수 호출을 defer 하면 두 가지 장점을 얻는다. 첫 번째로는, 파일을 닫는 것을 절대로 잊지 않도록 보장해주는 것이다. 이는 나중에 새로운 반환 경로를 추가해야 하는 경우 쉽게 만들 수 있는 실수이다. 두번째로, open 근처에 close가 있으면 함수의 끝 부분에 위치하는 것보다 명확한 코드가 된다.

defer되는 함수의 매개변수는 (함수가 메서드인 경우 리시버도 포함) 호출이 실행될 때가 아니라 defer가 실행될 때 평가된다. 또한 함수가 실행될 때 값을 변경하는 변수에 대한 걱정할 필요가 없다. 하나의 defer된 호출 위치가 여러 함수의 실행을 defer할 수 있다. 다음은 어리석은 예시이다.

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```

defer된 함수는 LIFO 순서로 실행된다. 그러므로 이 코드는 반환될 때 `4 3 2 1 0`을 출력할 것이다. 더 그럴듯한 예시로서 프로그램을 통해 함수의 실행을 추적하기 위한 간단한 방법이 있다. 아래와 같이 여러 간단한 추적 루틴을 작성할 수 있다.

```go
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
```

defer가 실행될 때 defer된 함수의 매개변수가 평가된다는 사실을 이용하면 더 잘 활용할 수 있다. tracing 루틴은 unstracing 루틴에 매개변수를 설정할 수 있다. 다음 예시를 보자.

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

이는 아래를 출력한다.

```bash
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

다른 언어에서의 블록 레벨 리소스 관리에 익숙한 프로그래머에게는, defer가 이상하게 보일 수 있지만, 가장 흥미롭고 강력한 애플리케이션은 분명 블록 기반이 아니라 함수 기반이라는 사실로부터 온다. panic과 recover 섹션에서 이러한 가능성에 대하여 다른 예시들을 살펴 볼 것이다.
