FROM postgres:14.2 as dumper

COPY testdb.sql /docker-entrypoint-initdb.d/

RUN ["sed", "-i", "s/exec \"$@\"/echo \"skipping...\"/", "/usr/local/bin/docker-entrypoint.sh"]

ENV POSTGRES_USER=ziyan
ENV POSTGRES_PASSWORD=postgres
ENV PGDATA=/data

RUN ["/usr/local/bin/docker-entrypoint.sh", "postgres"]

# final build stage
FROM postgres:14.2

COPY --from=dumper /data $PGDATA
