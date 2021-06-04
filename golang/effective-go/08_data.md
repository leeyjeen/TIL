# Data

## Allocation with `new`

Go언어에는 두 가지 메모리 할당 방식이 있다. 내장 함수인 `new`와 `make`이다. 이들은 서로 다른 일을 하며, 다른 형식에 적용되어, 혼란스러울 수 있으나, 규칙은 간단하다. 

우선 `new`에 대해서 살펴 보자. 이는 메모리를 할당하는 내장 함수이지만, 다른 언어에서의 `new`와는 다르게 메모리를 초기화하지 않고 단지 값을 zero 값으로 만든다. 즉, `new(T)`는 `T` 형식의 새로운 객체에 대하여 zero화된 저장 공간을 할당해주고, 그 객체의 주소인 `*T`의 값을 반환한다. Go의 용어로서 말해보면, zero 값으로 새로 할당된 형식 `T`를 가리키는 포인터를 반환한다.

`new`로 반환된 메모리는 zero화되어 있기 때문에, 각 형식의 zero 값을 추후의 초기화 없이 사용될 수 있도록 자료 구조를 설계하면 유용하다. 이는 자료 구조의 사용자가 `new`를 이용하여 생성 및 작업할 수 있다는 것을 의미한다. 예를 들어, `bytes.Buffer`에 대한 문서는 "Buffer에 대한 zero값은 사용할 준비가 된 비어 있는 buffer이다." 라고 말하고 있다. 유사하게, `sync.Mutex`는 명시된 생성자나 Init 메서드를 가지고 있지 않다. 대신, `sync.Mutex`의 zero 값은 unlocked mutex로 정의되어 있다.

zero값의 유용함은 전이적인 속성이 있다. 다음의 타입 선언을 살펴 보자.

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

`SyncedBuffer` 타입의 값은 할당 또는 선언 즉시 사용할 준비가 된다. 다음 코드에서, `p`와 `v`는 추가적인 정리 없이도 올바르게 동작한다.

```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

## Constructors and composite literals

때때로 zero값으로만으로는 충분하지 않고 초기화 생성자가 필요한 경우가 있다. package `os`로부터 가져온 예시를 보자.

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

여기서 반복적인 코드들을 많이 볼 수 있다. **힙성 리터럴(composite literal)**을 사용하여 간단하게 만들 수 있다. 이는 표현식이 실행될 때마다 새로운 인스턴스를 생성한다.

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

C와는 다르게, 지역 변수의 주소를 반환하는 것에 전혀 문제가 없다. 변수에 연결된 저장 공간은 함수가 반환된 후에도 살아 남는다. 사실 합성 리터럴의 주소를 취하는 것은 실행될 때마다 새로운 인스턴스를 할당하므로, 마지막 두 라인을 아래처럼 묶을 수 있다.

```go
    return &File{fd, name, nil, 0}
```

합성 리터럴의 필드는 순서대로 배열되고 반드시 모두 입력해야 한다. 그러나, 요소들에 `field:value` pair와 같이 명백하게 레이블링을 한다면, 순서에 상관 없이 초기화할 수 있으며, 입력되지 않은 요소는 각각의 zero값으로 초기화할 수 있다.

```go
    return &File{fd: fd, name: name}
```

제한적인 경우로서, 합성 리터럴이 아무 필드도 포함하지 않는다면, 타입에 대한 zero값을 생성한다. `new(File)`와 `&File{}`은 동일한 표현이다.

합성 리터럴은 필드 레이블로 인덱스 또는 맵의 키를 적절하게 사용하여 배열, 슬라이드, 맵에 대해서도 생성할 수 있다. 예시에서, 초기화는 `Enone`, `Eio`, `Einval`의 값에 상관 없이, 그들이 서로 다르기만 하면 동작한다.

```go
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```
예를 들면 다음과 같다.
```go
a := [...]string   {0: "no error", 1: "Eio", 2: "invalid argument"}
s := []string      {0: "no error", 1: "Eio", 2: "invalid argument"}
m := map[int]string{0: "no error", 1: "Eio", 2: "invalid argument"}
```

## Allocation with `make`

메모리 할당으로 돌아가보자. 내장 함수 `make(T, args)`는 `new(T)`와 다른 목적을 제공한다. 이는 슬라이스, 맵, 채널만 생성하며, T 타입의 (*T가 아닌) 초기화된 (zero값이 아닌) 값을 반환한다. 이렇게 구별하는 이유는 이 세 타입은 사용 전에 초기화되어야 하는 데이터 구조를 참조하기 때문이다. 예를 들어, 슬라이스는 데이터를 가리키는 포인터 (배열 내부의), 길이, 수용력을 포함하는 세 가지 항목의 디스크립터를 가지며, 이 항목들이 초기화될 때까지 슬라이스는 `nil`이다. 슬라이스, 맵, 채널에 대해 `make`는 내부 데이터 구조를 초기화하고 사용할 값을 준비힌다. 예를 들어,

```go
make([]int, 10, 100)
```

는 크기가 100인 배열을 할당하고 배열의 첫 10 요소를 가리키는 길이가 10이고 수용공간이 100인 슬라이스 구조를 생성한다.(슬라이스를 만들 때, 수용공간은 생략될 수 있다. 자세한 내용은 슬라이스 섹션을 참고하자.) 반대로, `new([]int)`는 새로 할당된 zero화된 슬라이스 구조를 가리키는 포인터를 반환하는데, 즉 nil 슬라이스 값의 포인터이다.

다음 예시는 `new`와 `make`의 차이점을 나타낸다.

```go
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

`make`는 맵, 슬라이스, 채널에만 적용되며 포인터를 반환하지 않는다는 것을 기억하자. 포인터를 얻고 싶다면 `new`를 사용하여 메모리를 할당하거나 변수의 주소를 명시적으로 사용하자.

## Arrays

배열은 메모리 레이아웃을 상세하게 계획할 때 유용하며 때로는 메모리 할당을 피하는 데 도움이 되지만, 주로 다음 섹션의 주제인 슬라이스의 재료로 쓰인다. 슬라이스에 대한 초석을 다지기 위하여 배열에 대해 짧게 확인해보자.

Go와 C에서 배열의 작동에는 주요한 차이점들이 있다. Go에서,

* 배열은 값이다. 하나의 배열을 다른 배열에 할당한다면 모든 요소가 복사된다.
* 특히, 함수에 배열을 넘겨준다면, 함수는 배열 포인터가 아니라 배열의 복사본을 받는다.
* 배열의 크기는 타입의 한 부분이다. [10]int과 [20]int 타입은 서로 다르다.

값 속성은 유용할 수 있으나 비용이 크기도 하다. C처럼 동작하거나 효율적이기를 원한다면, 배열 포인터를 전달할 수 있다.

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
```

그러나 이 스타일은 Go스럽지 않다. 대신 슬라이스를 사용해보자.

## Slices

슬라이스는 데이터 시퀀스에 더 일반적이고, 강력하고, 편리한 인터페이스를 제공하기 위하여 배열을 포장한다. 변환 매트릭스와 같은 명백한 차원의 항목들을 제외하고, Go에서 대부분의 배열 프로그래밍은 단순한 배열보다는 슬라이스를 이용하여 이루어진다.

슬라이스는 내부의 배열을 가리키는 참조를 쥐고 있으며, 한 슬라이스를 다른 슬라이스에 할당한다면, 둘은 동일한 배열을 참조한다. 함수가 슬라이스 매개변수를 취하였을 때, 슬라이스 요소에 변화를 줄 경우 호출자도 확인할 수 있다. 내부 배열을 가리키는 포인터를 전달하는 것과 유사하다. 따라서 `Read` 함수는 포인터와 카운트 대신 슬라이스 매개변수를 받을 수 있다. 슬라이스 길이는 읽을 수 있는 데이터의 상한으로 설정된다. 다음은 `os` 패키지의 `File`의 `Read` 메서드의 시그니처이다.

```go
func (f *File) Read(buf []byte) (n int, err error)
```
이는 읽은 바이트 수와 에러가 있다면 에러 값을 반환한다. 더 큰 버퍼 buf의 처음 32 바이트를 읽어 들이기 위해 버퍼를 슬라이싱한다.

```go
    n, err := f.Read(buf[0:32])
```

이러한 슬라이싱은 일반적으로 쓰이며 효율적이다. 사실 효율성을 잠시 제쳐둔다면, 다음 코드 조각 또한 버퍼의 처음 32 바이트를 읽어 들인다.

```go
    var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
```

슬라이스의 길이는 내부 배열의 한계 내에서는 바뀔 수 있다. 슬라이스 자체에 할당하면 된다. 슬라이스의 용량은 내장 함수 `cap`으로 접근 가능하며, 슬라이스가 취할 수 있는 최대 길이를 알려 준다. 다음은 슬라이스에 데이터를 추가하기 위한 함수이다. 데이터가 용량을 초과한다면, 슬라이스는 재할당된다. 결과 슬라이스가 반환된다. 함수는 `len`과 `cap`는 `nil` 슬라이스에 사용할 수 있으며, 이 경우 0을 반환한다는 사실을 이용한다.

```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

슬라이스는 처리 후 반환되어야 한다. `Append`가 슬라이스의 요소를 변경할 수 있지만, 슬라이스 자체(포인터, 길이, 용량을 쥐고 있는 런타임 데이터 구조)가 값으로 전달되기 때문이다.

슬라이스에 append하는 아이디어는 매우 유용하여 `append` 내장 함수로 만들어져 있다. 함수의 설계를 이해하기 위해서 더 많은 정보가 필요하므로, 나중에 살펴 보도록 하자.

## Two-dimensional slices

Go의 배열과 슬라이스는 1차원이다. 2차원 배열이나 슬라이스와 동일한 것을 만들기 위해서는, 다음과 같이 배열의 배열 또는 슬라이스의 슬라이스를 정의하여야 한다.

```go
type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte     // A slice of byte slices.
```

슬라이스는 크기가 변할 수 있기 때문에, 각 내부 슬라이스가 다른 길이를 가질 수 있다. `LinesOfText` 예시에서처럼 일반적인 상황이다. 각 라인은 독립적인 길이를 갖는다.

```go
text := LinesOfText{
	[]byte("Now is the time"),
	[]byte("for all good gophers"),
	[]byte("to bring some fun to the party."),
}
```

때때로 2차원 슬라이스를 할당해야하는 경우가 발생한다. 예를 들어, 픽셀 라인을 스캔하는 상황이다. 이 경우 두 가지 방법이 있다. 첫 번째는 각 슬라이스를 독립적으로 할당하는 것이다. 두 번째는 단일 배열을 할당하고 개별 슬라이스를 가리키게 하는 것이다. 슬라이스가 커지거나 작아질 수 있다면, 다음 라인을 덮어쓰는 것을 피하기 위하여 슬라이스는 독립적으로 할당되어야 한다. 그렇지 않다면, 한 번의 할당으로 객체를 생성하는 것이 더 효율적일 것이다. 참고로, 다음은 두 가지 방법에 대한 스케치이다. 첫 번째로, 한 번에 한 줄씩:

```go
// Allocate the top-level slice.
picture := make([][]uint8, YSize) // One row per unit of y.
// Loop over the rows, allocating the slice for each row.
for i := range picture {
	picture[i] = make([]uint8, XSize)
}
```

그리고 한 번의 메모리 할당으로, 라인으로 슬라이스하는 경우:

```go
// Allocate the top-level slice, the same as before.
picture := make([][]uint8, YSize) // One row per unit of y.
// Allocate one large slice to hold all the pixels.
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
// Loop over the rows, slicing each row from the front of the remaining pixels slice.
for i := range picture {
	picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

## Maps

맵은 편리하고 강력한 내장 데이터 구조로서 하나의 타입(key)의 값들을 다른 타입(element 또는 value)의 값들에 연결해준다. key는 동등 연산자가 정의되는 어떤 타입이어도 가능하다. (예: integers, floating point, complex numbers, strings, pointers, interfaces, structs, arrays) 슬라이스는 맵의 key로서 사용될 수 없다. 슬라이스에 대해서는 동등이 정의되지 않기 때문이다. 슬라이스와 같이, 맵은 내부 데이터 구조에 참조를 쥐고 있다. 맵의 내용을 변경하는 함수에 맵을 전달한다면, 해당 변경은 호출자에게도 보인다.

맵은 콜론으로 구분된 key-value 쌍의 일반적인 합성 리터럴 문법을 사용하여 생성될 수 있으므로, 초기화와 동시에 생성하는 것이 쉽다.

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

맵에 값을 할당하고 읽어오는 것은 인덱스가 정수일 필요가 없다는 것을 제외하고는 문법적으로 배열, 슬라이스와 거의 동일하다.

```go
offset := timeZone["EST"]
```

맵에 존재하지 않는 key로 value를 추출하려고 한다면 맵의 엔트리 타입에 해당하는 zero값을 반환한다. 예를 들어 맵이 정수를 포함한다면, 존재하지 않는 key를 조회할 때 0을 반환한다. value 타입을 `bool`로 맵을 사용함으로써 집합을 구현할 수 있다. 집합에 값을 넣기 위하여 맵 엔트리를 `true`로 설정하고, 간단한 인덱싱을 통해 테스트해보자.

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

때때로 zero값과 존재하지 않는 값을 구별해야 할 수 있다. "UTC"에 대한 엔트리 값인지, 맵에 존재하지 않기 때문에 0인지? 복수 할당의 형태로 구별할 수 있다.

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

명백한 이유로 이는 "comma ok" 관용구라고 부른다. 이 예시에서, `tz`가 존재한다면, `seconds`는 적절히 설정되고 `ok`는 `true`가 될 것이다. 존재하지 않는다면, `seconds`는 zero로 설정되고 `ok`는 `false`가 될 것이다. 다음은 좋은 에러 보고와 함께 comma ok 관용구를 사용한 함수의 예시이다.

```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

실제 값에 상관없이 맵에 존재 여부를 테스트하기 위해서, blank 식별자 (_)을 값의 변수 위치에 사용할 수 있다.

```go
_, present := timeZone[tz]
```

맵 엔트리를 삭제하기 위해, `delete` 내장 함수를 사용해보자. 매개변수는 맵과 지우려는 key이다. key가 맵에서 이미 존재하지 않는 경우에도 안전하게 사용할 수 있다.

```go
delete(timeZone, "PDT")  // Now on Standard Time
```

## Printing

Go에서 포맷된 출력은 C의 `printf`와 유사하지만 더 풍부하고 일반적이다. 함수들은 `fmt` 패키지에 있으며, 대문자 이름을 갖는다. `fmt.Printf`, `fmt.Fprintf`, `fmt.Sprintf` 등등. 문자열 함수(`Sprintf` 등)는 제공된 버퍼를 채우기 보다는 문자열을 반환한다.

반드시 형식 문자열을 제공할 필요는 없다. `Printf`, `Fprintf`, `Sprintf` 각각에 대해 짝을 이루는 다른 함수가 존재하는데, 예를 들어 `Print`, `Println`이다. 이 함수들은 형식 문자열을 취하지 않는 대신 각 매개변수에 대하여 기본 형식을 생성한다. `Println` 버전들은 매개변수 사이에 공백을 삽입하고 출력에 줄바꿈을 추가한다. `Print` 버전들은 인수 양쪽이 다 문자열이 아닌 경우에만 공백을 삽입한다. 아래 예시에서 각 라인은 동일한 출력을 만든다.

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

포맷팅된 print 함수 `fmt.Fprint`와 유사 함수들은 첫 번째 매개변수로 `io.Writer` 인터페이스를 구현한 객체를 취한다. `os.Stdout`와 `os.Stderr` 변수는 익숙한 예시이다.

여기서부터는 C로부터 갈라지기 시작한다. 첫째, `%d`와 같은 숫자 형식은 부호 플래그나 크기를 취하지 않는 대신, 출력 루틴은 이 속성을 정하기 위한 매개변수 타입을 사용한다.

```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```
위 예시는 다음과 같이 출력한다.

```
18446744073709551615 ffffffffffffffff; -1 -1
```
정수 소수점과 같이 기본 변환을 원하는 경우, 다목적의 형식 `%v`("value"를 위한)을 사용할 수 있다. 결과는 `Print`, `Println`의 출력과 같다. 또한, 해당 포맷은 어떤 값이어도 출력할 수 있으며, 배열 슬라이스, 구조체, 맵도 출력한다. 다음은 이전 섹션에서 정의한 타임존 맵에 대한 print 구문이다.

```go
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
```

이는 다음과 같이 출력한다.

```
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

맵의 경우, `Printf`와 유사 함수들은 출력을 key 사전순으로 정렬한다.

구조체를 출력할 때, 수정된 형식 `%+v`는 구조체의 필드를 필드명으로 주석을 달아주고, `%#v`는 어떤 값이든 완전한 Go 문법으로 값을 출력한다.

```go
type T struct {
    a int
    b float64
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```

```
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```
(`&`에 주목하라.) 인용된 문자열 포맷은 `%q`을 이용해 string 또는 []byte 타입의 값에 적용될 때 가능하다. 대안 포맷 `%#q`는 가능하다면 backquote를 사용한다.(`%q` 포맷은 integer, rune에 적용되며, single-quoted rune 상수를 만든다.) 또한, `%x`는 string, byte 배열, byte 슬라이스, integer에 적용되여 긴 16진수 문자열을 생성하는데, (`% x`) 포맷처럼 스페이스를 중간에 넣으면 출력하는 byte 사이에 공백을 넣어 준다.

또다른 유용한 포맷으로 `%T`는 값의 타입을 출력한다.

```go
fmt.Printf("%T\n", timeZone)
```
위의 예시는 다음을 출력한다.
```
map[string]int
```

커스텀 타입의 기본 포맷을 제어하기 위해서 필요한 것은 타입에 대한 string 시그니처 `String()` 메서드를 정의하는 것이다. 우리가 정의한 단순한 타입 T에 대하여는 아래와 같은 포맷을 가질 수 있다.

```go
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
```
다음과 같이 출력한다.
```
7/-2.35/"abc\tdef"
```
(타입 `T` 값과 `T`에 대한 포인터도 함께 출력해야하는 경우, `String`의 리시버는 값 타입이어야 한다. 예시에서는 포인터를 사용했는데, struct 타입에 대해서 더 효율적이고 Go언어답기 때문이다.)

우리가 정의한 String 메서드는 `Sprintf`를 호출할 수 있다. print 루틴이 재진입이 충분히 가능하고 예시와 같이 감싸도 되기 때문이다. 하지만 이 방식에 대해 한가지 이해하고 넘어가야 하는 매우 중요한 디테일이 있다. String 메서드를 만들면서 Sprintf를 호출할 때 다시 String 메서드로 영구적으로 재귀하는 방식을 사용하면 안 된다. 이는 Sprintf 호출이 리시버를 string처럼 직접 출력하는 경우 발생할 수 있는데, 그렇게 되면 다시 같은 메서들르 호출하게 될 것이다. 이는 다음 예시에서 보는 것처럼 흔하고 쉽게 하는 실수이다.

```go
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
}
```

이러한 실수는 쉽게 고칠 수 있다. 인수를 메서드를 갖지 않는 기본적인 문자열 타입으로 변환하는 것이다.

```go
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
```

initialization 섹션에서 재귀를 피하기 위한 또다른 기술을 볼 것이다.

또다른 출력 기법으로서 print 루틴의 인수를 직접 다른 루틴으로 전달하는 것이 있다. `Printf`의 시그니처는 마지막 인수로서 타입 `...interface{}`을 사용하여 임의의 숫자의 파라미터가 포맷 다음에 나타날 수 있음을 명시한다.

```go
func Printf(format string, v ...interface{}) (n int, err error) {
```

`Printf` 함수 내에, `v`는 타입 `[]interface{}`의 변수처럼 동작하지만 만약 다른 가변 인수 함수에 대입되면, 보통의 인수 리스트처럼 동작한다. 다음은 위에서 사용된 `log.Println` 함수의 구현이다. 실제 포맷팅을 위하여 인수를 `fmt.Sprintln`에 직접 전달한다.

```go
// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output takes parameters (int, string)
}
```

`Sprintln`의 중첩된 호출에서 `v` 다음에 `...`를 적음으로써 컴파일러에게 `v`를 인수 리스트로 취급하라고 말해준다. 그렇지 않으면 `v`를 하나의 slice 인수로 전달한다.

여기서 살펴본 출력에 관한 내용보다 훨씬 더 많은 내용이 있다. `fmt` 패키지에 대한 godoc 문서를 통해 자세한 내용을 확인하자.

그런데, `...` 파라미터는 특정한 타입을 가질 수도 있다. 예를 들어, 정수 리스트의 최소값을 고르는 `min` 함수의 `...int`를 보자.

```go
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

## Append

이제 `append` 내장함수의 설계를 설명한다. `append`의 시그니처는 위의 커스텀 `Append` 함수와 다르다. 도식적으로 다음과 같다.

```go
func append(slice []T, elements ...T) []T
```

`T`는 어떤 타입의 플레이스홀더이다. Go에서는 호출자에 의해 결정되는 타입 T를 쓰는 함수를 만들 수 없다. 그래서 append는 내장함수인 것이다. 컴파일러의 지원이 필요하다.

append가 하는 일은 slice의 끝에 요소들을 붙이고 결과를 반환하는 것이다. 결과는 반환되어야 하는데 왜냐하면, 손으로 작성한 `Append`와 같이 내부의 배열은 변할 수 있기 때문이다. 간단한 예시를 보자.

```go
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

```
[1 2 3 4 5 6]
```
append는 `Printf`처럼 임의의 수의 인수를 모으며 작동한다.

그러나 만약 `Append`와 같이 slice에 slice를 붙이고 싶다면 어떻게 해야 할까? 쉽다. 위에서 Output을 호출하면서 그랬듯이 호출 위치에서 `...`를 사용하는 것이다. 아래 코드는 위와 동일한 결과를 생성한다.

```go
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```
`...`이 없다면 타입이 틀리기 때문에 컴파일되지 않는다. `y`는 int 타입이 아니다.