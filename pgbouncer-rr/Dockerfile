FROM alpine:3.8
MAINTAINER Felipe Cavalcanti <fjfcavalcanti@gmail.com>

WORKDIR /opt/pgbouncer

RUN apk add --no-cache openssl \
      c-ares \
      python \
      libevent && \
      apk add --no-cache -t .build-deps \
      libevent-dev \
      openssl-dev \
      autoconf \
      m4 \
      git	\
      libtool	\
      python-dev \
      automake \
      build-base \
      gcc \
      wget \
      patch \
      c-ares-dev && \
      git clone https://github.com/pgbouncer/pgbouncer.git /opt/pgbouncer && \
      git clone https://github.com/awslabs/pgbouncer-rr-patch.git /opt/pgbouncer-rr-patch && \
      cd /opt/pgbouncer-rr-patch && ./install-pgbouncer-rr-patch.sh /opt/pgbouncer && \
      cd /opt/pgbouncer && \
      mkdir -p /etc/pgbouncer /var/log/pgbouncer /var/run/pgbouncer && \
      cp etc/pgbouncer.ini /etc/pgbouncer/pgbouncer.ini.example && \
      cp etc/userlist.txt /etc/pgbouncer/userlist.txt.example && \
      chown -R postgres /var/run/pgbouncer /etc/pgbouncer /var/log/pgbouncer && \
      git submodule init && \ 
      git submodule update && \
      ./autogen.sh && \
      ./configure && \
      make && \
      make install && \
      rm -rf /opt/pgbouncer /opt/pgbouncer-rr-patch && \
      apk del --purge .build-deps

ADD entrypoint.sh /entrypoint.sh

USER postgres
EXPOSE 5432
VOLUME /etc/pgbouncer

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/usr/local/bin/pgbouncer", "/etc/pgbouncer/pgbouncer.ini"]
