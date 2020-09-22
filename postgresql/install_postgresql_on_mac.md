# Install postgresql on Mac
```bash
# install postgresql
$ brew install postgresql

# run postgresql
$ pg_ctl -D /usr/local/var/postgres start
waiting for server to start....2020-09-22 16:21:56.722 KST [27793] LOG:  starting PostgreSQL 12.4 on x86_64-apple-darwin19.5.0, compiled by Apple clang version 11.0.3 (clang-1103.0.32.62), 64-bit
2020-09-22 16:21:56.724 KST [27793] LOG:  listening on IPv6 address "::1", port 5432
2020-09-22 16:21:56.724 KST [27793] LOG:  listening on IPv4 address "127.0.0.1", port 5432
2020-09-22 16:21:56.725 KST [27793] LOG:  listening on Unix socket "/tmp/.s.PGSQL.5432"
2020-09-22 16:21:56.735 KST [27794] LOG:  database system was shut down at 2020-09-22 16:16:45 KST
2020-09-22 16:21:56.739 KST [27793] LOG:  database system is ready to accept connections
 done
server started

# make sure pogstresql starts every time your computer starts up
$ brew services start postgresql

# check installed version
$ postgres -V
postgres (PostgreSQL) 12.4

# stop postgresql server
$ pg_ctl -D /usr/local/var/postgres stop
```

# Create Postgres Database
```bash
# create database
$ createdb [DB_NAME]

# connect to database via psql with db name
# it leads you to the psql shell..
$ psql [DB_NAME]
psql (12.4)
Type "help" for help.

[DB_NAME]=# exit

# connet to database via psql with options
$ psql -h [HOST_NAME] -p [PORT] -U [USER_NAME] -W -d [DB_NAME]
# or 
$ psql postgres://[USER_NAME]:[PASSWORD]@[HOST_NAME]:[PORT]/[DB_NAME]?sslmode=require
```

# Postgresql GUI Client Tool on Mac
* DBeaver: https://dbeaver.io/
