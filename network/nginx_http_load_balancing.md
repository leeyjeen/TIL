# Nginx - HTTP Load Balancing
NGINX 및 NGINX Plus를 로드 밸런서로 사용하는 방법에 대해 설명한다.

## Overview
여러 애플리케이션 인스턴스 간의 로드 밸런싱은 리소스 활용도를 최적화하고, 처리량을 극대화하며, 지연 시간을 줄이고, 오류 방지 구성을 보장하기 위해 일반적으로 사용되는 기술이다. 

NGINX 사용자가 대규모의 고가용성 웹 서비스를 구축하기 위해 사용하는 기술에 대해 자세히 알아 보려면 NGINX Plus for Load Balancing and Scaling Webinar를 시청하자. 

NGINX 및 NGINX Plus는 매우 효율적인 HTTP 로드 밸런서로 다양한 배포 시나리오에서 사용할 수 있다.

## Proxying HTTP Traffic to a Group of Servers
NGINX Plus 또는 NGINX 오픈소스를 사용하여 HTTP 트래픽을 서버 그룹으로 로드 밸런싱하려면 먼저 `upstream` 지시문으로 **그룹을 정의**해야 한다. 지시문은 `http` 컨텍스트에 배치된다.

그룹의 서버는 `server` 지시문을 사용하여 구성된다(NGINX에서 실행되는 가상 서버를 정의하는 서버 블록과 혼동하지 말 것). 예를 들어, 다음 구성은 **backend**라는 그룹을 정의하며 세 개의 서버 구성으로 구성된다(실제 서버가 세 개 이상일 수 있음).
```
http {
    upstream backend {
        server backend1.example.com weight=5;
        server backend2.example.com;
        server 192.0.0.1 backup;
    }
}
```

요청을 서버 그룹에 전달하기 위해 그룹 이름은 `proxy_pass` 지시문(또는 해당 프로토콜에 대한 `fastcgi_pass`, `memcached_pass`, `scgi_pass`, 또는 `uwsgi_pass` 지시문)에 지정된다. 다음 예에서 NGINX에서 실행되는 가상 서버는 모든 요청을 이전 예에서 정의한 **backend** upstream 그룹으로 전달한다.
```
server {
    location / {
        proxy_pass http://backend;
    }
}
```

다음 예는 위의 두 snippet을 결합하고 HTTP 요청을 backend 서버 그룹에 프록시하는 방법을 보여준다. 그룹은 세 개의 서버로 구성되며, 그 중 두 개는 동일한 애플리케이션의 인스턴스를 실행하고 세 번째는 백업 서버이다. upstream 블록에 로드 밸런싱 알고리즘이 지정되어 있지 않기 때문에 NGINX는 기본 알고리즘인 Round Robin을 사용한다.
```
http {
    upstream backend {
        server backend1.example.com;
        server backend2.example.com;
        server 192.0.0.1 backup;
    }

    server {
        location / {
            proxy_pass http://backend;
        }
    }
}
```

## Choosing a Load-Balancing Method
NGINX 오픈소스는 네 가지 로드 밸런싱 방법을 지원하며 NGINX Plus는 두 가지 방법을 더 추가한다.

1. **Round Robin** - 요청은 서버 가중치를 고려하여 서버 전체에 균등하게 분산된다. 이 방법이 기본적으로 사용된다(활성화하기 위한 지시문 없음).
```
upstream backend {
    # no load balancing method is specified for Round Robin
    server backend1.example.com;
    server backend2.example.com;
}
```

2. **Least Connections** - 서버 가중치를 다시 고려하여 활성 연결 수가 가장 적은 서버로 요청을 보낸다.
```
upstream backend {
    least_conn;
    server backend1.example.com;
    server backend2.example.com;
}
```

3. **IP Hash** - 요청이 전송되는 서버는 클라이언트 IP 주소에서 결정된다. 이 경우 IPv4 주소의 처음 세 옥텟 또는 전체 IPv6 주소가 해시 값을 계산하는 데 사용된다. 이 방법은 동일한 주소의 요청을 사용할 수 없는 경우가 아니면 동일한 서버에 도달하도록 보장한다.
```
upstream backend {
    ip_hash;
    server backend1.example.com;
    server backend2.example.com;
}
```

로드 밸런싱 순환에서 서버 중 하나를 일시적으로 제거해야 하는 경우 클라이언트 IP 주소의 현재 해싱을 유지하기 위해 `down` 매개변수로 표시할 수 있다. 이 서버에서 처리할 요청은 그룹의 다음 서버로 자동으로 전송된다.
```
upstream backend {
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com down;
}
```

4. **Generic Hash** - 요청이 전송되는 서버는 텍스트 문자열, 변수 또는 조합이 될 수 있는 사용자 정의 키에서 결정된다. 예를 들어 키는 쌍을 이루는 소스 IP 주소와 포트이거나 다음 예와 같이 URI일 수 있다.
```
upstream backend {
    hash $request_uri consistent;
    server backend1.example.com;
    server backend2.example.com;
}
```

`hash` 지시문에 대한 옵션 `consistent` 매개변수는 `ketama` consistent-hash 로드 밸런싱을 활성화한다. 요청은 사용자 정의 해시 키 값을 기반으로 모든 upstream 서버에 균등하게 분산된다. upstream 서버가 upstream 그룹에 추가되거나 제거되면 로드 밸런싱 캐시 서버 또는 상태를 누적하는 기타 응용 프로그램의 경우 캐시 누락을 최소화하는 몇 개의 키만 다시 매핑된다.

5. **Least Time** (NGINX Plus only) - 각 요청에 대해 NGINX Plus는 가장 낮은 평균 지연 시간과 가장 낮은 활성 연결 수를 가진 서버를 선택한다. 여기서 가장 낮은 평균 지연 시간은 `least_time` 지시문에 포함된 다음 매개변수를 기준으로 계산된다.
* `header` - 서버로부터 첫 바이트를 수신하는 데 걸리는 시간
* `last_byte` - 서버로부터 전체 응답을 수신하는 데 걸리는 시간
* `last_byte inflight` - 불완전한 요청을 고려하여 서버로부터 전체 응답을 수신하는 데 걸리는 시간
```
upstream backend {
    least_time header;
    server backend1.example.com;
    server backend2.example.com;
}
```

6. **Random** - 각 요청이 랜덤으로 선택된 서버로 전달된다. `two` 매개변수가 지정되면 먼저 NGINX는 서버 가중치를 고려하여 두 서버를 무작위로 선택한 다음 지정된 방법을 사용하여 다음 서버 중 하나를 선택한다.
* `least_conn` - 최소 활성 연결 수
* `least_time=header` (NGINX Plus) - 서버에서 응답 헤더를 수신하는 데 걸리는 최소 평균 시간 (`$upstream_header_time`)
* `least_time=last_byte` (NGINX Plus) - 서버에서 전체 응답을 수신하는 데 걸리는 최소 평균 시간
```
upstream backend {
    random two least_time=last_byte;
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com;
    server backend4.example.com;
}
```

**Random** 로드 밸런싱 방법은 여러 로드 밸런서가 동일한 backend 집합에 요청을 전달하는 분산 환경에 사용해야 한다. 로드 밸런서가 모든 요청을 전체적으로 볼 수 있는 환경의 경우 라운드 로빈, 최소 연결 및 최소 시간과 같은 다른 로드 밸런싱 방법을 사용하자.

## Server Weights
기본적으로 NGINX는 Round Robin 방법을 사용하여 가중치에 따라 그룹의 서버간에 요청을 분산한다. `server` 지시문에 대한 `weight` 매개변수는 서버의 가중치를 설정한다. 기본값은 1이다.
```
upstream backend {
    server backend1.example.com weight=5;
    server backend2.example.com;
    server 192.0.0.1 backup;
}
```
이 예시에서 **backend1.example.com**의 가중치는 `5`이다. 다른 두 서버는 기본 가중치 `1`을 갖지만 IP 주소가 `192.0.0.1`인 서버는 `backup` 서버로 표시되며 다른 두 서버를 모두 사용할 수 없는 경우가 아니면 요청을 수신하지 않는다. 이 가중치 구성을 사용하면 6개의 요청 중 5개는 **backend1.example.com**으로, 1개는 **backend2.example.com**으로 전송된다.

## Server Slow-Start
서버 느린 시작 기능은 최근에 복구된 서버가 연결로 인해 과부하되는 것을 방지하며 이로 인해 시간이 초과되어 서버가 다시 실패한 것으로 표시될 수 있다.

NGINX Plus에서 느린 시작을 사용하면 upstream 서버가 복구되거나 사용 가능해진 후 가중치를 `0`에서 공칭 값으로 점진적으로 복구할 수 있다. 이는 `server` 지시문에 대한 `slow_start` 매개변수로 수행할 수 있다.
```
upstream backend {
    server backend1.example.com slow_start=30s;
    server backend2.example.com;
    server 192.0.0.1 backup;
}
```
시간 값(여기서는 30초)은 NGINX Plus가 서버에 대한 연결 수를 전체 값으로 늘리는 시간을 설정한다.

그룹에 단일 서버만 있는 경우 `server` 지시문에 대한 `max_fails`, `fail_timeout`, `slow_start` 매개변수가 무시되고 서버는 절대 사용할 수 없는 것으로 간주되지 않는다는 점에 유의하라.

## Configuring HTTP Load Balancing Using DNS
서버 그룹의 구성은 DNS를 사용하여 런타임에 수정할 수 있다.

`server` 지시문에서 도메인 이름으로 식별되는 upstream 그룹의 서버의 경우 NGINX Plus는 해당 DNS 레코드의 IP 주소 목록에 대한 변경 사항을 감시할 수 있으며, 변경 사항을 재시작할 필요 없이 upstream 그룹의 로드 밸런싱에 자동으로 적용할 수 있다. 이는 `server` 지시문에 대한 `resolve` 매개변수와 함께 http 블록에 `resolver` 지시문을 포함함으로써 이루어질 수 있다.
```
http {
    resolver 10.0.0.1 valid=300s ipv6=off;
    resolver_timeout 10s;
    server {
        location / {
            proxy_pass http://backend;
        }
    }
    upstream backend {
        zone backend 32k;
        least_conn;
        # ...
        server backend1.example.com resolve;
        server backend2.example.com resolve;
    }
}
```
이 예에서 `server` 지시문에 대한 `resolve` 매개변수는 NGINX Plus에 주기적으로 **backend1.example.com** 및 **backend2.example.com** 도메인 이름을 IP 주소로 다시 확인하도록 지시한다.

`resolver` 지시문은 NGINX Plus가 요청을 보내는 DNS 서버의 IP 주소를 정의한다(여기서는 10.0.0.1). 기본적으로 NGINX Plus는 레코드의 TTL(Time-to-Live)에 지정된 빈도로 DNS 레코드를 다시 확인하지만 유효한 매개변수를 사용하여 TTL 값을 재정의할 수 있다. 예에서는 300초 또는 5분이다.

선택적 `ipv6=off` 매개변수는 기본적으로 IPv4 및 IPv6 주소 확인이 모두 지원되지만 로드 밸런싱에 IPv4 주소만 사용됨을 의미한다.

도메인 이름이 여러 IP 주소로 확인되면 주소가 upstream 구성에 저장되고 로드 밸런싱이 된다. 이 예에서 서버는 최소 연결 로드 밸런싱 방법에 따라 로드 밸런싱된다. 서버의 IP 주소 목록이 변경된 경우 NGINX Plus는 즉시 새 주소 집합에 대한 로드 밸런싱을 시작한다.

## Reference
* https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/