# Goroutines
## What are Goroutines?
`고루틴(Goroutine)`은 **다른 함수 또는 메서드와 동시에 실행되는 함수 또는 메서드**이다. 고루틴은 **경량 스레드**로 생각할 수 있다. 고루틴을 만드는 비용은 스레드와 비교할 때 적다. 따라서 Go 애플리케이션이 수천 개의 고루틴을 동시에 실행하는 것은 일반적이다.

## Advantages of Goroutines over threads
* 고루틴은 스레드와 비교할 때 매우 **비용이 적다**. **스택 크기가 몇 kb에 불과**하며 또한 애플리케이션의 요구에 따라 스택이 **늘어나거나 줄어들 수 있는** 반면 스레드의 경우 스택 크기를 지정하고 고정해야 한다.
* 고루틴은 **더 적은 수의 OS threads로** 다중화된다. 수천 개의 고루틴이 있는 프로그램에는 *단 하나의 thread만 있을 수 있다*. 해당 스레드 블록의 고루틴이 사용자 입력을 기다리고 있다고 하면 다른 OS 스레드가 생성되고 나머지 고루틴은 새 OS 스레드로 이동된다. 이 모든 것은 런타임에 의해 처리되며 프로그래머로서 우리는 이러한 복잡한 세부 사항에서 추상화되고 동시성 작업을 위한 깨끗한 API가 제공된다.
* 고루틴은 채널을 사용하여 통신한다. 채널은 고루틴을 사용하여 공유 메모리에 액세스할 때 **경쟁 조건이 발생하지 않도록** 설계되었다. 채널은 고루틴이 통신하는 파이프로 생각할 수 있다.

## How to start a Goroutine?
함수 또는 메서드 호출 앞에 `go` 키워드를 붙이면 동시에 새로운 고루틴이 실행된다.

고루틴을 생성해보자!

```go
package main

import (
    "fmt"
)

func hello() {
    fmt.Println("Hello world goroutine")
}

func main() {
    go hello()
    fmt.Println("main function")
}
```

`main()` 함수의 `go hello()`에서 새로운 고루틴을 시작한다. 이제 `hello()` 함수는 `main()` 함수와 함께 동시에 실행된다. main 함수는 자체 고루틴을 실행하며 이를 main 고루틴이라고 한다.

이 프로그램을 실행하면 놀라운 결과를 확인할 것이다!

이 프로그램은 오직 `main function`이라는 텍스트만 출력한다. 우리가 시작한 고루틴에 어떤 일이 발생한 것일까? 왜 이런 일이 발생하는지 이해하기 위해 고루틴의 두 주요 속성에 대해 이해하여야 한다.

* **새로운 고루틴이 시작됐을 때, 고루틴 호출이 즉시 반환된다. 함수와 달리 제어는 고루틴이 실행을 완료할 때까지 기다리지 않는다. 제어는 고루틴 호출 후 즉시 다음 코드 라인으로 돌아가고 고루틴의 모든 반환값은 무시된다.**
* **다른 고루틴을 실행하려면 main 고루틴이 실행 중이어야 한다. main 고루틴이 종료되면 프로그램이 종료되고 다른 고루틴은 실행되지 않는다.**

이제 고루틴이 실행되지 않은 이유를 이해할 수 있을 것이다. `go hello()` 호출 후, 제어는 hello 고루틴이 끝나기를 기다리지 않고 코드의 다음 라인으로 즉시 반환되고 `main function`을 출력했다. 그런 다음 실행할 다른 코드가 없기 때문에 main 고루틴이 종료되어 hello 고루틴을 실행할 기회를 얻지 못했다.

이제 이를 고쳐보자.

```go
package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello world goroutine")
}

func main() {
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}
```

`time.Sleep()` 호출을 통해 실행되고 있는 고루틴을 sleep하였다. 이 경우 main 고루틴은 1초 동안 휴면 상태가 된다. 이제 `go hello()` 호출은 main 고루틴이 종료되기 전에 실행하기에 충분한 시간이 있다. 이 프로그램은 먼저 `Hello world goroutine`을 출력하고 1초 동안 기다린 다음 `main function`을 출력한다.

다른 고루틴이 실행을 마칠 때까지 기다리기 위해 main 고루틴에서 sleep을 사용하는 방법은 고루틴의 작동 방식을 이해하기 위해 사용하는 방법이다. 다른 모든 고루틴이 실행을 마칠 때까지 main 고루틴을 블록하기 위해 우리는 채널을 사용할 수 있다.

## Starting multiple Goroutines
고루틴을 더 잘 이해하기 위하여 여러 고루틴을 시작하는 프로그램을 하나 더 작성해보자!

```go
package main

import (
    "fmt"
    "time"
)

func numbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d", i)
    }
}

func alphabets() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c", i)
    }
}

func main() {
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```
Output:
```
1 a 2 3 b 4 c 5 d e main terminated
```

위의 프로그램은 `main()`에서 두 개의 고루틴을 시작한다. 이 두 고루틴은 동시에 실행된다. numbers 고루틴은 250 밀리초 주기로 숫자를 출력한다. 마찬가지로 alphabets 고루틴은 400 밀리초 주기로 알파벳을 출력한다. main 고루틴은 두 개의 고루틴을 시작하고 3000 밀리초 동안 휴면 후 종료된다.

## Reference
* https://golangbot.com/goroutines/