# Node.js Introduction

## Intro
**Node.js**는 구글 크롬의 JavaScript V8 엔진을 기반으로 구축된 매우 강력한 **JavaScript 기반** 서버 사이드 플랫폼이다. 동영상 스트리밍 사이트, 단일 페이지 애플리케이션 및 기타 웹 애플리케이션과 같은 I/O 집약적인 웹 애플리케이션을 개발하는 데 사용된다. 오픈 소스로서 완전히 무료로 제공되며 전 세계 수천 명의 개발자들이 사용하고 있다.

## What is Node.js?
Node.js 공식 문서에서 설명하는 Node.js는 아래와 같다.
> Node.js는 크롬의 자바스크립트 런타임에 구축된 플랫폼으로 빠르고 확장 가능한 네트워크 애플리케이션을 쉽게 구축할 수 있다. Node.js는 event-driven, non-blocking I/O 모델을 사용하여 가볍고 효율적이며 분산된 장치에서 실행되는 데이터 집약적인 실시간 애플리케이션에 적합하다.

Node.js는 서버 사이드 및 네트워킹 애플리케이션을 개발하기 위한 오픈 소스 크로스 플랫폼 런타임 환경이다. Node.js 애플리케이션은 JavaScript로 작성되며 OS X, Microsoft Windows 및 Linux에서 Node.js 런타임 내에서 실행될 수 있다.

또한 Node.js는 Node.js를 사용하는 웹 애플리케이션의 개발을 매우 단순화하는 다양한 JavaScript 모듈의 풍부한 라이브러리를 제공한다.

```
Node.js = Runtime Environment + JavaScript Library
```

## Features of Node.js
다음은 소프트웨어 설계시 Node.js를 선택하도록 만드는 중요한 특징들이다.

* **비동기식(Asynchronous) & 이벤트 기반(Event driven)**:
    * Node.js 라이브러리의 모든 API는 `비동기`, 즉 `논블로킹`식이다. 이는 기본적으로 Node.js 기반 서버가 데이터를 반환할 때까지 기다리지 않는다는 것을 의미한다. 서버는 호출 후 다음 API로 이동하며, Node.js의 이벤트 알림 메커니즘은 서버가 이전 API 호출로부터 응답을 받을 수 있도록 도와준다.

* 매우 **빠른** 속도:
    * Google Chrome의 V8 JavaScript 엔진에 구축된 Node.js 라이브러리는 코드 실행 속도가 매우 빠르다.

* **단일 스레드**이지만 **뛰어난 확장성**:
    * Node.js는 이벤트 루핑이 있는 단일 스레드 모델을 사용한다. `이벤트 메커니즘`은 서버가 `논블로킹` 방식으로 응답할 수 있도록 돋고, 요청을 처리하기 위해 제한된 스레드를 만드는 기존의 서버와 달리 서버의 확장성이 매우 높게 만든다.
    * Node.js는 단일 스레드 프로그램을 사용하며 Apache HTTP서버와 같은 기존 서버보다 `훨씬 많은 수의 요청에 서비스를 제공`할 수 있다.

* **버퍼링 없음**:
    * Node.js 애플리케이션은 어떤 데이터도 버퍼링하지 않는다. 단순히 `데이터를 청크로 출력`한다.

* **라이센스**:
    * Node.js는 MIT 라이센스에 따라 릴리즈된다.

## Where to Use Node.js?
Node.js가 완벽한 기술 파트너임을 입증하고 있는 분야는 다음과 같다.

* I/O 바운드 애플리케이션
* 데이터 스트리밍 애플리케이션
* 데이터 집약적 실시간 애플리케이션 (DIRT)
* JSON API 기반 애플리케이션
* 단일 페이지 애플리케이션

## Where Not to Use Node.js?
CPU 집약적 애플리케이션에는 Node.js를 사용하는 것이 권장되지 않는다.

## Reference
* https://www.tutorialspoint.com/nodejs/nodejs_introduction.htm