# Socket - Ports and Services
클라이언트 프로세스가 서버에 연결하려는 경우 클라이언트는 연결하려는 서버를 식별할 방법이 있어야 한다. 클라이언트가 서버가 있는 호스트의 32 비트 인터넷 주소를 알고 있으면 해당 호스트에 연결할 수 있다. 그러나 클라이언트는 해당 호스트에서 실행중인 특정 서버 프로세스를 어떻게 식별할까?

호스트에서 실행되는 특정 서버 프로세스를 식별하는 문제를 해결하기 위해 TCP와 UDP 모두 잘 알려진 포트 그룹을 정의했다.

우리의 목적을 위해 포트는 1024에서 65535 사이의 정수로 정의된다. 1024보다 작은 모든 포트 번호는 잘 알려진 것으로 간주되기 때문이다. 예를 들어, 텔넷은 포트 23을 사용하고, http는 80을 사용하고, ftp는 21을 사용한다.

네트워크 서비스에 대한 포트 할당은 /etc/services 파일에서 찾을 수 있다. 자체 서버를 작성하는 경우 서버에 포트를 할당할 때 주의해야 한다. 해당 포트가 다른 서버에 할당되지 않도록 해야한다.

일반적으로 5000 이상의 포트 번호를 할당하는 것이 관례이다. 그러나 포트 번호가 5000 이상인 서버를 작성한 조직이 많이 있다. 예를 들어 Yahoo Messenger는 5050에서 실행되고 SIP Server는 5060에서 실행된다.

## Example Ports and Services
다음은 서비스 및 관련 포트의 작은 목록이다. IANA - TCP/IP 포트 할당에서 최신 인터넷 포트 및 관련 서비스 목록을 찾을 수 있다. 

Service | Port Number | Service Description
------- | ------ | ------
echo | 7 | UDP/TCP가 수신한 내용을 다시 보낸다.
discard | 9 | UDP/TCP가 입력을 버린다.
daytime | 13 | UDP/TCP가 ASCII 시간을 반환한다.
chargen | 19 | UDP/TCP가 문자를 반환한다.
ftp | 21 | TCP 파일 전송.
telnet | 23 | TCP 원격 로그인.
smtp | 25 | TCP 이메일.
daytime | 37 | UDP/TCP가 바이너리 시간을 반환한다.
tftp | 69 | UDP 사소한 파일 전송.
finger | 79 | 사용자에 대한 TCP 정보.
http | 80 | TCP World Wide Web.
login | 513 | TCP 원격 로그인.
who | 513 | UDP 사용자에 대한 다른 정보.
Xserver | 6000 | TCP X windows (N.B. > 1023).

## Port and Service Functions
Unix는 /etc/services 파일에서 서비스 이름을 가져오는 다음과 같은 function을 제공한다.
* struct servent *getservbyname(char *name, char *proto)
    * 이 호출은 서비스 이름과 프로토콜 이름을 가지고 해당 서비스에 해당하는 포트 번호를 반환한다.
* struct servent *getservbyport(int port, char *proto)
    * 이 호출은 포트 번호와 프로토콜 이름을 가지고 해당 서비스 이름을 반환한다.

```
struct servent {
    char    *s_name;
    char    **s_aliases;
    int     s_port;
    char    *s_proto;
}
```

Attribute | Values | Description
------- | ------ | ------
s_name | http | 서비스의 공식 이름. 예를 들어 SMTP, FTP POP3 등.
s_aliases | ALIAS | 서비스 별칭 리스트. 대부분 이는 NULL로 설정된다.
s_port | 80 | 연관된 포트 번호. HTTP의 경우 80
s_proto | TCP, UDP | 사용된 프로토콜이 설정된다. 인터넷 서비스는 TCP 또는 UDP를 사용하여 제공된다.

## Reference
* https://www.tutorialspoint.com/unix_sockets/ports_and_services.htm