# 초기화(Initialization)

C나 C++의 초기화와 겉으로는 많이 달라 보이지 않지만, Go에서의 초기화는 더 강력하다. 초기화하는 동안 복잡한 구조체가 생성될 수 있고 초기화된 객체 사이에서 순서 이슈가 다른 패키지 사이에서도 정확하게 처리된다.

## 상수(Constants)

Go에서 상수는 상수이다! 함수 내 지역적으로 정의되더라도, 컴파일 시점에 생성된다. 숫자, 문자(runes), 문자열, boolean값만 가능하다. 컴파일 시점 제한때문에, 상수를 정의하는 표현식은 컴파일러가 컴파일 시점에 실행 가능한 상수 표현식이어야 한다. 예를 들어, `1<<3`은 상수 표현식이지만, `math.Sin(math.Pi/4)`은 상수 표현식이 아니다. `math.Sin` 함수 호출이 런타임에 발생하기 때문이다.

Go에서는 열거형 상수를 iota라는 열거자를 이용하여 생성한다. iota는 표현식의 일부로 묵시적으로 반복될 수 있다. 이는 복잡한 값들로 구성된 집합들을 만들기 쉽게 해준다.

```go
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```

Go는 String과 같은 메서드를 유저가 정의한 타입에 붙일 수 있다. 이로서 임의의 값의 포맷을 자동으로 출력을 위해 바뀌게 할 수 있다. 이와 같은 경우는 struct에 일반적으로 적용되지만, 이 기법은 ByteSize와 같은 실수 스칼라 타입에 또한 유용하다.

```go
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}
```

`YB` 표현식은 `1.00YB`를 출력한다. `1e13`은 `9.09TB`를 출력한다.

ByteSize의 String 메서드를 구현하기 위한 여기서의 `Sprintf`의 사용은 안전하다. (무한 재귀를 피했다) 형변환 때문이 아니라 `Sprintf`를 string 타입이 아닌 `%f` 옵션으로 호출했기 때문이다. `Sprintf`는 문자열이 필요할 때만 `String` 메서드를 호출하고, `%f`는 실수값을 사용한다.

## 변수(Variables)

변수는 상수처럼 초기화되지만, 런타임에 계산되는 일반적인 표현식이 가능하다.

```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

## init 함수

마지막으로, 각 소스파일은 어떤 필요한 상태를 셋업하기 위하여 init 함수를 정의할 수 있다. (각 파일은 여러 init 함수를 가질 수 있다.) init 함수는 모든 임포트된 패키지들이 초기화되고 패키지 내의 모든 변수 선언이 평가된 이후에 호출된다.

선언의 형태로 표현할 수 없는 것들을 초기화하는 것 외에도, init 함수는 실제 프로그램의 실행이 일어나기 전에 프로그램의 상태를 검증하고 올바르게 복구하기 위해 자주 사용된다.

```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```