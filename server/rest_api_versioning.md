# REST API Versioning
복잡성 관리를 위하여 API 버전을 관리하는 것이 좋다. 버전 관리를 사용하면 필요한 변경 사항이 식별될 때 더 빠르게 반복할 수 있다.

> 시스템에 대한 지식과 경험이 향상됨에 따라 API 변경은 불가피해진다. 이러한 변경의 영향을 관리하는 것은 기존 클라이언트 통합을 중단할 위험이 있는 경우 상당히 어려울 수 있다.

## When to Version
API는 주요 변경 사항이 있을 때만 버전을 높이면 된다. 주요 변경 사항은 다음과 같다.
* 하나 이상의 호출에 대한 응답 데이터 형식의 변경
* 요청 또는 응답 유형의 변경 (예: 정수를 부동 소수점으로 변경)
* API의 일부 제거

## How to Version
REST는 특정 버전 관리 가이드라인을 제공하지 않지만 가장 일반적으로 사용되는 접근 방식은 다음 세 가지 범주로 분류된다.

### URI Versioning
URI를 사용하는 것은 URI가 고유한 리소스를 참조해야 한다는 원칙을 위반하지만 가장 간단한 방법이며 가장 일반적으로 사용된다. 또한 버전이 업데이트되면 클라이언트 통합을 중단할 수 있다.

예:
```
http://api.example.com/v1
http://apiv1.example.com
```

버전은 숫자이거나 "v[x]" 구문을 사용하여 지정할 필요가 없다. 대안으로는 API를 만드는 팀에게 충분히 의미 있고 버전이 변경됨에 따라 변경될 수 있을만큼 충분히 유연한 날짜, 프로젝트 이름, 시즌 또는 기타 식별자가 있다.

### Versioning using Custom Request Header
사용자 정의 헤더(예: Accept-version)를 사용하면 기존 Accept 헤더에 의해 구현된 콘텐츠 협상 동작을 효과적으로 복제하더라도 버전 간 URI를 유지할 수 있다.

예:
```
Accept-version: v1
Accept-version: v2
```

### Versioning using Accept header
콘텐츠 협상을 사용하면 깨끗한 URL 집합을 유지할 수 있지만 어딘가에 다른 버전의 콘텐츠를 제공하는 복잡성을 처리해야한다. 이 부담은 보낼 리소스 버전을 파악하는 책임이 있는 API 컨트롤러로 스택 위로 이동하는 경향이 있다. 클라이언트가 리소스를 요청하기 전에 지정할 헤더를 알아야 하므로 최종 결과는 더 복잡한 API가 되는 경향이 있다.

예:
```
Accept: application/vnd.example.v1+json
Accept: application/vnd.example+json;version=1.0
```

현실 세계에서 API는 완전히 안정적일 수 없다. 따라서 이 변경 사항을 관리하는 방법이 중요하다. API를 잘 문서화하고 점진적인 API 지원 중단은 대부분의 API에서 허용 가능한 관행일 수 있다.

## Reference
* https://restfulapi.net/versioning/