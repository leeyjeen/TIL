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

## Append

