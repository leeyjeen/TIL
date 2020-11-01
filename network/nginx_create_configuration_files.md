# Creating NGINX Plus and NGINX Configuration Files
NGINX 및 NGINX Plus는 특정 형식으로 작성된 텍스트 기반 구성 파일을 사용한다는 점에서 다른 서비스와 유사하다. 기본적으로 파일 이름은 `nginx.conf`이며 NGINX Plus의 경우 `/etc/nginx` 디렉터리에 배치된다. (NGINX 오픈소스의 경우 위치는 NGINX를 설치하는 데 사용된 패키지 시스템과 운영체제에 따라 다르다. 일반적으로 `/usr/local/nginx/conf`, `/etc/nginx` 또는 `/usr/local/etc/nginx` 중 하나이다.) 

## Directives
구성 파일은 지시문과 해당 매개변수로 구성된다. 세미콜론으로 끝나는 단순(single-line) 지시문. 다른 지시문은 관련 지시문을 그룹화하여 중괄호({})로 묶는 "컨테이너" 역할을 한다. 이를 블록이라고도 한다. 다음은 간단한 지시문의 몇 가지 예시이다.
```
user                nobody;
error_log           logs/error.log notice;
worker_processes    1;
```

## Feature-Specific Configuration Files
구성을 보다 쉽게 유지하려면 구성을 `/etc/nginx/conf.d` 디렉터리에 저장된 기능별 파일 세트로 분할하고 기본 `nginx.conf` 파일의 `include` 지시문을 사용하여 기능별 파일의 내용을 참조하는 것이 좋다. 
```
include conf.d/http;
include conf.d/stream;
include conf.d/exchange-enhanced;
```

## Contexts
컨텍스트라고 하는 몇 가지 최상위 지시문은 서로 다른 트래픽 유형에 적용되는 지시문을 함께 그룹화한다.
* `events`: 일반 연결 처리
* `http`: HTTP 트래픽
* `mail`: Mail 트래픽
* `stream`: TCP, UDP 트래픽

이러한 컨텍스트 외부에 배치된 지시문은 main 컨텍스트에 있다고 한다.

### Virtual Servers
각 트래픽 처리 컨텍스트에서 하나 이상의 server 블록을 포함하여 요청 처리를 제어하는 가상 서버를 정의한다. server 컨텍스트 내에 포함할 수 있는 지시문은 트래픽 유형에 따라 다르다.

HTTP 트래픽(`http` 컨텍스트)의 경우 각 `server` 지시문은 특정 도메인 또는 IP 주소에서 리소스에 대한 요청 처리를 제어한다. `server` 컨텍스트 내의 하나 이상의 `location` 컨텍스트는 특정 URI 집합을 처리하는 방법을 정의한다.

mail, TCP/UDP 트래픽(`mail`, `stream` 컨텍스트)의 경우 `server` 지시문은 각각 특정 TCP 포트 또는 UNIX 소켓에 도착하는 트래픽 처리를 제어한다.

### Sample Configuration File with Multiple Contexts
다음 구성은 컨텍스트 사용의 예시이다.
```
user nobody;    # a directive in the 'main' context

events {
    # configuration of connection processing
}

http {
    # configuration specific to HTTP and affecting all virtual servers
    server {
        # configuration of HTTP virtual server 1
        location /one {
            # configuration for processing URIs starting with '/one'
        }
        location /two {
            # configuration for processing URIs starting with '/two'
        }
    }

    server {
        # configuration of HTTP virtual server 2
    }
}

stream {
    # configuration specific to TCP/UDP and affecting all virtual servers
    server {
        # configuration of TCP virtual server 1
    }
}
```

### Inheritance
일반적으로 다른 컨텍스트(상위)에 포함된 하위 컨텍스트는 사위 수준에 포함된 지시문의 설정을 상속한다. 일부 지시문은 여러 컨텍스트에 나타날 수 있으며, 이 경우 지시문을 자식 컨텍스트에 포함하여 부모로부터 상속된 설정을 재정의할 수 있다. 예는 `proxy_set_header` 지시문을 참조하자.

## Reloading Configuration
구성 파일의 변경 사항을 적용하려면 reload해야 한다. nginx 프로세스를 다시 시작하거나 reload 신호를 보내 현재 요청 처리를 중단하지 않고 구성을 업그레이드할 수 있다. 자세한 내용은 Controlling NGINX Processes at Runtime을 참조하자.

NGINX Plus를 사용하면 구성을 다시 로드하지 않고도 업스트림 그룹의 서버간에 로드밸런싱을 동적으로 재구성할 수 있다. 또한 NGINX Plus API 및 key-value 저장소를 사용하여, 예를 들어 클라이언트 IP 주소를 기반으로 액세스를 동적으로 제어할 수 있다.

## Reference
* https://docs.nginx.com/nginx/admin-guide/basic-functionality/managing-configuration-files/