# The blank identifier

for range 반복문과 맵의 컨텍스트에서 blank identifier에 대해 여러 번 언급했다. blank identifier는 모든 타입의 값으로 할당 및 선언할 수 있으며, 값은 무시된다. Unix /dev/null 파일에 쓰는 것과 약간 비슷하다. 변수는 피룡하지만 실제 값은 관련이 없는 자리 표시자로 사용할 쓰기 전용 값을 나타낸다. 우리가 이미 본 것 이상의 용도가 있다.

## The blank identifier in multiple assignment

for range 반복문에서의 blank identifier의 사용은 일반적인 상황의 특수한 케이스인 다중 할당이다

어떤 할당이 왼쪽에 여러 값을 필요로 하지만, 프로그램에 의해 그 중 하나의 값이 사용되지 않을 경우, 할당의 왼쪽에 blank identifier를 사용하면 더미 변수를 만들 필요가 없고 값이 버려진다는 것을 명확히 할 수 있다. 예를 들어, 값과 오류를 반환하지만 오류만 중요한 함수를 호출할 때, blank identifier를 사용하여 관련없는 값을 버릴 수 있다.

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
	fmt.Printf("%s does not exist\n", path)
}
```

때로는 오류를 무시하기 위해 오류 값을 버리는 코드를 볼 수 있다. 이는 끔찍한 예시히다. 항상 오류 반환을 확인하여야 한다. 이유가 있어서 제공되는 것이다.

```go
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
```

## Unused imports and variables ¶

사용하지 않는 package를 import하거나 변수를 선언하는 것은 에러이다. 사용되지 않은 imports는 프로그램을 부풀리게 하고 컴파일을 느리게 하며, 초기화되었지만 사용되지 않은 변수는 최소 계산 낭비이며 아마도 더 큰 버그를 생성할 수 있다. 그러나 프로그램이 현재 개발 중일 때, 사용하지 않는 imports와 변수는 종종 발생하는데, 컴파일을 진행하기 위해 삭제하고 나중에 다시 필요로하는 경우에만 갖도록 하는 것은 성가실 수 있다. blank identifier가 해결 방법을 제공한다.

반으로 작성된 아래 프로그램은 사용되지 않은 두 개의 imports (`fmt`, `io`)와 사용되지 않은 변수(`fd`)를 갖고 있으므로, 컴파일되지 않을 것이다. 그러나 지금까지 코드가 올바른지 확인하는 것이 좋다.

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
```

사용하지 않은 imports에 대한 불만을 없애려면 blank identifier를 사용하여 import된 package의 심볼을 가리키도록 하자. 마찬가지로, 사용하지 않은 변수 `fd`를 blank identifier에 할당하면 사용되지 않은 변수 오류가 사라진다. 아래 버전의 프로그램을 컴파일된다.

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

관례상, import 에러를 침묵시키는 글로벌 선언은 import 직후에 나오고 주석을 달아야 한다. 이 둘은 쉽게 찾을 수 있도록 하고 나중에 정리할 것을 상기시켜 준다.

## Import for side effect

이전의 예시에서 `fmt`, `io`와 같이 사용되지 않은 import는 결국 사용하거나 제거해야한다. blank 할당은 코드를 진행중인 작업으로 식별한다. 그러나 때때로 명시적 사용 없이 부작용만을 위해 package를 import하는 것은 유용할 수 있다. 예를 들어, `init` 함수 동안, `net/http/pprof` package는 디버깅 정보를 제공하는 HTTP 핸들러를 등록한다. export된 API를 가지고 있지만, 대부분의 클라이언트는 웹페이지를 통해 핸들러 등록과 데이터 액세스만 필요하다. 부작용에 대해서만 package를 import하려면 package 이름을 blank identifier로 변경하자.

```go
import _ "net/http/pprof"
```

이 import 형식은 package를 부작용을 위하여 import하고 다른 용도로 사용할 수 없음을 분명히 한다. 이 파일에는 package의 이름이 없다. (만약 우리가 그 이름을 사용하지 않는다면, 컴파일러는 프로그램을 거부할 것이다.)

## Interface checks

위의 인터페이스에 대한 논의에서 보았듯이, 어떤 타입이 인터페이스를 구현한다고 명시적으로 선언할 필요는 없다. 대신, 타입은 인터페이스의 메서드를 구현함으로써 인터페이스를 구현한다. 실제로 대부분의 인터페이스 변환은 정적이므로 컴파일 타임에 확인된다. 예를 들어, `*os.File`을 `io.Reader`를 기대하는 함수에 전달하는 것은 `*os.File`이 `io.Reader` 인터페이스를 구현하지 않는다면 컴파일되지 않는다.

그러나 일부 인터페이스 검사는 런타임에 수행된다. 한 인스턴스는 `encoding/json` package에 있으며, 이 package는 `Marshaler` 인터페이스를 정의한다. JSON 인코더가 해당 인터페이스를 구현하는 값을 수신하면 인코더는 값의 mashaling 메서드를 호출하여 표준 변환을 수행하는 대신 JSON으로 변환한다. 인코더는 다음과 같은 type assertion으로 런타임에 이 속성을 확인한다.

```go
m, ok := val.(json.Marshaler)
```

오류 확인의 일부로 인터페이스 자체를 실제로 사용하지 않고 타입이 인터페이스를 구현하는지 여부만 궁금한 경우, blank identifier를 사용하여 type-asserted 값을 무시한다.

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```

이러한 상황이 발생하는 한 장소는 실제로 타입을 구현하는 package 내에서 인터페이스를 만족시키는지 보증할 필요가 있는 때이다. 예를 들어 `json.RawMessage`의 타입은 사용자 정의 JSON 표현이 필요하며, json을 구현해야 한다. `Marshaler`, 그러나 컴파일러가 이를 자동으로 확인하도록 하는 정적 변환은 없다. 타입이 실수로 인터페이스를 충족하지 못하는 경우에도 JSON 인코더는 계속 작동하지만 사용자 지정 구현을 사용하지 않는다. 구현이 올바른지 확인하려면 blank identifier를 사용하는 전역 선언을 package에서 사용할 수 있다.

```go
var _ json.Marshaler = (*RawMessage)(nil)
```

이 선언에서 `*RawMessage`를 `Marshaler`로 변환하는 것과 관련된 할당은 `*RawMessage`가 `Marshaler`를 구현하도록 요구하며, 컴파일시 해당 속성을 확인한다. `json.Marshaler` 인터페이스가 변경되면 이 package는 더 이상 컴파일되지 않으며 업데이트해야 한다는 알림을 받게 된다.

이 구조에서 blank identifier의 등장은 선언이 변수를 만들기 위한 것이 아니라 타입 확인을 위한 것임을 나타낸다. 그러나 인터페이스를 충족하는 모든 타입에 대해 이 작업을 수행하지는 마라. 관례에 따라, 이러한 선언은 이미 존재하는 정적 변환이 없을 때만 사용되며, 이는 드문 상황이다.
