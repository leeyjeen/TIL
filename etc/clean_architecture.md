# Clean Architecture
<img src="./images/CleanArchitecture.jpg" width="80%" alt="Insert an element at the end of the queue">
<br><br>

지난 몇 년간 시스템 아키텍처에 관한 모든 범위의 아이디어를 봐왔다. 이는 다음을 포함한다.
* Hexagonal Architecture
* Onion Architecture
* Screaming Architecture
* DCI
* BCE

이러한 아키텍처들은 모두 세부사항에서 다소 차이가 있지만, 매우 유사하다. 그들은 모두 동일한 목표, 즉 관심사 분리를 가지고 있다. 그들은 모두 소프트웨어를 레이어(계층)로 나누어 이러한 분리를 달성한다. 각각에는 비즈니스 규칙을 위한 계층과 인터페이스에 대한 계층을 하나 이상 가지고 있다.

이들 아키텍처 각각은 다음과 같은 시스템을 생성한다.

1. 프레임워크와 독립적이다. 아키텍처는 일부 기능이 포함된 소프트웨어 라이브러리의 존재에 의존하지 않는다. 이를 통해 제한된 제약에 시스템을 밀어 넣을 필요 없이 이러한 프레임워크를 도구로 사용할 수 있다.
2. 테스트가 가능하다. 비즈니스 규칙은 UI, 데이터베이스, 웹 서버 또는 기타 외부 요소 없이 테스트할 수 있다. 
3. UI와 독립적이다. UI는 시스템의 나머지 부분을 변경하지 않고도 쉽게 변경할 수 있다. 예를 들어 비즈니스 규칙을 변경하지 않고 웹 UI를 콘솔 UI로 바꿀 수 있다.
4. 데이터베이스 독립적이다. Oracle 또는 SQL Server, MongoDB, CouchDB 등 다른 것으로 교체할 수 있다. 비즈니스 규칙은 데이터베이스에 바인딩되지 않는다.
5. 외부 기관과 독립적이다. 실제로 비즈니스 규칙은 외부 세계에 대해 전혀 알지 못한다.

맨 위에 있는 다이어그램은 이러한 모든 아키텍처를 하나의 실행 가능한 아이디어로 통합하려는 시도이다.

# The Dependency Rule | 종속성 규칙
동심원은 소프트웨어의 여러 영역을 나타낸다. 일반적으로 내부로 갈수록 고수준의 소프트웨어이다. 바깥 쪽 원은 메커니즘이다. 내부 원은 정책이다.

이 아키텍처가 작동하도록 만드는 우선 규칙은 Dependency Rule(종속성 규칙)이다. 이 규칙은 소스 코드 종속성이 내부로만 향할 수 있다고 말한다. 내부 원의 어떤 것도 외부 원의 어떤 것에 대해 전혀 알 수 없다. 특히 외부 원에 선언된 이름은 내부 원에 있는 코드에서 언급되지 않아야 한다. 여기에는 함수, 클래스, 변수 또는 기타 명명된 소프트웨어 개체가 포함된다.

마찬가지로, 외부 원에서 사용되는 데이터 형식은 특히 외부 원의 프레임워크에 의해 생성되는 경우 내부 원에서 사용해서는 안된다. 바깥 쪽 원 안에 있는 어떤 것도 안쪽 원에 영향을 주는 것을 원하지 않는다.

# Entities

# Use Cases

# Interface Adapters

# Frameworks and Drivers

# Only Four Circles?

# Crossing boundaries

# What data crosses the boundaries

# Conclusion

## Reference
* https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html