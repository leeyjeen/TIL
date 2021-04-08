# SSL Certificate
## What is TLS

**TLS**는 **SSL** 기반이며, SSLv3의 알려진 취약성에 대응한 대체품으로서 개발되었다.

SSL은 일반적으로 사용되는 용어이며, 요즘에는 일반적으로 TLS를 나타낸다.

## Security Provided

SSL/TLS는 데이터 암호화, 데이터 무결성 및 인증을 제공한다.

즉, SSL/TLS를 사용하는 경우 다음을 확신할 수 있다.

- 아무도 메시지를 읽지 않았다.
- 아무도 메시지를 변경하지 않았다.
- 의도한 사람(서버)과 통신 중이다.

두 사람 간에 메시지를 보낼 때 해결해야 할 두 문제가 있다.

- 아무도 메시지를 읽지 않았다는 걸 어떻게 알까?
- 아무도 메시지를 변경하지 않았다는 걸 어떻게 알까?

이러한 문제의 해결 방법은 다음과 같다.

- **암호화**
    - 이렇게 하면 내용을 읽을 수 없으므로 메시지를 보는 모든 사람이 이해할 수 없다.
- **서명**
    - 메시지를 보낸 사람이 본인이며, 메시지가 변경되지 않았음을 수신인이 확신할 수 있다.

이 두 프로세스 모두 키를 사용해야 한다.

이러한 키는 단순한 숫자(일반적으로 128비트인 경우)로, 메시지를 암호화하거나 서명하기 위해 일반적으로 알고리즘(예: RSA)으로 알려진 특정 방법을 사용하여 메시지와 결합된다.

## Symmetrical Keys and Public and Private Keys


## Keys and SSL Certificates


## Obtaining a Digital Certificate



## Reference
- http://www.steves-internet-guide.com/ssl-certificates-explained/