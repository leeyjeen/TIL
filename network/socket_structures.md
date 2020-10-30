# Socket Structures
Unix 소켓 프로그래밍에서는 주소, 포트 및 기타 정보에 대한 정보를 보유하기 위해 다양한 구조가 사용된다. 대부분의 소켓 함수에는 소켓 주소 구조에 대한 포인터가 인수로 필요하다. 이 페이지에서 정의된 구조는 Internet Protocol Family와 관련이 있다. 

## sockaddr
첫 번째 구조는 소켓 정보를 보유하는 `sockaddr`이다.
```
struct sockaddr {
    unsigned short sa_family;
    char    sa_data[14];
}
```
이는 일반적인 소켓 주소 구조로, 대부분의 소켓 함수 호출에서 전달된다.

다음 테이블은 멤버 필드에 대한 설명을 제공한다.

Attribute | Values | Description
------- | ------ | ------
sa_family | AF_INET, AF_UNIX, AF_NS, AF_IMPLINK | address family를 나타낸다. 대부분의 인터넷 기반 응용 프로그램에서 AF_INET을 사용한다.
sa_data | Protocol-specific Address	| 14 바이트 프로토콜 특정 주소의 내용은 주소 유형에 따라 해석된다. 인터넷 제품군의 경우 아래 정의된 *sockaddr_in* 구조로 표시되는 포트 번호 IP 주소를 사용한다.

## sockaddr in
소켓의 요소를 참조하는 데 도움이 되는 두 번째 구조는 다음과 같다.
```
struct sockaddr_in {
    short int           sin_family;
    unsigned short int  sin_port;
    struct in_addr      sin_addr;
    unsigned char       sin_zero[8];
}
```
멤버 필드에 대한 설명은 다음과 같다.

Attribute | Values | Description
------- | ------ | ------
sin_family | AF_INET, AF_UNIX, AF_NS, AF_IMPLINK | address family를 나타낸다. 대부분의 인터넷 기반 응용 프로그램에서 AF_INET을 사용한다.
sin_port | Service Port | 네트워크 바이트 순서의 16 비트 포트 번호
sin_addr | IP Address | 네트워크 바이트 순서의 32 비트 IP 주소
sin_zero | Not Used | 사용되지 않으므로 이 값을 Null로 설정한다.

## in addr
이 구조는 위의 구조에서만 구조 필드로 사용되며 32 비트 netid/hostid를 갖는다.
```
struct in_addr {
    unsigned long s_addr;
}
```

Attribute | Values | Description
------- | ------ | ------
s_addr | IP Address | 네트워크 바이트 순서의 32 비트 IP 주소

## hostent
이 구조는 호스트와 관련된 정보를 유지하는 데 사용된다.
```
struct hostent {
    char *h_name;
    char **h_aliases;
    int h_addrtype;
    int h_length;
    char **h_addr_list

#define h_addr h_addr_list[0]
};
```

Attribute | Values | Description
------- | ------ | ------
h_name | ti.com 등.. | 호스트의 공식 이름. 예시: google.com
h_aliases | TI | 호스트 이름 별칭 리스트
h_addrtype | AF_INET | address family를 포함하여 인터넷 기반 응용 프로그램의 경우 항상 AF_INET이다.
h_length | 4 | IP 주소의 길이를 담는다. 인터넷 주소의 길이는 4이다.
h_addr_list | in_addr | 인터넷 주소의 경우 포인터 배열 h_addr_list[0], h_addr_list[1] 등은 in_addr 구조를 가리킨다.

## servent
이 특정 구조는 서비스 및 관련 포트와 관련된 정보를 유지하는 데 사용된다.

```
struct servent {
    char *s_name;
    char **s_aliases;
    int s_port;
    char *s_proto;
}
```

Attribute | Values | Description
------- | ------ | ------
s_name | http | 서비스의 공식 이름이다. 예시: SMTP, FTP POP3 등
s_aliases | ALIAS | 서비스 별칭 리스트이다. 대부분의 경우 NULL로 설정된다.
s_port | 80 | 연결된 포트 번호가 있다. 예를 들어 HTTP의 경우 80이다.
s_proto | TCP, UDP | 사용된 프로토콜이 설정된다. 인터넷 서비스는 TCP 또는 UDP를 사용하여 제공된다.

## Reference
* https://www.tutorialspoint.com/unix_sockets/socket_structures.htm