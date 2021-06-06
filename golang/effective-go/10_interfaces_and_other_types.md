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

## Generality

## Interfaces and methods
