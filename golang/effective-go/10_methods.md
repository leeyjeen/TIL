# Methods

## Pointers vs. Values
`ByteSize`에서 보았듯이, 메서드는 (포인터나 인터페이슬르 제외한) 모든 명명된 유형에 대해 정의될 수 있다.

위의 슬라이스에 대한 논의에서 `Append` 함수를 작성했다. 대신 슬라이스에 대한 메서드로 정의할 수 있다. 이를 위해 먼저 메서드를 바인딩할 수 있는 명명된 유형을 선언한 다음 메서드의 리시버를 해당 유형의 값으로 설정한다.

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
```

이는 여전히 업데이트된 슬라이스를 반환하는 메서드가 필요하다. **포인터**를 `ByteSlice`에 리시버로서 취하는 메서드를 재정의함으로써 이 메서드가 호출자의 슬라이스를 덮어쓸 수 있도록 하여 난잡함을 제거할 수 있다.

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

사실 훨씬 더 잘 할 수 있다. 만약 함수를 표준 `Write` 메서드처럼 보이게 한다면, 다음과 같이,

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

`*ByteSlice` 타입은 표준 인터페이스 `io.Writer`를 만족한다. 예를 들어, 하나로 출력할 수 있다.

```go
    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)
```


`*ByteSlice`만이 `io.Writer`를 만족시키기 때문에 `ByteSlice`의 주소를 전달시킨다. 포인터 vs 리시버 값에 대한 규칙은 값 메서드가 포인터 및 값에서 호출될 수 있지만, 포인터 메서드는 포인터에서만 호출될 수 있다.

이 규칙은 포인터 메서드가 리시버를 수정할 수 있기 때문에 발생한다. 값을 호출하면 메서드가 값의 복사본을 수신하므로 수정 사항이 삭제된다. 그러므로 언어는 이 실수를 허용하지 않는다. 그러나 편리한 예외가 있다. 값이 주소 지정이 가능할 때, 언어는 자동으로 주소 연산자를 삽입하여 값에 포인터 메서드를 호출하는 일반적인 경우를 처리하여 값의 복사본을 수신하므로 수정 사항이 무시된다. 그러므로 언어는 이 실수를 허용하지 않는다. 우리의 예시에서, 변수 b는 주소 지정 가능하므로, 우리는 단지 `b.Write`로 Write 메서드를 호출할 수 있다. 컴파일러가 이를 다시 `(&b).Write`로 쓴다.

그런데 바이트 슬라이스에 `Write`를 사용하는 아이디어는 `bytes.Buffer` 구현의 중심이다.