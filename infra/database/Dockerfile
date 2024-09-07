FROM mysql:8 as database
ENV MYSQL_ROOT_PASSWORD=myunsafepassword+123
COPY ./ /docker-entrypoint-initdb.d/
# docker build -t "to-do-list-database:1.0" --target database .
# docker run -p 3306:3306 --name "to-do-list-database" "to-do-list-database:1.0"
# docker run -d -e MYSQL_USER=$DBUSER -e MYSQL_PASSWORD=$DBPASS -p 3306:3306 --name "to-do-list-database" "to-do-list-database:1.0"
    # EX: docker run -d -e MYSQL_USER=todolistuser -e MYSQL_PASSWORD=myunsafepassword+123 -p 3306:3306 --name "to-do-list-database" "to-do-list-database:1.0"
# docker run -d -e MYSQL_ROOT_USER=root -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 --name "to-do-list-database" "to-do-list-database:1.0"
# docker exec -it to-do-list-database bash
