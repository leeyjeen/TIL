# Network Host Names
숫자의 측면에서 호스트 이름은 기억하기 어려우므로 그들은 paris, seoul과 같은 평범한 이름으로 불린다. 우리는 주어진 이름에 해당하는 점으로 구분된 IP 주소를 찾기 위해 소프트웨어 응용 프로그램을 작성한다.

주어진 영숫자 호스트 이름을 기반으로 점으로 구분된 IP 주소를 찾는 프로세스를 **hostname resolution**(호스트 이름 확인)이라고 한다. 

호스트 이름 확인은 고용량 시스템에 있는 특수 소프트웨어에 의해 수행된다. 이러한 시스템을 DNS(Domain Name Systems)라고 하며 IP 주소 및 해당 일반 이름의 매핑을 유지한다. 

## The /etc/hosts File
호스트 이름과 IP 주소 간의 대응은 `hosts`라는 파일에서 유지된다. 대부분의 시스템에서 이 파일은 `/etc` 디렉터리에 있다.

이 파일의 항목은 다음과 같다.
```
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
66.249.89.104   www.google.com
::1             localhost
# Added by Docker Desktop
# To allow the same kube context to work on the host and the container:
127.0.0.1 kubernetes.docker.internal
# End of section
```

두 개 이상의 이름이 주어진 IP 주소와 연관될 수 있다. 이 파일은 IP 주소에서 호스트 이름으로 또는 그 반대로 변환하는 동안 사용된다.

이 파일을 편집할 권한이 없으므로 IP 주소와 함께 호스트 이름을 입력하려면 루트 권한이 있어야 한다.

## Reference
* https://www.tutorialspoint.com/unix_sockets/network_host_names.htm