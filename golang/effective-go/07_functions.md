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

