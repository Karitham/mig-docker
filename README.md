# mig-docker

A simple cli tool to copy migration files from their original location (`./db/migrations`) to the specified out directory.

Params:
`-up` specifies the up filename. default is `up.sql`.
`-migrations` specifies the migration directory. default is `./db/migrations`.

This tool assumes your migration directories are named with only numbers (such as unix timestamps of creation, or sequential numbers).

It will order them by this same number found in the directory name.

## Usage

```docker
FROM postgres:14.2-bullseye

RUN apt-get update -y
RUN apt-get install wget -y
RUN wget https://github.com/Karitham/mig-docker/releases/download/v0.1.0/mig-docker_0.1.0_linux_x86_64.tar.gz -O mig-docker.tar.gz
RUN tar -xvf mig-docker.tar.gz

COPY ./db/schema.sql /docker-entrypoint-initdb.d/0.sql
COPY ./db/migrations ./

RUN ./mig-docker -migrations ./ /docker-entrypoint-initdb.d/
```
