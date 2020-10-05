# About Golang...
## Go란?
**Go**는 **시스템 프로그래밍**을 염두에 두고 고안된 범용어이다. 2007년에 로버트 그라이저, 롭 파이크, 켄 톰슨에 의해 **구글**에서 개발되었다. **강타입**, **정적 타입**이며 **가비지 컬렉션** 내장 지원 제공 및 **동시 프로그래밍**을 지원한다. 프로그램은 효율적인 의존성 관리를 위해 **패키지**를 사용하여 구성된다. Go 프로그래밍 구현은 실행 가능한 이진 파일을 생성하기 위해 기존의 컴파일 및 링크 모델을 사용한다.

> cf. 강타입:
컴파일러가 소스 코드 변환시 검사 과정에서 형이 다를 경우 형변환을 하지 않고 컴파일을 중단한다. 

> cf. 정적타입:
타입을 컴파일 시에 결정한다.
컴파일 시 타입에 맞지 않은 값이 입력되어 있을 경우 컴파일 에러가 발생한다.

## Go에는 exeception이 없다 !
Go는 다르게 접근한다. 일반적인 에러 핸들링의 경우, Go의 다중값 반환이 반환값에 과부하를 주지 않고 에러를 보고하기 쉽게 해준다. Go 코드는 비정상적인 상태를 표시하기 위해서 **에러 값을 사용**한다.
```go
func Open(name string) (file *File, err error)
```
```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// ...
```

## Go를 사용하는 것의 장점
*Go*는 새롭다. 동시성을 지원한다. 가비지 컬렉션(GC)을 제공한다. 컴파일 속도가 빠르다. 그 외 다음의 이점을 갖는다.:

* 단일 컴퓨터에서 몇 초 안에 대규모 Go 프로그램을 컴파일할 수 있다.

* Go는 종속성 분석을 쉽게 하고 C 스타일의 include 파일과 라이브러리의 오버헤드를 피할 수 있는 소프트웨어 구축 모델을 제공한다.

* Go의 타입 시스템은 계층 구조가 없으므로 타입간 관계를 정의하는 데 시간이 소요되지 않는다. 또한 Go는 정적 타입을 가지고 있으나, 일반적인 객체 지향 언어보다 타입을 가벼워지도록 만든다.

* Go는 완전히 가비지 컬렉트 처리되며 동시 실행 및 통신을 위한 기본적인 지원을 제공한다.

* 이 설계에 의해, Go는 멀티코어 기계에 시스템 소프트웨어를 구축하기 위한 접근법을 제안한다.

## 단일 선언으로 여러 타입의 변수 선언 가능
```go
var a, b, c = 3, 4, "foo"
```

## Goroutine이란?
* Goroutine은 다른 function 또는 method와 동시에 실행되는 function 또는 method이다. 

* Goroutine은 경량 thread라고 생각할 수 있다.

* Goroutine 생성 비용은 thread에 비해 적다.

* Go 애플리케이션은 수천 개의 Goroutine을 동시에 실행하는 것이 일반적이다.

## Go 프로그래밍의 이점
다음은 Go 프로그래밍의 이점이다.:
* 동적 언어와 유사한 패턴을 채택하는 환경 지원 (예: 타입 추론=> `x := 0`은 int 타입의 변수 x의 유효한 선언이다.)

* 컴파일 시간이 빠르다.

* 내장된 동시성 지원: 고루틴을 통한 경량 프로세스, 채널, select문.

* 간결성, 단순성 및 안전성

* 인터페이스 및 타입 임베딩 지원

* Go 컴파일러는 정적 링크를 지원한다. 모든 go 코드는 하나의 큰 바이너리에 정적으로 연결될 수 있으며 종속성에 대한 걱정없이 클라우드 서버에 쉽게 배포할 수 있다.

## 하나의 function에서 여러 변수 반환 가능
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("A", "B")
    fmt.Println(a, b)
}
```

## Go 동시성이란?
일반적으로 대형 프로그램은 여러 개의 작은 하위 프로그램으로 구성된다. 예를 들어, 서버는 웹 브라우저를 통해 생성된 여러 request를 처리하고 response로 HTML 웹 페이지를 제공한다. 이 경우, 각각의 request는 작은 프로그램으로 간주된다.

Golang은 동시성을 통해 이들 각 프로그램의 작은 부품들을 동시에 실행할 수 있게 한다. 고루틴과 채널을 이용한 동시성 지원도 광범위하게 지원하고 있다.

## Interface 구현 방법
Golang interface는 다른 언어에서와 다르다. Golang에서는 type은 메소드 구현을 통하여 interface를 구현한다. 명시적 선언이나 "implements"와 같은 키워드가 존재하지 않는다.

예:
```go
package main

import "fmt"
type I interface {
    M()
}

type T struct {
    S string
}

// M() this method means type T implements the interface I.
// but we don't need to explicitly declare that it does so.
func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"hello"}
    i.M()
}
```

## Golang은 왜 빠른가?
To be continued..


## Reference
* https://www.fullstack.cafe/interview-questions/golang
* https://www.bestinterviewquestion.com/go-interview-questions