FROM postgres:15.3-alpine

LABEL author=Dima191

ENV POSTGRES_DB=temp

COPY ./init.sql /docker-entrypoint-initdb.d/