FROM mysql:8 as database
ENV MYSQL_ROOT_PASSWORD=myunsafepassword+123
COPY ./ /docker-entrypoint-initdb.d/
