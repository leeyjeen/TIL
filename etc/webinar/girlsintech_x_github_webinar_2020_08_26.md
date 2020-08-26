# 2020_08_26 | Girls in Tech x GitHub webinar
## 컴퓨팅 사고력이란?
*복잡한 문제*를 효율적으로 다루고 해결하는 사고 능력

### Decomposition | 문제 분해
웹 서비스에서 문제가 생겼다?
- 외부 패키지 잘못됨
- 코드 변경 사항에서 버그 발생
- ...
- ...

=> 해결 가능한 단위로 *문제를 쪼갠다*

### Pattern Recognition | 패턴 인식 
과거와 비교해서 *비슷한 문제가 있는지* 분석

### Abstraction | 추상화
*중요한 정보만 집중*, 중요하지 않은 정보는 노이즈화 시켜서 살짝 지워버린다.
문제 원인이 외부 패키지 문제인 경우 그것에만 집중.

### Algorithmic Design | 알고리즘 설계
이 문제를 해결할 수 있는 과정을 *스텝 바이 스텝으로 작성*.  
ex)
1. 문제가 발생한 컴포넌트가 사용하는 외부 패키지가 무엇인지 확인한다.
2. 외부 패키지의 공식 status page를 확인한다.
3. 공식 페이지에 확인되지 않는다면, 커뮤니티에 최근 비슷한 문제를 겪는 사람들이 있는지 확인한다.
4. 일전에 동일한 known issue가 fix된 이력이 있는지 확인한다. (regression bug 여부 확인)
5. 최근 외부 패키지에 코드 변경 사항이 있었는지 확인한다.

## Git, GitHub 소개

### Git의 필요성
- 한 줄만 바뀌어도, 점 하나만 달라져도 프로그램의 동작 여부가 달라진다
- 누가 어떤 부분을 작성하였는지
- 버전 관리

### GitHub를 이용한 형상 관리 및 협업
- 오픈소스 기여
- 변경 사항에 대한 의견 주고 받기
- 버전 관리 용이 -> 어떤 부분이 어떻게 바뀌었는지 쉽게 추적

### Git vs GitHub
- Git:
```bash
형상관리 기술의 일종!
컴퓨터 파일의 변경 사항을 추적,
여러 명의 사용자들 간에 협업하기 위한 분산 버전 관리 시스템
=> 리눅스 만든 사람(리눅스 토발즈)이 개발하였음
```

- GitHub:
```bash
분산 버전 관리 툴인 Git을 사용하는 프로젝트를 지원하는 웹 호스팅 서비스
(지금은 GitHub가 지원하는 기능이 워낙 많아져서 위의 정의가 100% 일치하진 않음)
```

### Git에서 형상관리를 위해 사용하는 개념 및 용어
버전 관리 프로그램을 만든다고 상상해보자.

#### 버전 관리 프로그램이 갖추어야 할 요소?
- 어느 단위로 *변경된 사항을 저장*해야 할지? -> Commit
- 같은 부분을 다른 사람들이 *동시에 변경*하면 어떻게 처리할 것인지? -> Conflict
- 개별적인 *버전을 관리*하고 싶다면 어떻게 처리해야 하는지? -> Branch
- 개별적인 버전을 *최종본에 합치고 싶을 땐* 어떻게 해야 하는지? -> Merge

### Git에서 사용하는 용어 정리
- Commit
- Conflict
- Branch
- Pull Request (PR)
- Review
- Merge

## create-react-app 프로젝트 GitHub로 관리!
### Branch Policy
Settings > Branches > master > Edit > Protect matching branches에서  

- merge 전 최소 몇 명에게 pull request 리뷰를 받아라.  
- merge 전 Build가 잘 되는지 확인해라.  
- 관리자 또한 이 규칙을 따르게 하라.

등의 *Branch protection rule*을 적용할 수 있다.

### Azure Board 툴 GitHub 연동하기
이슈명을 커밋메시지 입력시 Board에 Commit이 연동된다.

### Draft Pull Request
변경 내용이 많고 아직 완성된 버전은 아니지만 다른 사람들이 미리 미리 봐야할 경우 사용한다.

### CodeSpace
- Cloud VM에서 호스팅되는 개발 환경 인스턴스
- *Open in codespace* 클릭 (Beta버전으로 출시되어서 현재 메뉴에서 보이지 않는다. 미리 보고 싶으면 베타 서비스 신청하면 된다.)

#### 코드 스페이스 장점:
- 내 컴퓨터가 더러워지는 것 방지할 수 있다.
- 좀 더 빨리 리뷰 가능하다.  
- 로컬과 동일한 경험을 할 수 있다.
- 비쥬얼 스튜디오를 웹으로 동작시킨다고 생각할 수 있다.
- 디버깅도 여기서 가능하다.

예를 들어 클라우드 버츄얼 터미널에서 아래와 같이 실행 가능
```bash
> npm start 
```
테스트도 가능하다
```bash
> npm test
```


### Actions 탭
CI/CD 기능, 빌드 자동화   
```.github/workflows/buildAndDeploy.yml```

new workflow 버튼을 누르면 템플릿 목록이 있다.
원하는 템플릿을 복사 붙여넣기해서 필요한 내용만 작성하면 된다.

1) 언제 실행이 되는지 (on)
2) 환경 변수 설정 (env)
3) 무슨 일을 하는가 (jobs)
인증, 허가 관련 정보는 secret으로 처리할 수 있도록 변수로 처리한다. 이 변수 설정은 github에서 제공한다. 

Settings > Secrets > Repository secrets 에서 변수 관리


## 기타 링크
- 코드 스페이스 베타 신청: aka.ms/gitsdemo/codespace
- 타입스크립트, 루비, 고 등 제공: aka.ms/gitsdemo/learn/lab
- Microsoft Learn: aka.ms/gitsdemo/learn
- azure web app을 쓰고 싶다면: aka.ms/gitsdemo/learn/awa 




