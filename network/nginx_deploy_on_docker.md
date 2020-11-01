# Deploying NGINX and NGINX Plus on Docker
고성능 애플리케이션 제공 플랫폼, 로드밸런서 및 웹 서버인 NGINX Plus는 Docker 컨테이너로 사용할 수 있다.

## Prerequisites
* Docker 설치
* Docker Hub 계정 (NGINX 오픈소스)
* `nginx-repo.crt` 및 `nginx-repo.key` 파일, Docker 이미지 생성용 Dockerfile (NGINX Plus)

## Running NGINX Open Source in a Docker Container
Docker Hub의 NGINX 오픈소스 이미지를 사용하여 Docker 컨테이너에 NGINX 인스턴스를 생성할 수 있다.

1. 다음 명령으로 기본 NGINX 구성을 사용하여 컨테이너에서 실행중인 NGINX 인스턴스를 시작한다.
```
$ docker run --name mynginx1 -p 80:80 -d nginx
```
* `mynginx1`: NGINX 이미지에 근거하여 생성된 컨테이너 이름
* `-d` 옵션: 컨테이너가 분리된 모드로 실행되도록 지정한다. 컨테이너는 중지될 때까지 계속 실행되지만 명령줄에서 실행되는 명령에는 응답하지 않는다.
* `-p` 옵션: NGINX 이미지(포트80)에 의해 컨테이너에 노출된 포트를 Docker 호스트의 지정된 포트에 매핑하도록 Docker에 지시한다. 첫 번째 매개변수는 Docker 호스트의 포트를 지정하고 두 번째 매개변수는 컨테이너에 노출된 포트에 매핑된다.

이 명령은 긴 형식의 컨테이너 ID `fcd1fb01b14557c7c9d991238f2558ae2704d129cf9fb97bb4fadf673a58580d`를 반환한다. 이 형식의 ID는 로그 파일의 이름에 사용된다.

2. 컨테이너가 생성되었고 `docker ps` 명령으로 실행 중인지 확인한다.
```
$ docker ps
CONTAINER ID  IMAGE         COMMAND               CREATED         STATUS        ...  
fcd1fb01b145  nginx:latest  "nginx -g 'daemon of  16 seconds ago  Up 15 seconds ... 

    ... PORTS              NAMES
    ... 0.0.0.0:80->80/tcp mynginx1
```

이 명령을 사용하면 이전 단계에서 설정한 포트 매핑을 확인할 수도 있다. output의 PORTS 필드는 Docker 호스트의 포트 80이 컨테이너의 포트 80에 매핑되었음을 알린다.

## Running NGINX Plus in a Docker Container
Docker는 NGINX Plus와 함께 사용할 수도 있다. NGINX 오픈소스와 함께 Docker를 사용하는 것의 차이점은 상업용 제품으로 NGINX Plus를 Docker Hub에서 사용할 수 없기 때문에 먼저 NGINX Plus 이미지를 만들어야 한다는 것이다.

> 참고: NGINX Plus 이미지를 Docker Hub와 같은 공용 저장소에 업로드하지 마라. 그렇게 하면 라이센스 계약에 위배된다.

### Creating NGINX Plus Docker Image
NGINX Plus 이미지를 생성하기 위하여,
1. Docker build 컨텍스트 또는 Dockerfile을 생성한다.
```dockerfile
FROM debian:buster-slim

LABEL maintainer="NGINX Docker Maintainers <docker-maint@nginx.com>"

# Define NGINX versions for NGINX Plus and NGINX Plus modules
# Uncomment this block and the versioned nginxPackages block in the main RUN
# instruction to install a specific release
# ENV NGINX_VERSION 21
# ENV NJS_VERSION 0.3.9
# ENV PKG_RELEASE 1~buster

# Download certificate and key from the customer portal (https://cs.nginx.com)
# and copy to the build context
COPY nginx-repo.crt /etc/ssl/nginx/
COPY nginx-repo.key /etc/ssl/nginx/

RUN set -x \
# Create nginx user/group first, to be consistent throughout Docker variants
    && addgroup --system --gid 101 nginx \
    && adduser --system --disabled-login --ingroup nginx --no-create-home --home /nonexistent --gecos "nginx user" --shell /bin/false --uid 101 nginx \
    && apt-get update \
    && apt-get install --no-install-recommends --no-install-suggests -y ca-certificates gnupg1 \
    && \
    NGINX_GPGKEY=573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62; \
    found=''; \
    for server in \
        ha.pool.sks-keyservers.net \
        hkp://keyserver.ubuntu.com:80 \
        hkp://p80.pool.sks-keyservers.net:80 \
        pgp.mit.edu \
    ; do \
        echo "Fetching GPG key $NGINX_GPGKEY from $server"; \
        apt-key adv --keyserver "$server" --keyserver-options timeout=10 --recv-keys "$NGINX_GPGKEY" && found=yes && break; \
    done; \
    test -z "$found" && echo >&2 "error: failed to fetch GPG key $NGINX_GPGKEY" && exit 1; \
    apt-get remove --purge --auto-remove -y gnupg1 && rm -rf /var/lib/apt/lists/* \
# Install the latest release of NGINX Plus and/or NGINX Plus modules
# Uncomment indivisual modules if necessary
# Use versioned packages over defaults to specify a release
    && nginxPackages=" \
        nginx-plus \
        # nginx-plus=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-xslt \
        # nginx-plus-module-xslt=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-geoip \
        # nginx-plus-module-geoip=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-image-filter \
        # nginx-plus-module-image-filter=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-perl \
        # nginx-plus-module-perl=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-njs \
        # nginx-plus-module-njs=${NGINX_VERSION}+${NJS_VERSION}-${PKG_RELEASE} \
    " \
    && echo "Acquire::https::plus-pkgs.nginx.com::Verify-Peer \"true\";" >> /etc/apt/apt.conf.d/90nginx \
    && echo "Acquire::https::plus-pkgs.nginx.com::Verify-Host \"true\";" >> /etc/apt/apt.conf.d/90nginx \
    && echo "Acquire::https::plus-pkgs.nginx.com::SslCert     \"/etc/ssl/nginx/nginx-repo.crt\";" >> /etc/apt/apt.conf.d/90nginx \
    && echo "Acquire::https::plus-pkgs.nginx.com::SslKey      \"/etc/ssl/nginx/nginx-repo.key\";" >> /etc/apt/apt.conf.d/90nginx \
    && printf "deb https://plus-pkgs.nginx.com/debian buster nginx-plus\n" > /etc/apt/sources.list.d/nginx-plus.list \
    && apt-get update \
    && apt-get install --no-install-recommends --no-install-suggests -y \
                        $nginxPackages \
                        gettext-base \
                        curl \
    && apt-get remove --purge --auto-remove -y && rm -rf /var/lib/apt/lists/* /etc/apt/sources.list.d/nginx-plus.list \
    && rm -rf /etc/apt/apt.conf.d/90nginx /etc/ssl/nginx

# Forward request logs to Docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]
```

```dockerfile

FROM alpine:3.11

LABEL maintainer="NGINX Docker Maintainers <docker-maint@nginx.com>"

# Define NGINX versions for NGINX Plus and NGINX Plus modules
# Uncomment this block and the versioned nginxPackages in the main RUN
# instruction to install a specific release
# ENV NGINX_VERSION 21
# ENV NJS_VERSION   0.3.9
# ENV PKG_RELEASE   1

# Download certificate and key from the customer portal (https://cs.nginx.com)
# and copy to the build context
COPY nginx-repo.crt /etc/apk/cert.pem
COPY nginx-repo.key /etc/apk/cert.key

RUN set -x \
# Create nginx user/group first, to be consistent throughout Docker variants
    && addgroup -g 101 -S nginx \
    && adduser -S -D -H -u 101 -h /var/cache/nginx -s /sbin/nologin -G nginx -g nginx nginx \
# Install the latest release of NGINX Plus and/or NGINX Plus modules
# Uncomment individual modules if necessary
# Use versioned packages over defaults to specify a release
    && nginxPackages=" \
        nginx-plus \
        # nginx-plus=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-xslt \
        # nginx-plus-module-xslt=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-geoip \
        # nginx-plus-module-geoip=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-image-filter \
        # nginx-plus-module-image-filter=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-perl \
        # nginx-plus-module-perl=${NGINX_VERSION}-${PKG_RELEASE} \
        # nginx-plus-module-njs \
        # nginx-plus-module-njs=${NGINX_VERSION}.${NJS_VERSION}-${PKG_RELEASE} \
    " \
    KEY_SHA512="e7fa8303923d9b95db37a77ad46c68fd4755ff935d0a534d26eba83de193c76166c68bfe7f65471bf8881004ef4aa6df3e34689c305662750c0172fca5d8552a *stdin" \
    && apk add --no-cache --virtual .cert-deps \
        openssl \
    && wget -O /tmp/nginx_signing.rsa.pub https://nginx.org/keys/nginx_signing.rsa.pub \
    && if [ "$(openssl rsa -pubin -in /tmp/nginx_signing.rsa.pub -text -noout | openssl sha512 -r)" = "$KEY_SHA512" ]; then \
        echo "key verification succeeded!"; \
        mv /tmp/nginx_signing.rsa.pub /etc/apk/keys/; \
    else \
        echo "key verification failed!"; \
        exit 1; \
    fi \
    && apk del .cert-deps \
    && apk add -X "https://plus-pkgs.nginx.com/alpine/v$(egrep -o '^[0-9]+\.[0-9]+' /etc/alpine-release)/main" --no-cache $nginxPackages \
    && if [ -n "/etc/apk/keys/nginx_signing.rsa.pub" ]; then rm -f /etc/apk/keys/nginx_signing.rsa.pub; fi \
    && if [ -n "/etc/apk/cert.key" && -n "/etc/apk/cert.pem"]; then rm -f /etc/apk/cert.key /etc/apk/cert.pem; fi \
# Bring in gettext so we can get `envsubst`, then throw
# the rest away. To do this, we need to install `gettext`
# then move `envsubst` out of the way so `gettext` can
# be deleted completely, then move `envsubst` back.
    && apk add --no-cache --virtual .gettext gettext \
    && mv /usr/bin/envsubst /tmp/ \
    \
    && runDeps="$( \
        scanelf --needed --nobanner /tmp/envsubst \
            | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
            | sort -u \
            | xargs -r apk info --installed \
            | sort -u \
    )" \
    && apk add --no-cache $runDeps \
    && apk del .gettext \
    && mv /tmp/envsubst /usr/local/bin/ \
# Bring in tzdata so users could set the timezones through the environment
# variables
    && apk add --no-cache tzdata \
# Bring in curl and ca-certificates to make registering on DNS SD easier
    && apk add --no-cache curl ca-certificates \
# Forward request and error logs to Docker log collector
    && ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]

# vim:syntax=Dockerfile
```

2. NGINX 오픈소스와 마찬가지로 기본 NGINX Plus 이미지에는 동일한 기본 설정이 있다.
* access logs, error logs가 Docker log 컬렉터에 연결되어 있다.
* 볼륨이 지정되지 않는다. Dockerfile을 사용하여 지정된 볼륨으로 새 이미지를 생성할 수 있는 기본 이미지를 생성하거나 볼륨을 수동으로 지정할 수 있다.
```dockerfile
VOLUME /usr/share/nginx/html
VOLUME /etc/nginx
```
* 컨테이너가 생성될 때 Docker 호스트에서 파일이 복사되지 않는다. 각 Dockerfile에 `COPY` 정의를 추가하거나 생성한 이미지를 다른 이미지의 기반으로 사용할 수 있다.

3. MyF5 Customer Portal에 로그인하고 nginx-repo.crt 및 nginx-repo.key 파일을 다운로드한다. NGINX Plus 평가판의 경우 평가판 패키지와 함께 파일이 제공된다.

4. Dockerfile이 위치한 디렉터리에 파일을 복사한다.

5. Docker 이미지를 생성한다. 예를 들어 `nginxplus`라는 이미지.
```
$ docker build --no-cache -t nginxplus .
```
--no-cache 옵션은 Docker에 처음부터 이미지를 빌드하도록 지시하고 최신 버전의 NGINX Plus 설치를 보장한다. 이전에 Dockerfile을 사용하여 --no-cache 옵션없이 이미지를 빌드한 경우 새 이미지는 Docker 캐시에서 이전에 빌드된 이미지의 NGINX Plus 버전을 사용한다.

6. `nginxplus` 이미지가 성공적으로 생성되었음을 `docker images` 명령을 통해 확인한다.
```
$ docker images nginxplus
REPOSITORY  TAG     IMAGE ID      CREATED        SIZE
nginxplus   latest  ef2bf65931cf  6 seconds ago  91.2 MB
```

7. 이 이미지를 기반으로 컨테이너를 생성한다. 예를 들어, `mynginxplus` 컨테이너를 생성.
```
$ docker run --name mynginxplus -p 80:80 -d nginxplus
```

8. `docker ps` 명령을 사용하여 `mynginxplus` 컨테이너가 실행되고 있는지 확인한다.
```
$ docker ps
CONTAINER ID  IMAGE         COMMAND               CREATED         STATUS        ...  
eb7be9f439db  nginx:latest  "nginx -g 'daemon of  1 minute ago  Up 15 seconds ... 

    ... PORTS              NAMES
    ... 0.0.0.0:80->80/tcp mynginxplus
```

NGINX Plus 컨테이너는 NGINX 오픈소스 컨테이너와 동일한 방식으로 제어 및 관리된다.

## Managing Content and Configuration Files
NGINX 및 NGINX 구성 파일에서 제공하는 콘텐츠는 여러 방법으로 관리할 수 있다.

* 파일은 Docker 호스트에서 유지된다.
* 파일은 Docker 호스트에서 컨테이너로 복사된다.
* 파일은 컨테이너에 유지된다.

### Maintaining Content and Configuration Files on the Docker Host
컨테이너가 생성되면 Docker 호스트의 로컬 디렉터리를 컨테이너의 디렉터리에 탑재할 수 있다. NGINX 이미지는 `/usr/share/nginx/html`을 컨테이너의 루트 디렉터리로 사용하고 구성 파일을 `/etc/nginx`에 저장하는 기본 NGINX 구성을 사용한다. 로컬 디렉터리 `/var/www`에 콘텐츠가 있고 `/var/nginx/conf`에 구성 파일이 있는 Docker 호스트의 경우 다음 명령을 실행한다.

```
$ docker run --name mynginx2 --mount type=bind,source=/var/www,target=/usr/share/nginx/html,readonly --mount source=/var/nginx/conf,target=/etc/nginx/conf,readonly -p 80:80 -d nginx
```

Docker 호스트의 로컬 디렉터리 `/var/www` 및 `/var/nginx/conf`에 있는 파일에 대한 변경 사항은 컨테이너의 `/usr/share/nginx/html` 및 `/etc/nginx` 디렉터리에 반영된다. readonly 옵션은 이러한 디렉터리가 컨테이너 내부가 아닌 Docker 호스트에서만 변경될 수 있음을 의미한다.

### Copying Content and Configuration Files from the Docker Host
Docker는 컨테이너 생성 중에 Docker 호스트의 로컬 디렉터리에서 콘텐츠 및 구성 파일을 복사할 수 있다. 컨테이너가 생성되면 파일이 변경될 때 새 컨테이너를 생성하거나 컨테이너의 파일을 수정하여 파일을 유지한다.

파일을 복사하는 간단한 방법은 NGINX 이미지를 기반으로 새 Docker 이미지를 생성하는 동안 실행되는 명령으로 Dockerfile을 만드는 것이다. Dockerfile의 파일 복사(COPY) 명령의 경우 로컬 디렉터리 경로는 Dockerfile이 있는 빌드 컨텍스트에 상대적이다. 

콘텐츠 디렉터리가 콘텐츠이고 구성 파일의 디렉터리가 Dockerfile이 있는 디렉터리의 두 하위 디렉터리인 conf라고 가정해보겠다. NGINX 이미지에는 `/etc/nginx/conf.d` 디렉터리에 `default.conf`를 포함한 기본 NGINX 구성 파일이 있다. Docker 호스트에서만 구성 파일을 사용하려면 RUN 명령으로 기본 파일을 삭제한다.

```dockerfile
FROM nginx
RUN rm /etc/nginx/conf.d/default.conf
COPY content /usr/share/nginx/html
COPY conf /etc/nginx
```

Dockerfile이 있는 디렉터리에서 명령을 실행하여 NGINX 이미지를 만든다. 명령 끝에 있는 마침표(".")는 현재 디렉터리를 Dockerfile 및 복사할 디렉터리를 포함하는 빌드 컨텍스트로 정의한다.
```
$ docker build -t mynginx_image1 .
```

`mynginx_image1` 이미지를 기반으로 `mynginx3` 컨테이너를 생성한다.
```
$ docker run --name mynginx3 -p 80:80 -d mynginx_image1
```

컨테이너의 파일을 변경하려면 다음 섹션에 설명된대로 helper 컨테이너를 사용한다.

### Maintaining Content and Configuration Files in the Container
NGINX 컨테이너에 액세스하는 데 SSH를 사용할 수 없으므로 콘텐츠 또는 구성 파일을 직접 편집하려면 셸 액세스 권한이 있는 helper 컨테이너를 생성해야 한다. helper 컨테이너가 파일에 액세스할 수 있도록 이미지에 대해 정의된 적절한 Docker 데이터 볼륨이 있는 새 이미지를 생성한다.

1. nginx 콘텐츠 및 구성 파일을 복사하고 Dockerfile을 사용하여 이미지의 볼륨을 정의한다.
```dockerfile
FROM nginx
COPY content /usr/share/nginx/html
COPY conf /etc/nginx
VOLUME /usr/share/nginx/html
VOLUME /etc/nginx
```

2. 다음 명령을 실행하여 새 NGINX 이미지를 생성한다.
```
$ docker build -t mynginx_image2 .
```

3. `mynginx_image2` 이미지를 기반으로 `mynginx4` NGINX 컨테이너를 생성한다.
```
$ docker run --name mynginx4 -p 80:80 -d mynginx_image2
```

4. 셸이 있는 helper 컨테이너 `mynginx4_files`를 시작하여 방금 만든 `mynginx4` 컨테이너의 콘텐츠 및 구성 디렉터리에 대한 액세스를 제공한다.
```
$ docker run -i -t --volumes-from mynginx4 --name mynginx4_files debian /bin/bash
root@b1cbbad63dd1:/#
```
* 새로운 `mynginx4_files` helper 컨테이너는 영구 표준 입력 (-i 옵션) 및 tty(-t 옵션)을 사용하여 포그라운드에서 실행된다. `mynginx4`에 정의된 모든 볼륨은 helper 컨테이너의 로컬 디렉터리로 마운트된다.
* debian 인수는 helper 컨테이너가 Docker Hub의 Debian 이미지를 사용함을 의미한다. NGINX 이미지도 Debian을 사용하기 때문에 Docker가 다른 운영체제를 로드하도록 하는 것보다 helper 컨테이너로 Debian을 사용하는 것이 가장 효율적이다.
* `/bin/bash` 인수는 bash 셸이 helper 컨테이너에서 실행됨을 의미하며 필요에 따라 파일을 수정하는 데 사용할 수 있는 셸 프롬프트를 표시한다.

컨테이너를 시작하고 중지하기 위해 명령을 실행한다.
```
$ docker start mynginx4_files
$ docker stop mynginx4_files
```

셸을 종료하고 컨테이너는 계속 실행하려면 `Ctrl+p`를 누른 다음 `Ctrl+q`를 누른다. 실행중인 컨테이너ㅔ 대한 셸 액세스 권한을 다시 얻으려면 다음 명령어를 실행한다.
```
$ docker attach mynginx4_files
```

셸을 나가고 컨테이너를 종료하려면 `exit` 명령을 실행한다.

## Managing Logging
기본 로깅을 사용하거나 로깅을 커스터마이징할 수 있다.

### Using Default Logging
기본적으로 NGINX 이미지는 NGINX access log 및 error log를 Docker log 컬렉터로 보내도록 구성된다. 이는 stdout 및 stderr에 연결되어 수행된다. 두 로그의 모든 메시지는 Docker 호스트의 `/var/lib/docker/containers/container-ID/container-ID-json.log` 파일에 기록된다. container-ID는 컨테이너를 만들 때 반환되는 긴 형식의 ID이다. 긴 형식의 ID를 표시하려면 다음 명령을 실행한다.
```
$ docker inspect --format '{{ .Id }}' container-name
```

Docker 명령 줄과 Docker Engine API를 모두 사용하여 로그 메시지를 추출할 수 있다.

명령줄에서 로그 메시지를 추출하려면 다음 명령을 실행한다.
```
$ docker logs container-name
```

Docker Remote API를 사용하여 로그 메시지를 추출하려면 Docker Unix Sock을 사용하여 GET 요청을 전송한다.
```
$ curl --unix-sock /var/run/docker-sock http://localhost/containers/container-name/logs?stdout=1&stderr=1
```

출력에 access log 메시지만 포함하려면 stdout=1만 포함한다. 출력을 error log 메시지로 제한하려면 stderr=1만 포함한다. 사용 가능한 다른 옵션은 Docker Engine API 문서의 컨테이너 로그 가져오기 섹션을 참조하자.

### Using Customized Logging
특정 구성 블록(예: server {}, location {})에 대해 로깅을 다르게 구성하려면 컨테이너에 로그 파일을 저장할 디렉터리에 대한 Docker 볼륨을 정의하고 로그 파일에 액세스하기 위한 helper 컨테이너를 생성한다. 로깅 도구를 사용한다. 이를 구현하려면 로깅 파일에 대한 볼륨을 포함하는 새 이미지를 생성한다. 예를 들어 `/var/log/nginx/log`에 로그 파일을 저장하도록 NGINX를 구성하려면 디렉터리에 대한 VOLUME 정의를 Dockerfile에 추가한다(컨테이너에서 콘텐츠 및 구성 파일이 관리되는 경우).

```dockerfile
FROM nginx
COPY content /usr/share/nginx/html
COPY conf /etc/nginx
VOLUME /var/log/nginx/log
```

그런 다음 이미지를 생성하고 이를 사용하여 NGINX 컨테이너와 로깅 디렉터리에 대한 액세스 권한이 있는 helper 컨테이너를 생성할 수 있다. helper 컨테이너에는 원하는 로깅 도구를 설치할 수 있다.

## Controlling NGINX
NGINX 컨테이너의 명령 줄에 직접 액세스할 수 없기 때문에 NGINX 명령을 컨테이너로 직접 보낼 수 없다. 대신 Docker `kill` 명령을 통해 신호를 컨테이너로 보낼 수 있다.

NGINX 구성을 다시 로드하려면 HUP 신호를 Docker로 보낸다.

```
$ docker kill -s HUP container-name
```

NGINX를 재시작하려면 컨테이너를 재시작하는 명령을 실행한다.
```
$ docker restart container-name
```

## Reference
* https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-docker/