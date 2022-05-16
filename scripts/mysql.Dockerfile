FROM mysql:latest
ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_ROOT_PASSWORD

ENV MYSQL_DATABASE=calendar
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=test
ENV MYSQL_ROOT_PASSWORD=test
ENV LANG=C.UTF-8

ADD scripts/schema.sql /docker-entrypoint-initdb.d/1.init.sql
ADD scripts/data.sql /docker-entrypoint-initdb.d/2.data.sql

EXPOSE 3306

CMD ["mysqld"]