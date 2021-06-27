# Concurrency

## Share by communicating

동시성 프로그래밍은 큰 주제이며 여기에는 Go만의 특화된 하이라이트를 위한 공간이 있다.

많은 환경에서 동시성 프로그래밍은 공유 변수에 대한 올바른 액세스를 구현하는 데 필요한 세부사항으로 인해 어려워진다. Go는 **공유된 값이 채널에서 전달**되고 실제로 별도의 실행 스레드에 의해 활발히 공유되지 않는 **다른 접근 방식**을 장려한다. 지정된 시간에 하나의 고루틴만 값에 접근할 수 있다. 데이터 경주는 설계상 발생할 수 없다. 이러한 사고 방식을 장려하기 위해 이를 슬로건으로 축소해보았다.

> 메모리를 공유하여 통신하지 마라. 대신, 통신하여 메모리를 공유하라.

이 접근 방식은 도가 지나칠 수 있다. 예를 들어, 참조 횟수는 정수 변수 주위에 뮤텍스를 두는 것이 가장 좋다. 그러나 고급 접근 방식으로서 채널을 사용하여 액세스를 제어하면 명확하고 정확한 프로그램을 쉽게 작성할 수 있다.

이 모델에 대해 생각하는 한 가지 방법은 하나의 CPU에서 실행되는 일반적인 단일 스레드 프로그램을 고려하는 것이다. 이는 동기화 기본 요소가 필요하지 않다. 이제 다른 인스턴스를 실행한다. 이 또한 동기화가 필요 없다. 이제 그 둘이 통신하도록 한다. 통신이 동기화 장치인 경우 다른 동기화가 필요하지 않다. 예를 들어 유닉스 파이프라인은 이 모델에 완벽하게 맞다. 동시성에 대한 Go의 접근 방식은 Hoare의 CSP (Communication Sequential Processes)에서 비롯되었지만 Unix 파이프의 type-safe 일반화로도 볼 수 있다.

## Goroutines

스레드, 코루틴, 프로세스 등의 기존 용어가 부정확한 의미를 전달하기 때문에 goroutine(고루틴)이라고 부른다. 고루틴은 간단한 모델을 갖는다. 동일한 주소 공간에서 다른 고루틴과 동시에 실행되는 함수이다. 이는 경량이며, 스택 공간 할당 비용보다 약간 더 든다. 또한 스택은 소규모로 시작되므로 비용이 적게 들며 필요에 따라 힙 스토리지를 할당(및 해제)하여 확장할 수 있다. 고루틴은 여러 OS 스레드에 멀티플렉싱되므로 I/O를 기다리는 동안처럼 하나가 차단되어야 하는 경우 다른 스레드는 계속 실행된다. 이러한 설계는 스레드 생성 및 관리의 많은 복잡성을 숨긴다.

새로운 고루틴에서 호출을 실행하려면 함수 또는 메서드 호출 앞에 go 키워드를 붙인다. 호출이 완료되면 고루틴이 조용히 종료된다. (이 효과는 백그라운드에서 명령을 실행하기 위한 Unix 셸의 & 표기법과 유사하다.)

```go
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

함수 리터럴은 고루틴 호출에 유용할 수 있다. 

```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

Go에서 함수 리터럴은 클로저이다. 구현은 함수가 참조하는 변수가 활성 상태인 동안 지속되도록 보장한다.

이 예시는 함수가 완료 신호를 보낼 방법이 없기 때문에 실용적이지 않다. 이를 위해, channel(채널)이 필요하다. 

## Channels

맵과 마찬가지로 채널도 `make`로 할당되며, 결과 값은 기본 데이터 구조에 대한 참조 역할을 한다. 옵션 정수 파라미터가 제공된 경우 **채널의 버퍼 크기**를 설정한다. 기본값은, 버퍼링되지 않았거나 동기화된 채널의 경우 0이다.

```go
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

버퍼링되지 않은 채널은 통신(값 교환)과 동기화를 결합하여 두 계산(고루틴)이 알려진 상태에 있음을 보장한다.

채널을 사용하는 좋은 idiom들이 많이 있다. 이제 시작하겠다. 이전 섹션에서 우리는 배경에 어떤 종류의 것을 시작했다. 채널에서는 시작 고루틴이 정렬이 완료될 때까지 대기할 수 있다.

```go
c := make(chan int)  // 채널 할당
// 고루틴에서 정렬 시작. 정렬이 완료되면 채널에 신호를 전송한다.
go func() {
    list.Sort()
    c <- 1  // 신호를 보낸다. 값은 중요하지 않다.
}()
doSomethingForAWhile()
<-c   // 정렬이 완료될 때까지 기다린 후 전송된 값을 버린다.
```

수신할 데이터가 있을 때까지 receiver는 항상 블록된다. 채널이 버퍼링되지 않으면, receiver가 값을 수신할 때까지 sender는 블록된다. 채널에 버퍼가 있는 경우 sender는 값이 버퍼에 복사될 때까지만 블록된다. 버퍼가 가득 찼다면, 이는 일부 receiver가 값을 되찾아올 때까지 기다린다는 의미이다.

버퍼링된 채널을 세마포어처럼 사용하여 처리량을 제한할 수 있다. 다음 예시에서는 수신 요청이 handle로 전달되어 채널로 값을 보내고 요청을 처리한 다음 채널에서 값을 수신하여 다음 소비자를 위한 "세마포어"를 준비한다. 채널 버퍼 용량은 처리할 동시 호출 수를 제한한다.

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // 활성 대기열이 소모될 때까지 기다린다
    process(r)  // 시간이 오래 걸릴 수 있다
    <-sem       // 완료. 다음 요청을 실행할 수 있다
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // handle이 끝날 때까지 기다리지 않는다
    }
}
```

일단 `MaxOutstanding`개의 핸들러가 프로세스를 실행하고 나면, 기존 핸들러 중 하나가 완료되어 버퍼로부터 수신될 때까지 더 이상 채워진 채널 버퍼로 보내려는 시도를 차단한다.

그러나 이 설계에는 문제가 있다. `Serve`는 들어오는 모든 요청에 대해 새로운 루틴을 생성한다. 단, 이러한 요청 중 `MaxOutstanding`만큼만 언제든지 실행할 수 있다. 따라서 요청이 너무 빨리 들어오면 프로그램에서 리소스를 무제한으로 사용할 수 있다. 우리는 `Serve`를 고루틴의 생성 게이트로 바꿈으로써 그 결함을 해결할 수 있다. 다음은 명백한 해결 방법이다. 하지만 나중에 해결할 버그가 있다는 점에 유의하자.

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // 버그가 있다. 아래 설명을 참조하라.
            <-sem
        }()
    }
}
```

버그는 Go의 for 반복문에서 루프 변수가 각 반복에 재사용되므로 `req` 변수가 모든 루틴에서 공유된다는 것이다. 그건 우리가 원하는 것이 아니다. 우리는 그 `req`가 각 고루틴마다 고유한지 확인해야 한다. 이렇게 하는 한 가지 방법으로 고루틴의 closure에 인수로 `req` 값을 전달하는 것이다.

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}
```

이 버전을 이전 버전과 비교하여 closure 선언 및 실행 방식의 차이를 확인하라. 또 다른 솔루션은 이 예시와 동일한 이름의 새 변수를 생성하는 것이다.

```go
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // 고루틴에 대한 새 req 인스턴스를 생성한다
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

쓰는 것이 이상해 보일지도 모른다.

```go
req := req
```

그러나 이는 Go에서 합법적이고 관용적이다. 동일한 이름을 가진 새로운 버전의 변수를 가져와 의도적으로 루프 변수를 로컬로 섀도잉하지만 각 고루틴에 고유하다.

서버 작성의 일반적인 문제로 돌아가서, 리소스를 잘 관리하는 또 다른 접근 방식은 요청 채널에서 모두 읽는 고정된 수의 `handle` 고루틴을 시작하는 것이다. 고루틴의 수는 처리할 동시 호출 수를 제한한다. 이 `Serve` 함수는 종료하라는 채널도 받는다. 고루틴을 시작한 후 해당 채널에서 수신을 차단한다.

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // 핸들러 시작
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // 종료하라는 메시지를 기다린다
}
```

## Channels of channels


## Parallelization


## A leaky buffer