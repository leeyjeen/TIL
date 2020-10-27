# What happens when you type google.com into your browser's address box and press enter?

## Parse URL
* `Protocol` "http"
    * 'Hyper Text Transfer Protocol'을 사용한다.
* `Resource` "/"
    * main (index) 페이지를 검색한다.

## Is it a URL or a search term?
프로토콜이나 유효한 도메인 이름이 주어지지 않은 경우 브라우저는 주소창에 제공된 텍스트를 브라우저의 기본 웹 검색 엔진에 공급한다. 대부분의 경우 URL에는 특정 브라우저의 URL 표시 줄에서 가져온 것임을 검색 엔진에 알리기 위해 특별한 텍스트가 추가된다. 

## Convert non-ASCII Unicode characters in hostname
* 브라우저는 `a-z`, `A-Z`, `0-9`, `-` 또는 `.`에 없는 문자에 대해 호스트 이름을 확인한다.

## Check HSTS list
* 브라우저는 "사전 로드된 HSTS(HTTP Strict Transport Security)" 목록을 확인한다. 이는 HTTPS를 통해서만 접속하도록 요청한 웹 사이트 목록이다.
* 웹사이트가 목록에 존재한다면, 브라우저는 요청을 HTTP 대신 HTTPS를 통해 전송한다. 그렇지 않으면 초기 요청이 HTTP를 통해 전송된다. (웹 사이트는 HSTS 목록에 없어도 HSTS 정책을 계속 사용할 수 있다. 사용자가 웹 사이트에 보내는 첫 번째 HTTP 요청은 사용자가 HTTPS 요청만 보내도록 요청하는 응답을 받게 된다. 그러나 이 단일 HTTP 요청은 잠재적으로 사용자가 다운그레이드 공격에 취약할 수 있으므로 HSTS 목록이 현대 웹 브라우저에 포함되는 것이다.)

## DNS lookup
* 브라우저는 **도메인이 캐시에 존재하는지** 확인한다. (Chrome에서 DNS 캐시를 보려면 chrome://net-internals/#dns로 이동.)
* 찾을 수 없는 경우 브라우저는 `gethostbyname` 라이브러리 함수(OS에 따라 다름)를 호출하여 조회한다.
* `gethostbyname`은 DNS를 통해 호스트 이름을 확인하기 전에 **로컬 호스트 파일**(OS에 따라 위치가 다름)에서 참조로 호스트 이름을 확인할 수 있는지 확인한다.
* `gethostbyname`에 캐시되지 않았거나 호스트 파일에서 찾을 수 없는 경우 **네트워크 스택에 구성된 DNS 서버에 요청**을 보낸다. 이는 일반적으로 **로컬 라우터** 또는 **ISP의 캐싱 DNS 서버**이다.
* DNS 서버가 동일한 서브넷에 있는 경우 네트워크 라이브러리는 **DNS 서버**에 대해 아래의 **ARP 프로세스**를 따른다.
* DNS 서버가 다른 서브넷에 있는 경우 네트워크 라이브러리는 **기본 게이트웨이 IP**에 대해 아래의 **ARP 프로세스**를 따른다.

## ARP process
ARP(Address Resolution Protocol) broadcast를 보내려면 네트워크 스택 라이브러리에서 조회할 대상 **IP 주소**가 필요하다. 또한 ARP broadcast를 보내는 데 사용할 인터페이스의 **MAC 주소**를 알아야 한다.

ARP 캐시는 우선 대상 IP에 대한 ARP 항목을 체크한다. 캐시에 있는 경우 라이브러리 함수는 다음 결과를 반환한다.: (대상 IP = MAC)

항목이 ARP 캐시에 없는 경우:

* 대상 IP 주소가 로컬 라우팅 테이블의 서브넷에 있는지 확인하기 위해 라우팅 테이블을 조회한다. 있다면 라이브러리는 해당 서브넷과 연결된 인터페이스를 사용한다. 없다면 라이브러리는 기본 게이트웨이의 서브넷이 있는 인터페이스를 사용한다. 
* 선택한 네트워크 인터페이스의 MAC 주소가 조회된다. 
* 네트워크 라이브러리는 레이어 2(OSI 모델의 데이터 링크 레이어)에 ARP 요청을 보낸다.:

`ARP Request`:
```
Sender MAC: interface:mac:address:here
Sender IP: interface.ip.goes.here
Target MAC: FF:FF:FF:FF:FF:FF (Broadcast)
Target IP: target.ip.goes.here
```

컴퓨터와 라우터 사이에 있는 하드웨어 유형에 따라..:
* Directly connected:
    * 컴퓨터가 라우터에 직접 연결된 경우 라우터는 ARP reply로 응답한다.(아래 참조)
* Hub:
    * 컴퓨터가 허브에 연결된 경우 허브는 ARP 요청을 다른 모든 포트로 브로드 캐스트한다. 라우터가 동일한 "와이어"에 연결되어 있으면 ARP reply로 응답한다.(아래 참조)
* Switch:
    * 컴퓨터가 스위치에 연결된 경우 스위치는 로컬 CAM/MAC 테이블을 확인하여 우리가 찾고 있는 MAC 주소가 있는 포트를 확인한다. 스위치에 MAC 주소에 대한 항목이 없으면 ARP 요청을 다른 모든 포트로 다시 브로드 캐스트한다.
    * 스위치의 MAC/CAM 테이블에 항목이 있으면 우리가 찾고 있는 MAC 주소가 있는 포트로 ARP 요청을 보낸다.
    * 라우터가 동일한 "와이어"에 있는 경우 ARP reply로 응답한다. (아래 참조)

`ARP Reply`:
```
Sender MAC: target:mac:address:here
Sender IP: target.ip.goes.here
Target MAC: interface:mac:address:here
Target IP: interface.ip.goes.here
```

이제 네트워크 라이브러리에 DNS 서버 또는 기본 게이트웨이의 IP 주소가 있으므로 DNS 프로세스를 재개할 수 있다.
* DNS 클라이언트는 1023 이상의 소스 포트를 사용하여 DNS 서버의 UDP 포트 53에 소켓을 설정한다.
* 응답 크기가 너무 크면 대신 TCP가 사용된다.
* 로컬/ISP DNS 서버에 기능이 없는 경우 재귀 검색이 요청되고 SOA에 도달할 때까지 DNS 서버 목록 위로 이동하며 발견되면 응답이 반환된다.

## Opening of a socket
브라우저가 대상 서버의 IP 주소를 받으면 해당 IP 주소와 URL에서 지정된 포트 번호를 가져오고(HTTP 프로토콜의 기본값은 포트 80, HTTPS는 포트 443), socket이라는 시스템 라이브러리 함수를 호출하고 TCP 소켓 스트림(`AF_INET/AF_INET6`, `SOCK_STREAM`)을 요청한다.
* 이 요청은 먼저 TCP 세그먼트가 만들어진 전송 레이어로 전달된다. 대상 포트가 헤더에 추가되고 소스 포트가 커널의 동적 포트 범위(Linux의 경우 ip_local_port_range)에서 선택된다.
* 이 세그먼트는 추가 IP 헤더를 래핑하는 네트워크 레이어로 전송된다. 대상 서버의 IP 주소와 현재 컴퓨터의 IP 주소가 삽입되어 패킷을 형성한다. 
* 다음으로 패킷은 링크 레이어에 도착한다. 머신 NIC의 MAC 주소와 게이트웨이(로컬 라우터)의 MAC 주소를 포함하는 프레임 헤더가 추가된다. 이전과 마찬가지로 커널이 게이트웨이의 MAC 주소를 모르는 경우 이를 찾기 위해 ARP 쿼리를 브로드 캐스트해야 한다.

이 시점에서 패킷은 다음 중 하나를 통해 전송할 준비가 되었다.
* Ethernet
* Wifi
* Cellular data network

대부분의 가정 또는 소규모 비즈니스 인터넷 연결의 경우 패킷은 컴퓨터에서 로컬 네트워크를 통해 전달된 다음 디지털 1과 0을 전화, 케이블 또는 무선 전화 연결을 통한 전송에 적합한 아날로그 신호로 변환하는 **모뎀**(MOdulator/DEModulator)을 통해 전달된다. 연결의 다른 쪽 끝에는 아날로그 신호를 다시 디지털 데이터로 변환하여 다음 네트워크 노드에서 처리할 수 있는 또 다른 모뎀이 있다. 여기서 시작 및 끝 주소가 추가로 분석된다.

대부분의 대기업과 일부 새로운 주거용 연결에는 광섬유 또는 직접적인 이더넷 연결이 있으며, 이 경우 데이터는 디지털로 유지되고 처리를 위해 다음 네트워크 노드로 직접 전달된다. 

결국 패킷은 로컬 서브넷을 관리하는 라우터에 도달한다. 계속해서 자율 시스템의 경계 라우터, 다른 자율 시스템, 마지막으로 대상 서버로 이동한다. 각 라우터는 IP 헤더에서 대상 주소를 추출하여 적절한 다음 홉으로 라우팅한다. IP 헤더의 TTL(Time to Live) 필드는 통과하는 각 라우터에 대해 하나씩 감소한다. TTL 필드가 0에 도달하거나 현재 라우터의 대기열에 공간이 없는 경우(네트워크 정체로 인하여) 패킷이 삭제된다.

이 송수신은 TCP 연결 흐름에 따라 여러 번 발생한다.:
* 클라이언트가 ISN(초기 시퀀스 번호)을 선택하고 ISN을 설정하고 있음을 나타내기 위해 SYN 비트가 설정된 서버로 패킷을 보낸다.
* **서버가 SYN을 수신하고 기분이 좋은 경우:**
    * 서버는 자체 초기 시퀀스 번호를 선택한다.
    * 서버는 ISN을 선택하고 있음을 나타내기 위해 SYN을 설정한다.
    * 서버는 (클라이언트 ISN + 1)을 ACK 필드에 복사하고 ACK 플래그를 추가하여 첫 번째 패킷의 수신을 승인하고 있음을 나타낸다.
* **클라이언트는 패킷을 전송하여 연결을 확인한다.**
    * 자체 시퀀스 번호를 증가시킨다.
    * 수신자 승인 번호를 증가시킨다.
    * ACK 필드를 설정한다.
* **데이터는 다음과 같이 전송된다.**
    * 한 쪽이 N 데이터 바이트를 보내면 해당 수만큼 SEQ가 증가한다.
    * 상대방이 해당 패킷(또는 패킷 문자열)의 수신을 확인하면 다른 쪽에서 마지막으로 수신한 시퀀스와 동일한 ACK 값을 가진 ACK 패킷을 보낸다.
* **연결을 종료하려면:**
    * 클로저가 FIN 패킷을 보낸다.
    * 상대방은 FIN 패킷을 ACK하고 자체 FIN을 보낸다.
    * 클로저가 ACK로 상대방의 FIN을 인식한다.

## TLS handshake
To be continued..

## HTTP protocol

## HTTP Server Request Handle

## Behind the scenes of the Browser

## Browser

## HTML parsing

## CSS interpretation

## Page Rendering

## GPU Rendering

## Window Server

## Reference
* https://github.com/alex/what-happens-when