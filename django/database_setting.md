# DATABASES settings in django
```DATABASES``` 설정은 ```default``` 데이터베이스를 구성해야 하며, 추가 데이터베이스도 얼마든지 설정할 수 있다.

가능한 가장 간단한 설정 파일은 SQLite를 이용한 단일 데이터베이스 설정 파일이다. 이는 다음과 같이 구성할 수 있다.:

```python
# SQLite
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': 'mydatabase',
    }
}
```

MySQL, Oracle 또는 PostgreSQL과 같은 다른 데이터베이스 백엔드에 연결할 때는 추가적인 커넥션 매개 변수가 필요하다. 다른 데이터베이스 유형을 지정하는 방법은 아래 ```ENGINE``` 설정을 참조하자. 이 예는 ```PostgreSQL```에 관한 것이다.:

```python
# PostgreSQL
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql_psycopg2',
        'NAME': 'mydatabase',
        'USER': 'mydatabaseuser',
        'PASSWORD': 'mypassword',
        'HOST': '127.0.0.1',
        'PORT': '5432',
    }
}
```

다음은 보다 복잡한 구성에 필요할 수 있는 내부 옵션들 중 일부이다.

## ENGINE
Default: ''
사용할 데이터베이스 백엔드를 설정한다. 내장 데이터베이스 백엔드는 다음과 같다.:

- ```'django.db.backends.postgresql_psycopg2'```
- ```'django.db.backends.mysql'```
- ```'django.db.backends.sqlite3'```
- ```'django.db.backends.oracle'```

```ENGINE을``` 정규화된 경로(i.e. ```mypackage.backends.whatever```)로 설정하여 Django와 함께 제공되지 않는 데이터베이스를 사용할 수 있다.

## HOST
Default: ''

데이터베이스에 연결할 때 사용할 호스트를 설정한다. 빈 문자열은 ```localhost```를 의미한다. SQLite와는 쓰이지 않는다.

이 값이 슬래시('/')로 시작하고, MySQL을 사용하는 경우 MySQL은 Unix 소켓을 통해 지정된 소켓에 연결된다. 예를 들면 다음과 같다.:

```
"HOST": '/var/run/mysql'
```

**MySQL**을 사용 중인데 이 값이 슬래시로 시작하지 않는다면, 이 값은 호스트로 가정된다.

**PostgreSQL**을 사용한다면 기본적으로 데이터베이스 연결이 UNIX 도메인 소켓(‘local’ lines in **pg_hba.conf**)을 통해 이루어진다. UNIX 도메인 소켓이 표준 위치에 없는 경우 **postgresql.conf**의 **unix_socket_directory**과 동일한 값을 사용한다. TCP 소켓을 통해 연결하려면 HOST를 ‘localhost’ 또는 ‘127.0.0.1’으로 설정한다. Windows 환경에서는 UNIX 도메인 소켓이 가능하지 않기 때문에 항상 **HOST**를 정의해야 한다. 

## NAME
Default: ''

사용할 데이터베이스의 이름을 설정한다. SQLite의 경우 데이터베이스 파일에 대한 전체 경로이다. 경로를 명시할 때, 항상 슬래시를 사용한다. (e.g. ```C:/homes/user/mysite/sqlite3.db```).

## PASSWORD
Default: ''

데이터베이스에 연결할 때 사용할 패스워드를 설정한다. SQLite의 경우 사용되지 않는다.

## PORT
Default: ''

데이터베이스에 연결할 때 사용할 포트를 설정한다. 빈 문자열은 기본 포트를 의미한다. SQLite의 경우 사용되지 않는다.

## USER
Default: ''

데이터베이스에 연결할 때 사용할 username을 설정한다. SQLite의 경우 사용되지 않는다.

## TEST
Default: {}

테스트 데이터베이스를 위한 설정 딕셔너리 값이다. 

다음은 테스트 데이터베이스 설정 예시이다.:

```python
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'USER': 'mydatabaseuser',
        'NAME': 'mydatabase',
        'TEST': {
            'NAME': 'mytestdatabase',
        },
    },
}
```

## Reference
- https://docs.djangoproject.com/en/1.8/ref/settings/#databases