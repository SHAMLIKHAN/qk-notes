FROM postgres:10-alpine

ENV POSTGRES_USER postgres
ENV POSTGRES_DB qknote
ENV POSTGRES_PASSWORD password
ENV POSTGRES_PORT 5432

EXPOSE 5432

ADD postgres_up.sql /docker-entrypoint-initdb.d/