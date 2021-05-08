# 이름

다른 언어들에서와 마찬가지로 Go에서도 이름은 중요하다. 이름은 의미적으로 효과를 가질 수 있다. 이름의 첫 번째 문자가 대문자인지 여부에 따라서 package 외부로의 이름의 노출 여부가 결정된다. 그러므로 Go 프로그램에서 네이밍 컨벤션에 대해 살펴볼 필요가 있다.

## package 이름

package가 import 된다면, package 이름은 내용에 대한 접근자가 된다.

```go
import "bytes"
```
위와 같이 import한 다음에, import한 package는 `bytes.Buffer`를 사용할 수 있다. package를 사용하는 모든 사람이 내용을 참조하기 위해 동일한 이름을 사용할 수 있다는 것은, package 이름이 잘 작성되어야 함을 의미한다. 짧고, 간결하고, 상기시키는 이름이어야 한다. 컨벤션에 의해, package는 소문자이며, 한 단어의 이름이어야 한다. 언더스코어나 대소문자 혼용이 필요 없어야 한다. package를 사용하는 모든 사람이 해당 이름을 입력할 것이기 때문에 간결성 측면에서 중요하다. 그리고 충돌에 대해 미리 걱정하지 말자. package 이름은 오직 import를 위한 이름이다. 모든 소스 코드에서 유니크할 필요는 없으며, 드문 경우 충돌이 일어난다면 지역적으로 package을 import할 때 다른 이름을 사용할 수 있다. 어느 경우에나, import에 있는 파일 이름이 어떤 pakcage를 사용할 지 결정하기 때문에 혼란은 드물다.

또 다른 컨벤션은 package 이름이 소스 디렉토리 이름 기반이라는 것이다. `src/encoding/base64` 안의 package는 "base"라는 이름을 가지고 있지만, "encoding/base64"로 import된다. "encoding_base64"이나 "encodingBase64"으로 사용되지 않는다.

package를 import하는 입장에서 내용을 참조하여 이름을 사용할 것이므로, package에서 export되는 이름은 헷갈리는 것을 피하기 위하여 이와 같은 사실을 이용할 수 있다. (`import .` 표기법을 사용하지 마라. 이는 package 밖에서 실행해야 하는 테스트를 단순화할 수 있지만 그렇지 않다면 피해야 한다.) 예를 들어, `bufio` package의 buffered reader 형식은 `BufReader`가 아니라 `Reader`로 불린다. 왜냐하면 사용자들이 명확하고 간결한 이름인 `bufio.Reader`로 볼 수 있기 때문이다. 또한, import된 개체들이 항상 package 이름과 함께 불려지기 때문에 `bufio.Reader`는 `io.Reader`와 충돌하지 않는다. 비슷하게 `ring.Ring`의 인스턴스를 생성하는 함수는 보통 `NewRing`이라고 불린다. 그러나 `Ring`은 package에 의해서 export되는 유일한 타입이며, package는 `ring`이라고 불리기 때문에 이 함수는 그냥 `New`라고 부르고 `ring.New`와 같이 사용한다. 좋은 이름을 위해 package 구조를 활용하라.

또 다른 간단한 예시는 `once.Do`이다. `once.Do(setup)`은 쉽게 읽히며 `once.DoOrWaitUntilDone(setup)`로 개선되지 않는다. 긴 이름은 더 읽기 좋게 만들지 않는다. 문서에 주석을 다는 것이 과도하게 긴 이름보다 가치있을 것이다.

## Getters

Go는 getter와 setter를 자체적으로 제공하지 않는다. getter와 setter를 직접 만드는 것은 문제될 것이 없으며, 그렇게 하는 것이 적절하고 일반적이다. 그러나 getter의 이름에 Get을 넣는 것은 Go언어 답거나 필수적이지 않다. 만약 `owner`라는 필드(소문자, export되지 않음)를 가지고 있을 때, getter 메서드는 `GetOwner`가 아니라 `Owner`(대문자, export됨)이어야 한다. export를 위해 대문자 이름을 사용하는 것은 메서드로부터 필드를 구별하기 위한 hook을 제공한다. setter 함수는 만약 필요하다면 `SetOwner`라고 부른다. 두 이름 모두 읽기 쉽다.

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

## Interface 이름

관례적으로, 하나의 메서드를 갖는 인터페이스는 메서드 이름에 `-er` 접미사를 붙이거나 agent 명사를 생성하는 유사한 변형에 의해 명명된다. (예: `Reader`, `Writer`, `Formatter`, `CloseNotifier` 등)

이러한 이름은 여러 가지가 있으며 이를 사용하는 것은 생산적이다. Read, Write, Close, Flush, String 등은 각각 표준이 되는 표시와 의미를 가지고 있다. 혼란을 피하기 위하여 같은 용법과 의미를 갖는 것이 아니라면 이들과 같은 이름을 갖는 메서드를 생성하지 마라. 반대로, 구현한 타입이 잘 알려진 타입의 메서드와 동일한 의미를 갖는 메서드를 구현한다면, 같은 이름을 부여하면 된다. 예를 들어, 문자 변환 메서드를 만든다면 `ToString` 대신 `String`을 사용하라.

## 대소문자 혼합(MixedCaps)

마지막으로, Go에서의 관례는 여러 단어로 된 이름을 사용할 때 언더바(_) 보다는 대소문자 혼합(MixedCaps 또는 mixedCaps)을 사용하는 것이다.