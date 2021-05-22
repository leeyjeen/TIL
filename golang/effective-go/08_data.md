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

## Arrays

## Slices

## Two-dimensional slices

## Maps

## Printing

## Append

