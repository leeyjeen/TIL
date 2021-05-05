# 주석

Go는 C 스타일의 /* */ 블록 주석과 C++ 스타일의 // 라인 주석을 제공한다. 라인 주석이 일반적이고, 블록 주석은 대부분 package 주석으로서 사용되지만, 표현식 안에서 또는 많은 코드를 주석 처리할 때 유용하다.

프로그램이면서 웹 서버인 godoc은 package 내용에 대한 문서를 추출하도록 Go 소스 파일을 처리한다. 최상위 선언 이전에 개행 없이 나타나는 주석은 해당 항목의 설명으로서 선언과 함께 추출된다. 이러한 주석의 본성과 스타일은 godoc이 생성하는 문서의 퀄리티를 결정한다.

모든 package는 package 구문 전에 블록 주석 형태의 package 주석이 있어야 한다. 여러 파일로 구성된 package의 경우, package 주석은 하나의 파일에만 존재해야 하며, 그것이 사용된다. package 주석은 package를 소개해야 하고, 전체 패키지에 관련된 정보를 제공해야 한다. 이는 godoc 페이지의 맨 앞에 나타나며, 이후의 상세 문서도 작성해야 한다.

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

package가 단순하다면, package 주석 또한 간단할 수 있다.

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

주석은 별 배너와 같은 과한 포맷은 필요하지 않다. 생성된 출력이 고정된 너비의 폰트 안에서 표현되지 않을 수 있으므로, 스페이스나 정렬 등에 의존하지 말자. 주석은 해석되지 않는 일반 텍스트이다. HTML과 _this_같은 다른 주석은 말 그대로 생성될 것이며 사용하지 않는 것이 좋다. godoc이 하는 한 가지 수정은 들여쓰기된 텍스트를 고정폭 폰트로 보여주는 것으로, 프로그램 코드조각 같은 것에 적합하다. fmt package의 package 주석이 좋은 예이다.

컨텍스트에 따라서, godoc은 주석을 다시 포맷팅하지 않을 수 있다. 그러므로 좋게 보일 수 있도록 확실히 하도록 하자. 정확한 스펠링, 구두법, 문장 구조를 사용하라. 긴 라인을 줄여라.

package 내부에서 최상위 선언 이전의 주석은 선언에 대한 문서 주석으로 제공된다. 프로그램 안에서 모든 export된 (대문자의) 이름은 문서 주석이 필요하다.

문서 주석은 다양한 자동화된 표현을 허용하는 완전한 문장일 때 가장 효과가 좋다. 첫 문장은 이름의 선언으로 시작하는 한 문장의 요약이어야 한다.

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

만약 모든 문서 주석이 해당 항목의 이름으로 시작한다면, 문서 서브 커맨드를 go tool로 사용할 수 있고, grep을 통해 출력을 실행할 수 있다. 정규식을 파싱하는 함수를 찾고 있을 때 "Compile"이라는 이름을 기억하지 못하여서 다음과 같은 명령어를 실행하였다고 상상해보라.

```
$ go doc -all regexp | grep -i parse
```

만약 package 내부의 모든 문서 주석이 "This function..."으로 시작하였다면, grep은 이름을 기억하도록 돕지 못할 것이다. 그러나 package의 각 문서 주석이 이름으로 시작하기 때문에, 이와 같이 찾고 싶은 단어를 상기시키는 결과를 볼 수 있을 것이다.

```
$ go doc -all regexp | grep -i parse
    Compile parses a regular expression and returns, if successful, a Regexp
    MustCompile is like Compile but panics if the expression cannot be parsed.
    parsed. It simplifies safe initialization of global variables holding
```

Go의 선언 문법은 선언의 그룹화가 가능하다. 단일 문서 주석은 연관된 상수나 변수의 그룹을 소개할 수 있다. 전체 선언이 제공되기 때문에 그러한 주석은 형식적일 수 있다.

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

그룹화는 또한 변수 집합이 뮤텍스에 의해 보호된다는 사실과 같은 항목들 사이의 관계를 나타낼 수 있다.

```go
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```