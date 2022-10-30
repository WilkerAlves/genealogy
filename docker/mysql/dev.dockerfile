FROM mysql:latest

COPY ./resources/create_and_populate.sql /tmp

CMD [ "mysqld", "--init-file=/tmp/create_and_populate.sql" ]