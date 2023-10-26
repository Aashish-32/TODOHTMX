from postgres:14 

# Copying your custom initialization script into the container

copy InitDB/init.sql  /docker-entrypoint-initdb.d

ENV POSTGRES_USER=todo
ENV POSTGRES_PASSWORD=todo1234
ENV POSTGRES_DB=todo

# Specifies the CMD to run the PostgreSQL server when container is created.

CMD ["docker-entrypoint.sh", "postgres"]