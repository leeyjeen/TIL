# 소개

Go는 새로운 언어이다. 비록 기존 언어들의 아이디어를 차용했으나, 동류의 언어로 작성된 프로그램과 성격이 다른 효과적인 Go 프로그램을 만들어주는 특이한 속성을 갖는다. C++ 또는 Java 프로그램을 Go로 직역하면 만족스러운 결과를 만들지 못할 것이다. Java 프로그램은 Java로 작성되었지, Go로 작성되지 않았다. 반면에, Go의 관점으로 문제를 생각해보면, 성공적이지만 상당히 다른 프로그램을 만들 수 있다. 다시 말해, Go를 잘 작성하기 위해서는 Go의 속성과 표현 양식을 잘 이해하는 것이 중요하다. 또한 작성한 프로그램을 다른 Go 프로그래머들이 이해하기 쉽도록, Go 프로그래밍에 자리잡은 컨벤션을 아는 것도 중요하다.

이 문서는 명확하고, 자연스러운 Go 코드를 작성하는 팁을 제공한다. [언어 명세](https://golang.org/ref/spec), [Tour of Go](https://tour.golang.org/), [Go 코드 작성법](https://golang.org/doc/code)과 우선적으로 읽어야 할 모든 지식들을 보강한다.

## 예제

[Go 패키지 소스](https://golang.org/src/)는 코어 라이브러리로 뿐 아니라 언어 사용법 예제로써도 제공될 수 있도록 만들어졌다. 더 나아가, 많은 패키지들이 동작하는, 독립적으로 실행 가능한 예제들을 포함하고 있으며 [golang.org](https://golang.org/) 웹사이트에서 실행할 수 있다(필요하다면, "예제"라는 단어를 클릭해서 열어보라). 문제에 어떻게 접근해야하는지 혹은 무언가가 어떻게 구현되어 있는지 궁금하다면, 라이브러리의 문서, 코드와 예제들이 해결책, 아이디어, 배경지식을 제공해 줄 것이다.