FROM debian:stretch-slim

# RUN groupadd -r mysql && useradd -r -g mysql mysql

# RUN apt-get update && apt-get install -y --no-install-recommends gnupg dirmngr && rm -rf /var/lib/apt/lists/*

# RUN mkdir /docker-entrypoint-initdb.sh

ENV MYSQL_MAJOR 5.6
ENV MYSQL_VERSION 5.6.51-1debian9

# RUN echo 'deb http://repo.mysql.com/apt/debian/ stretch mysql-5.6' > /etc/apt/sources.list.d/mysql.list

VOLUME /varlib/mysql

# COPY docker-entrypoint.sh /usr/local/bin/

#RUN ln -s usr/local/bin/docker-entrypoint.sh /entrypoint.sh 
#ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 3306
#CMD ["mysqld"]
