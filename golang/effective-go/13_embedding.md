# Embedding

Go는 서브클래싱의 전형적이고 타입 중심적인 개념을 제공하지 않지만, 구조체 또는 인터페이스를 타입에 포함시킴으로써 구현의 일부를 "차용"할 수 있는 능력을 가지고 있다.

인터페이스 임베딩은 매우 간단하다. 우리는 `io.Reader`와 `io.Writer` 인터페이스에 대해 전에 언급했다. 다음은 그들의 정의이다.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`io` package는 또한 이러한 여러 메서드를 구현할 수 있는 개체를 지정하는 여러 다른 인터페이스도 export한다. 예를 들어, `io.ReadWriter`가 있다. `Read`, `Write` 모두 포함하는 인터페이스이다. 두 메서드를 명시적으로 나열하여 `io.ReadWriter`를 명시할 수 있다. 그러나 다음과 같이 두 인터페이스를 포함시켜 새로운 인터페이스를 구성하는 것이 더 쉽고 자극적이다.

```go
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

`ReadWriter`는 `Reader`와 `Writer`의 작업을 수행할 수 있다. 임베디드 인터페이스의 결합이다. 인터페이스에는 인터페이스만 포함할 수 있다.

동일한 기본 아이디어가 구조체에 적용되지만, 보다 광범위한 의미를 내포하고 있다. `bufio` package에는 두 가지 구조체 타입인 `bufio`가 있다. `Reader`와 `bufio.Writer`, 각각은 물론 `io` package에서 유사한 인터페이스를 구현한다. 또한 `bufio`는 buffered reader/writer를 구현하는데, 이는 임베딩을 사용하여 reader와 writer를 하나의 구조체로 결합함으로써 이루어진다. 구조체 내에 타입을 나열하지만 필드 이름은 지정하지 않는다.

```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```

내장된 구성 요소는 구조체에 대한 포인터이며, 사용되기 전에 먼저 유효한 구조체를 가리키도록 초기화되어야 한다.  `ReadWriter` 구조체는 다음과 같이 작성되어야 한다.

```go
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```

그러나 필드의 메서드를 홍보하고 io 인터페이스를 만족시키기 위해서는 다음과 같은 포워딩 메서드도 제공해야 한다.

```go
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

구조체를 직접 내장함으로써, bookkeeping을 피한다. 내장된 타입의 메서드들은 무상으로 따라오는데, 이는 `bufio.ReadWriter`는 `bufio.Reader`와 `bufio.Writer`의 메서드 뿐만 아니라, 다음 세 인터페이스도 만족함을 의미한다. `io.Reader`, `io.Writer`, `io.ReadWriter`.

임베딩이 서브클래싱과 다른 중요한 방식이 있다. 하나의 타입을 임베딩할 때, 그 타입의 메서드는 외부 타입의 메서드가 되지만 그들이 호출될 때 그 메서드의 리시버는 외부 타입이 아니라 내부 타입이다. 이 예시에서, `bufio.ReadWriter`의 `Read`가 호출될 때, 위에서 작성된 포워딩 메서드와 동일한 효과를 갖는다. 리시버는 `ReadWriter` 자체가 아니라 `ReadWriter`의 reader 필드이다.

임베딩은 또한 간단한 편리함이 될 수 있다. 다음 예시는 이름이 지정된 일반 필드와 함께 임베딩된 필드를 보여준다.

```go
type Job struct {
    Command string
    *log.Logger
}
```

`Job` 타입은 `Print`, `Printf`, `Println`와 `*log.Logger`의 다른 메서드들을 갖는다. 물론 `Logger`에 필드 이름을 붙일 수도 있었지만 그렇게 할 필요는 없다. 이제 초기화되면 `Job`에 로그를 기록할 수 있다.

```go
job.Println("starting now...")
```

`Logger`는 `Job` 구조체의 정규 필드이기 때문에 `Job`의 생성자 내에서 일반적인 방법으로 초기화할 수 있다. 다음과 같다.

```go
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```

또는 합성 리터럴을 사용할 수 있다.

```go
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```

우리가 직접 임베디드 필드를 참조할 필요가 있다면, package 한정자를 무시하는 필드의 타입명은 `ReadWriter` 구조체의 `Read` 메서드에서와 같이 필드 이름으로 기능한다. `Job` 변수 job의 `*log.Logger`에 접근해야 한다면, `job.Logger`을 작성할 수 있다. `Logger` 메서드를 재정의하고 싶다면 유용할 것이다.

```go
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
```

타입을 임베딩하는 것은 이름 충돌 문제를 일으키지만 이를 해결하기 위한 규칙은 간단하다. 먼저 필드 또는 메서드 X는 타입의 더 깊이 중첩된 부분에 있는 다른 항목 X를 숨긴다. `log.Logger`에 `Command`라는 필드 또는 메서드가 포함되어 있으면 `Job`의 `Command` 필드가 이를 지배한다.

둘째, 동일한 이름이 동일한 중첩 수준에 나타나면 대개 오류이다. `Job` 구조체에 `Logger`라는 다른 필드 또는 메서드가 포함된 경우 `log.Logger`를 내장하면 오류가 발생할 것이다. 

그러나 타입 정의 외부에 있는 프로그램에서 중복 이름이 언급되지 않는 경우는 괜찮다. 이 자격은 외부에서 임베디드된 타입에 대한 변경으로부터 일부 보호를 제공한다. 두 필드 중 어느 필드도 사용되지 않는 경우 다른 하위 타입의 다른 필드와 충돌하는 필드가 추가되더라도 문제가 없다.