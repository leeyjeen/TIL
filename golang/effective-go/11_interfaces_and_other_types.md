# Interfaces and other types

## Interfaces

Go언어에서 인터페이스는 객체의 행위를 지정하는 방법을 제공한다. 어떤 것이 `this`를 사용할 수 있다면, 여기서(here) 사용될 수 있다는 것을 의미한다. 여러 간단한 예시들을 이미 보아 왔다. 커스텀 프린터는 `String` 메서드를 통해 구현될 수 있다. `Fprintf`는 `Write` 메서드를 가진다면 출력을 생성할 수 있다. 하나 또는 두 개의 메서드만을 갖는 인터페이스는 Go 코드에서 일반적이며, 메서드로부터 이름을 따오는 경우가 많다. 예를 들면, `Write`를 구현하는 `io.Writer`가 그 경우이다.

하나의 타입은 여러 인터페이스를 구현할 수 있다. 예를 들어, 컬렉션이 `sort.Interface`(`Len()`, `Less(i, j int) bool`, `Swap(i, j int)`, 그 외 커스텀 formatter를 포함하는)를 구현하고 있다면 그 컬렉션은 sort 패키지로 정렬될 수 있다. 아래 예시 Sequence는 두 가지를 모두 만족한다.

```go
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

## Conversions

위의 Sequence의 `String` 메서드는 `Sprintf`가 슬라이스에 이미 하고 있는 기능을 다시 만든 것이다. (이 또한 시간 복잡도가 O(N²)라서 좋지 않다.) `Sprint`를 호출하기 전에 Sequence를 평범한 []int로 변환한다면 수고가 줄어 들고 속도를 빠르게 할 수 있다.

```go
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

위의 메서드는 `String` 메서드에서 `Sprintf`를 안전하게 호출하기 위한 변환 기술의 한 예시이다. 두 가지 타입 (Sequence, []int)는 타입명을 무시한다면 동일하기 때문에, 둘 사이에 변환을 수행하는 것은 문제가 되지 않는다. 변환은 새로운 값을 생성하지 않으며, 임시로 기존 값이 새로운 타입을 갖는 것처럼 동작하게 한다. (integer에서 floating point로 변환처럼 다른 합법적인 변환도 있다. 이는 새 값을 생성한다.)

다른 메서드 집합에 접근하기 위해 표현의 타입을 변환하는 것은 Go 프로그램에서 많이 쓰인다. 예를 들어, `sort.IntSlice`라는 타입을 사용하여 전체 예시를 다음과 같이 줄일 수 있다.

```go
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

이제, `Sequence`가 여러 인터페이스를 구현(sorting, printing)하게 하는 대신 하나의 데이터 항목이 여러 타입(`Sequence`, `sort.IntSlice`, `[]int`)으로 변환될 수 있는 능력을 이용한다. 이런 각 타입이 작업의 일정 부분을 담당하게 된다. 일반적으로 사용되지는 않지만 효율적일 수 있다.

## Interface conversions and type assertions

타입 스위치(Type switch)는 변환의 한 형태이다. 인터페이스를 받아서 switch 구문의 각 case마다 돌면서 case의 타입으로 변환해준다. 다음은 `fmt.Printf` 코드가 타입 스위치를 이용하여 값을 문자열로 바꾸는 방법의 간단한 버전을 보여준다. 만약 이미 문자열이라면, 인터페이스의 실제 문자열 값을, String 메서드를 갖고 있다면, 메서드 호출 결과를 사용한다.

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

첫 번째 case는 구체적인 값을 찾는다. 두 번째 case는 인터페이스를 또다른 인터페이스로 변환한다. 이런 방법으로 타입을 섞어서 사용해도 상관 없다.

우리가 신경 써야 하는 타입이 하나만 있다면 어떻게 할까? 값이 string 타입인 것을 알고 있고 단지 추출하고 싶을 경우는? case가 하나인 타입 스위치를 사용할 수 있지만, type assertion을 사용할 수 있다. type assertion은 인터페이스 값을 받아서 이로부터 명시된 타입의 값을 추출한다. 문법은 타입 스위치로부터 빌려왔지만, type 키워드 대신 타입을 명시하여 사용한다.

```go
value.(typeName)
```

결과로 `typeName`이라는 정적 타입의 새로운 값이 나온다. 해당 타입은 인터페이스로 전달된 구체적인 타입이거나, 값이 변환될 수 있는 두 번째 인터페이스 타입이어야 한다. 값에서 문자열을 추출하기 위해 다음과 같이 작성할 수 있다.

```go
str := value.(string)
```

값이 문자열을 포함하지 않는다면 프로그램은 런타임 에러가 발생할 것이다. 이를 방지하기 위해 "comma, ok" 구문을 사용하여 값이 문자열인지 안전하게 테스트하자.

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

type assertion이 실패한다면, `str`은 여전히 존재하며 string 타입이지만, zero 값인 비어 있는 문자열이 될 것이다.

용량(capability)에 대한 설명으로서, 다음은 이번 섹션의 앞에 나왔던 타입 스위치와 동일한 if-else 구문이다.

```go
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```

## Generality

어떤 타입이 인터페이스를 구현하기 위해서만 존재하고 인터페이스의 어떤 메서드도 외부에 노출(export)하지 않는다면, 해당 타입은 노출시킬 필요가 없다. 인터페이스만 노출시킨다면, 값이 인터페이스에 설명된 것 이상의 흥미로운 동작을 하지 않는다는 것을 알 수 있다. 또한 공통된 메서드의 모든 인스턴스에 문서를 반복해서 작성할 필요가 없어진다.

이러한 경우, 생성자는 구현 타입보다는 인터페이스 값을 반환해야 한다. 예를 들어, 해시 라이브러리의 `crc32.NewIEEE`와 `adler32.New`는 둘 다 인터페이스 타입 `hash.Hash32`를 반환한다. Go 프로그램에서 Adler-32에 대한 CRC-32 알고리즘을 대체하려면 생성자 호출만 변경하면 된다. 나머지 코드는 알고리즘 변경에 영향받지 않는다.

유사한 접근 방식을 사용하면 다양한 `crypto` 패키지의 스트리밍 암호 알고리즘을 함께 연결하는 블록 암호와 분리할 수 있다. `crypto/cipher` 패키지의 블록 인터페이스는 단일 데이터 블록의 암호화를 제공하는 블록 암호의 동작을 지정한다. 그런 다음, `bufio` 패키지와 유사하게, 이 인터페이스를 구현하는 암호 패키지는 블록 암호화의 세부 사항을 알지 못한 채 스트림 인터페이스로 표현되는 스트리밍 암호를 구성하는 데 사용될 수 있다.

`crypto/cipher` 인터페이스는 다음과 같다.

```go
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
```

다음은 블록 암호를 스트리밍 암호로 변환하는 카운터 모드(CTR) 스트림의 정의이다. 블록 암호의 세부 정보는 추상화된다.

```go
// NewCTR returns a Stream that encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewCTR(block Block, iv []byte) Stream
```

`NewCTR`는 하나의 특정한 암호화 알고리즘과 데이터 소스에만 적용되지 않고, 블록 인터페이스와 스트림의 모든 구현에 적용된다. 이는 인터페이스 값을 반환하기 때문에, 다른 암호화 모드를 CRT 암호화로 대체하는 것은 로컬 변화이다. 생성자 호출은 수정되어야 하지만, 둘러싼 코드가 결과를 `Stream`으로만 간주해야 하므로, 차이를 인식하지 못한다.

## Interfaces and methods

거의 모든 것은 메서드를 붙일 수 있기 때문에, 즉 모든 것은 인터페이스를 만족시킬 수 있다. 하나의 예시가 `Handler` 인터페이스를 정의하는 `http` 패키지에 있다. `Handler`를 구현하는 모든 개체는 HTTP 요청을 처리할 수 있다.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`ResponseWriter`는  클라이언트에게 응답을 반환하기 위해 필요한 메서드에 대한 액세스를 제공하는 인터페이스이다. 해당 메서드들은 표준 `Write` 메서드를 포함하므로, `http.ResponseWriter`는 `io.Writer`가 사용될 수 있는 곳이면 어디서나 사용할 수 있다. `Request`는 클라이언트로부터 파싱된 표현을 포함하는 구조체이다.

간단하게, POST를 무시하고 HTTP 요청이 항상 GET이라고 가정해보자. 이러한 단순화는 핸들러 설정 방식에 영향을 미치지 않는다. 다음은 페이지 방문 횟수를 계산하는 간단한 처리기 구현이다.

```go
// Simple counter server.
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

(우리 주제를 가지고, `Fprintf`가 `http.ResponseWriter`로 출력하는 방법을 참고해보자.) 실제 서버에서 `ctr.n`에 대한 접근은 동시 접근으로부터 보호가 필요하다. 자세한 내용은 `sync`, `atomic` 패키지를 참조하자.

참고로 URL 트리의 노드에 서버를 연결하는 방법은 다음과 같다.

```go
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

하지만 왜 `Counter`를 구조체로 만들었을까? 정수 하나만이 필요한게 다인데. (리시버는 포인터여야 호출자가 increment를 볼 수 있다.)

```go
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

만약 프로그램에서 페이지 방문에 대한 알림을 받아야 하는 내부 상태값들을 가진다면 어떻게 해야 할까? 웹페이지에 채널(channel)을 연결하자.

```go
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```

마지막으로, 서버 바이너리를 호출할 때 사용되는 인수를 `/args`에 표시하려고 한다고 해보자. 인수를 출력하는 함수를 작성하는 것은 쉽다.

```go
func ArgServer() {
    fmt.Println(os.Args)
}
```

어떻게 이것을 HTTP 서버로 바꿀 수 있을까? `ArgServer`를 우리가 무시하는 값의 어떤 타입의 메서드로 만들 수 있지만, 더 깔끔한 방법이 있다. 포인터와 인터페이스를 제외한 어떤 타입에 대해서도 메서드를 정의할 수 있기 때문에, 함수에 대한 메서드를 작성할 수 있다. `http` 패키지는 다음 코드를 포함하고 있다.

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```

`HandlerFunc`는 메서드 `ServeHTTP`를 갖는 타입으로, 해당 타입의 값들은 `HTTP` 요청을 처리할 수 있다. 메서드의 구현을 보자. 리시버는 함수 f이고, 메서드는 f를 호출한다. 이상하게 보일 수 있지만, 이는 리시버가 채널로서 메서드가 채널에 메시지를 전송하는 것과 다를 것이 없다.

`ArgServer`를 `HTTP` 서버로 만들기 위해, 우선 적절한 시그니처를 갖도록 수정해야 한다.

```go
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```

`ArgServer`는 이제 `HandlerFunc`와 동일한 시그니처를 가지므로, 메서드에 접근하기 위한 타입으로 변환될 수 있다. 이전에 `Sequence`를 `IntSlice.Sort`에 대한 접근을 위하여 `IntSlice`로 변환했던 것과 비슷하다. 이를 셋업하기 위한 코드는 간결하다.

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

누군가 페이지 `/args`에 방문하였을 때, 페이지에 설치된 핸들러는 `ArgServer` 값과 `HandlerFunc` 타입을 갖는다. HTTP 서버는 `ArgServer`를 리시버로 사용하여, 해당 타입의 `ServeHTTP` 메서드를 호출하고, 다시 `ArgServer`를 호출한다. (`HandlerFunc.ServeHTTP` 내부의 `f(w, req)` 호출을 통하여). 인수가 출력될 것이다.

이번 섹션에서 구조체, 정수, 채널, 함수로부터 HTTP 서버를 만들어 보았다. 인터페이스는 거의 모든 타입에 대해 정의될 수 있는 메서드의 집합이기 때문이다.