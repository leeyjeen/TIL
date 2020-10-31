# Nginx - Beginner's Guide
이 가이드는 nginx에 대한 기본적인 소개를 제공하며 nginx로 수행할 수 있는 몇 가지 간단한 작업에 대해 설명한다. nginx가 이미 독자의 컴퓨터에 설치되어 있다고 가정한다. 그렇지 않은 경우 nginx 설치 페이지를 참조하자. 이 가이드에서는 nginx를 시작 및 중지하는 방법, 구성을 다시 로드하는 방법, 구성 파일의 구조를 설명하고, 정적 콘텐츠를 제공하도록 nginx를 설정하는 방법, 프록시 서버로 nginx를 구성하는 방법 및 FastCGI 응용 프로그램과 연결하는 방법을 설명한다. 

nginx에는 하나의 master 프로세스와 여러 worker 프로세스가 있다. master 프로세스의 주요 목적은 구성을 읽고 평가하고 worker 프로세스를 유지하는 것이다. worker 프로세스는 요청의 실제 처리를 수행한다. nginx는 이벤트 기반 모델 및 OS 종속 메커니즘을 사용하여 worker 프로세스 간에 요청을 효율적으로 분산한다. worker 프로세스 수는 구성 파일에 정의되어 있으며 지정된 구성에 대해 고정되거나 사용 가능한 CPU 코어 수에 맞게 자동으로 조정될 수 있다(worker_processes 참조).

nginx와 그 모듈의 작동 방식은 구성 파일에서 결정된다. 기본적으로 구성 파일의 이름은 `nginx.conf`이며 `/usr/local/nginx/conf`, `/etc/nginx` 또는 `/usr/local/etc/nginx` 디렉터리에 배치된다.

## Starting, Stopping, and Reloading Configuration
nginx를 시작하기 위해 실행 파일을 실행한다. nginx가 시작되면 `-s` 매개변수로 실행 파일을 호출하여 제어할 수 있다. 다음 구문을 사용하자:
```
nginx -s signal
```
`signal`은 다음 중 하나일 수 있다.
* `stop`: 빠른 종료
* `quit`: 우아한 종료
* `reload`: 구성 파일을 다시 로드
* `reopen`: 로그 파일 다시 열기

예를 들어 worker 프로세스가 현재 요청 처리를 완료할 때까지 대기하면서 nginx 프로세스를 중지하려면 다음 명령을 실행할 수 있다.
```
nginx -s quit
```
> 이 명령은 nginx를 시작한 동일한 사용자로 실행해야 한다.

구성 파일에서 변경한 사항은 reload 구성 명령이 nginx로 전송되거나 nginx가 재시작될 때까지 적용되지 않는다. 구성을 다시 로드하려면 다음을 실행한다.
```
nginx -s reload
```

master 프로세스는 구성을 다시 로드하라는 신호를 받으면 새 구성 파일의 구문 유효성을 확인하고 제공된 구성을 적용하려고 한다. 성공하면 master 프로세스는 새 worker 프로세스를 시작하고 이전 worker 프로세스에 메시지를 보내 종료하도록 요청한다. 그렇지 않으면 master 프로세스가 변경 사항을 롤백하고 이전 구성으로 계속 작업한다. 이전 worker 프로세스는 종료 명령을 받고 새 연결 수락을 중지하고 이러한 모든 요청이 처리될 때까지 현재 요청을 계속 처리한다. 그 후 이전 worker 프로세스가 종료된다.

`kill` 유틸리티와 같은 Unix 도구를 사용하여 nginx 프로세스에 신호를 보낼 수도 있다. 이 경우 신호는 주어진 프로세스 ID를 가진 프로세스로 직접 전송된다. nginx master 프로세스의 프로세스 ID는 기본적으로 `/usr/local/nginx/logs` 또는 `/var/run` 디렉터리의 `nginx.pid`에 기록된다. 예를 들어 master 프로세스 IDrk 1628인 경우 QUIT 신호를 보내 nginx의 정상 종료를 발생 시키려면 다음을 실행한다.
```
kill -s QUIT 1628
```

실행중인 모든 nginx 프로세스의 목록을 가져오려면 `ps` 유틸리티를 다음과 같이 사용할 수 있다.
```
ps -ax | grep nginx
```

## Configuration File's Structure
nginx는 구성 파일에 지정된 지시문에 의해 제어되는 모듈로 구성된다. 지시문은 단순 지시문과 블록 지시문으로 나뉜다. 단순 지시문은 공백으로 구분된 이름과 매개변수로 구성되며 세미콜론(;)으로 끝난다. 블록 지시문은 단순 지시문과 동일한 구조를 갖지만 세미콜론 대신 중괄호({), (})로 묶인 추가 명령 세트로 끝난다. 블록 지시문이 중괄호 안에 다른 지시문을 포함할 수 있는 경우 이를 컨텍스트(예: event, http, server 및 location)라고 한다.

컨텍스트 외부의 구성 파일에 배치된 지시문은 기본 컨텍스트에 있는 것으로 간주된다. event 및 http 지시어는 main 컨텍스트에, server는 http, location은 server 내에 있다.

'#' 기호 뒤의 나머지 라인은 주석으로 간주된다.

## Serving Static Content
중요한 웹 서버 작업은 파일(예: 이미지 또는 정적 HTML 페이지)을 제공하는 것이다. 요청에 따라 파일이 다른 로컬 디렉터리인 `/data/www`(HTML 파일을 포함할 수 있음)와 `/data/images`(이미지 포함)에서 제공되는 예를 구현해보자. 이를 위해 구성 파일을 편집하고 두 개의 location 블록이 있는 http 블록 내에 server 블록을 설정해야 한다.

먼저 `/data/www` 디렉터리를 생성하고 텍스트 내용이 있는 `index.html` 파일을 여기에 넣고 `/data/images` 디렉터리를 생성하고 그 안에 이미지를 배치한다.

다음으로 구성 파일을 연다. 기본 구성 파일에는 이미 대부분 주석 처리된 server 블록의 몇 가지 예가 포함되어 있다. 지금은 이러한 모든 블록을 주석 처리하고 새 server 블록을 시작한다.

```
http {
    server {
    }
}
```

일반적으로 구성 파일에는 수신 대기하는 포트와 서버 이름으로 구분되는 여러 server 블록이 포함될 수 있다. nginx가 요청을 처리하는 서버를 결정하면 server 블록 내에 정의된 location 지시어의 매개변수에 대해 요청의 헤더에 지정된 URI를 테스트한다.

server 블록에 다음 location 블록을 추가한다.

```
location / {
    root /data/www;
}
```

이 location 블록은 요청의 URI와 비교하여 "/" 접두사를 지정한다. 일치하는 요청의 경우 root 지시어에 지정된 경로, 즉 `/data/www`에 URI가 추가되어 로컬 파일 시스템에서 요청된 파일의 경로를 형성한다. 일치하는 location 블록이 여러 개인 경우 nginx는 접두사가 가장 긴 블록을 선택한다. 위의 location 블록은 길이가 1인 가장 짧은 접두사를 제공하므로 다른 모든 location 블록이 일치 항목을 제공하지 못하는 경우에만 이 블록이 사용된다.

다음으로 두 번째 location 블록을 추가한다.

```
location /images/ {
    root /data;
}
```

`/images`로 시작하는 요청과 매치된다. (location /도 이러한 요청과 매치되지만 접두어가 더 짧다.)

server 블록의 결과 구성은 다음과 같아야 한다.

```
server {
    location / {
        root /data/www;
    }

    location /images/ {
        root /data;
    }
}
```

이는 이미 표준 포트 80에서 수신하는 서버의 작동 구성이며 `http://localhost/`의 로컬 시스템에서 액세스할 수 있다. `/images`로 시작하는 URI가 있는 요청에 대한 응답으로 서버는 `/data/images` 디렉터리에서 파일을 보낸다. 예를 들어, `http://localhost/images/example.png` 요청에 대한 응답으로 nginx는 `/data/images/example.png` 파일을 보낸다. 해당 파일이 존재하지 않는다면 nginx는 404 에러를 나타내는 응답을 보낸다. URI가 `/images/`로 시작하지 않는 요청은 `/data/www` 디렉터리에 매핑된다. 예를 들어, `http://localhost/some/example.html` 요청에 대한 응답으로 nginx는 `/data/www/some/example.html` 파일을 보낸다.

새 구성을 적용하려면 nginx가 아직 시작되지 않은 경우 시작하거나 다음을 실행하여 reload 신호를 nginx의 master 프로세스로 보낸다.
```
nginx -s reload
```

> 예상대로 작동하지 않는 경우 `/usr/local/nginx/logs` 또는 `/var/log/nginx` 디렉터리의 `access.log` 및 `error.log` 파일에서 원인을 찾아볼 수 있다.

## Setting Up a Simple Proxy Server
nginx의 잦은 사용 중 하나는 프록시 서버로 설정하는 것인데, 이는 요청을 수신하여 프록시 서버로 전달하고, 그 서버에서 응답을 검색하여 클라이언트로 전성하는 서버를 의미한다. 

로컬 디렉터리의 파일로 이미지 요청을 제공하고 다른 모든 요청을 프록시 서버로 전송하는 기본 프록시 서버를 구성할 것이다. 이 예에서 두 서버는 단일 nginx 인스턴스에서 정의된다.

먼저 다음 내용으로 nginx의 구성 파일에 server 블록을 하나 더 추가하여 프록시 서버를 정의한다.

```
server {
    listen 8080;
    root /data/up1;

    location / {
    }
}
```

이는 포트 8080에서 수신 대기하고(이전에는 표준 포트 80이 사용되었기 때문에 listen 지시어가 지정되지 않았음) 모든 요청을 로컬 파일 시스템의 `/data/up1` 디렉터리에 매핑하는 간단한 서버이다. 이 디렉터리를 만들고 `index.html` 파일을 디렉터리에 넣자. root 지시어는 서버 컨텍스트에 배치된다는 점에 유의하자. 이러한 root 지시어는 요청을 처리하기 위해 선택한 location 블록이 자체 root 지시어를 포함하지 않을 때 사용된다.

다음으로 이전 섹션의 서버 구성을 사용하고 수정하여 프록시 서버 구성으로 만든다. 첫 번째 location 블록의 매개변수에 지정된 프록시 서버의 프로토콜, 이름 및 포트와 함께 `proxy_pass` 지시어를 입력한다(이 경우에는 `http://localhost:8080`).

```
server {
    location / {
        proxy_pass http://localhost:8080;
    }

    location /images/ {
        root /data;
    }
}
```

현재 `/images/` 접두사가 있는 요청을 `/data/images` 디렉터리 아래의 파일에 매핑하는 두 번째 location 블록을 수정하여 일반적인 파일 확장자를 가진 이미지 요청과 일치하도록 한다. 수정된 location 블록은 다음과 같다.

```
location ~ \.(git|jpg|png)$ {
    root /data/images;
}
```

매개변수는 `.gif`, `.jpg`, `.png`로 끝나는 모든 URI와 일치하는 정규 표현식이다. 정규식 앞에 `~`가 와야한다. 해당 요청은 `/data/images` 디렉터리에 매핑된다.

nginx가 요청을 처리할 location 블록을 선택할 때 먼저 접두사를 지정하는 location 지시어를 확인하고 가장 긴 접두사가 있는 location을 기억한 다음 정규식을 확인한다. 정규 표현식과 일치하는 항목이 있으면 nginx는 이 location을 선택하고 그렇지 않으면 이전에 기억된 location을 선택한다.

프록시 서버의 결과 구성은 다음과 같다.

```
server {
    location / {
        proxy_pass http://localhost:8080/;
    }

    location ~ \.(gif|jpg|png)$ {
        root /data/images;
    }
}
```

이 서버는 `.gif`, `.jpg`, `.png`로 끝나는 요청을 필터링하고 root 지시어의 매개변수에 URI를 추가하여 `/data/images` 디렉터리에 매핑하고 다른 모든 요청을 위에 구성된 프록시 서버로 전달한다.

새 구성을 적용하려면 이전 섹션에서 설명한대로 reload 신호를 nginx에 전송한다.

프록시 연결을 추가로 구성하는 데 사용할 수 있는 더 많은 지시어가 있다.

## Setting Up FastCGI Proxying
nginx는 PHP와 같은 다양한 프레임워크 및 프로그래밍 언어로 빌드된 애플리케이션을 실행하는 FastGCI 서버로 요청을 라우팅하는 데 사용할 수 있다.

FastGCI 서버에서 작동하는 가장 기본적인 nginx 구성에는 proxy_pass 지시어 대신 fastcgi_pass 지시어를 사용하고 FastGCI 서버에 전달된 매개변수를 설정하는 fastcgi_param 지시어를 사용하는 것이 포함된다. FastGCI 서버가 `localhost:9000`에서 액세스 가능하다고 가정하자. 이전 섹션의 프록시 구성을 기반으로 하여 proxy_pass 지시어를 fastgci_pass 지시어로 바꾸고 매개변수를 `localhost:9000`으로 변경한다. PHP에서 SCRIPT_FILENAME 매개변수는 스크립트 이름을 결정하는 데 사용되며 QUERY_STRING 매개변수는 요청 매개변수를 전달하는 데 사용된다. 결과 구성은 다음과 같다.

```
server {
    location / {
        fastcgi_pass    localhost:9000;
        fastcgi_param   SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param   QUERY_STRING    $query_string;
    }

    location ~ \.(gif|jpg|png)$ {
        root /data/images;
    }
}
```

이렇게 하면 정적 이미지에 대한 요청을 제외한 모든 요청을 FastCGI 프로토콜을 통해 `localhost:9000`에서 작동하는 프록시 서버로 라우팅하는 서버가 설정된다.

## Reference
* http://nginx.org/en/docs/beginners_guide.html