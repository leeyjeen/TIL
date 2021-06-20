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



## Interface checks ¶

