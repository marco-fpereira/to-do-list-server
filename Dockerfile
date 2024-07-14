FROM mysql:8

ENV MYSQL_ROOT_PASSWORD=root
COPY ./database/ /docker-entrypoint-initdb.d/
# docker build -t "to-do-list-database:1.0" .
# docker run -d -e MYSQL_USER=$DBUSER -e MYSQL_ROOT_PASSWORD=$DBPASS -p 3306:3306 --name "to-do-list-database" "to-do-list-database:1.0"
# docker exec -it to-do-list-database bash
