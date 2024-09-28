# Docker Formatter

This Go project is a Docker formatter tool that formats the output of Docker commands such as `docker ps` and `docker ps -a`. The tool enhances readability and makes the output easier to parse and work with.

`docker ps` output is:

```bash
CONTAINER ID   IMAGE         COMMAND                  CREATED              STATUS              PORTS                                                                                                                NAMES
157b98bc2e57   backend-app   "docker-entrypoint.s…"   2 seconds ago        Up 2 seconds        3000/tcp                                                                                                             nodejs-app-container
220a34d3a427   nginx         "/docker-entrypoint.…"   2 seconds ago        Up 2 seconds        0.0.0.0:80->80/tcp, 0.0.0.0:4343->4343/tcp, 0.0.0.0:6000->6000/tcp, 0.0.0.0:7070->7070/tcp, 0.0.0.0:9090->9090/tcp   nginx-container
d7eefeabceff   postgres:14   "docker-entrypoint.s…"   About a minute ago   Up About a minute   5432/tcp                                                                                                             postgres-container
27d6ce5e3b2c   redis         "docker-entrypoint.s…"   About a minute ago   Up About a minute   6379/tcp                                                                                                             redis-container
67f53bddbf26   server-app    "docker-entrypoint.s…"   3 minutes ago        Up 3 minutes        0.0.0.0:4000->8000/tcp                                                                                               app-v2
```

our `docker fps` (formatted ps) output will be:
```bash
 # SHOWING ONLY RUNNING CONTAINERS:
·--------------·-------------·----------------------·-------------------·-------------------------·
| CONTAINER ID |    IMAGE    |    CONTAINER NAME    |      STATUS       |          PORTS          |
·--------------·-------------·----------------------·-------------------·-------------------------·
| 157b98bc2e57 | backend-app | nodejs-app-container | Up 7 seconds      | 3000/tcp                |
|              |             |                      |                   |                         |
| 220a34d3a427 | nginx       | nginx-container      | Up 7 seconds      | 0.0.0.0:80->80/tcp,     |
|              |             |                      |                   | 0.0.0.0:4343->4343/tcp, |
|              |             |                      |                   | 0.0.0.0:6000->6000/tcp, |
|              |             |                      |                   | 0.0.0.0:7070->7070/tcp, |
|              |             |                      |                   | 0.0.0.0:9090->9090/tcp  |
|              |             |                      |                   |                         |
| d7eefeabceff | postgres:14 | postgres-container   | Up About a minute | 5432/tcp                |
|              |             |                      |                   |                         |
| 27d6ce5e3b2c | redis       | redis-container      | Up About a minute | 6379/tcp                |
|              |             |                      |                   |                         |
| 67f53bddbf26 | server-app  | app-v2               | Up 3 minutes      | 0.0.0.0:4000->8000/tcp  |
·--------------·-------------·----------------------·-------------------·-------------------------·
```


## Features

- Formats the output of `docker ps` for running containers.
- Formats the output of `docker ps -a` to include stopped containers as well.

currently this formatter only formats those two docker commands, other commands may be added later

## Prerequisites

Before using this tool, make sure you have the following installed:

- **Go**: [Install Go](https://golang.org/doc/install)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/tazbin/docker-formatter.git
cd docker-formatter
```

2. Create the executable
```bash
go build -o docker-formatter main.go
```
3. Place the executable in the `usr/local/bin` directory
```bash
mv docker-formatter /usr/local/bin
```

4. Add this to your `.zshrc` file
```bash
docker() {
    if [ "$1" = "fps" ] && [ "$2" = "-a" ]; then
        /usr/local/bin/docker-formatter docker ps -a
    elif [ "$1" = "fps" ]; then
        /usr/local/bin/docker-formatter docker ps
    else
        command docker "$@"
    fi
}
```

5. Now finally source `.zshrc` file
```bash
source ~/.zshrc
```****

6. Now running `docker fps` command will show output in formatter way
```bash
 # SHOWING ONLY RUNNING CONTAINERS:
·--------------·-------------·----------------------·-------------------·-------------------------·
| CONTAINER ID |    IMAGE    |    CONTAINER NAME    |      STATUS       |          PORTS          |
·--------------·-------------·----------------------·-------------------·-------------------------·
| 157b98bc2e57 | backend-app | nodejs-app-container | Up 7 seconds      | 3000/tcp                |
|              |             |                      |                   |                         |
| 220a34d3a427 | nginx       | nginx-container      | Up 7 seconds      | 0.0.0.0:80->80/tcp,     |
|              |             |                      |                   | 0.0.0.0:4343->4343/tcp, |
|              |             |                      |                   | 0.0.0.0:6000->6000/tcp, |
|              |             |                      |                   | 0.0.0.0:7070->7070/tcp, |
|              |             |                      |                   | 0.0.0.0:9090->9090/tcp  |
|              |             |                      |                   |                         |
| d7eefeabceff | postgres:14 | postgres-container   | Up About a minute | 5432/tcp                |
|              |             |                      |                   |                         |
| 27d6ce5e3b2c | redis       | redis-container      | Up About a minute | 6379/tcp                |
|              |             |                      |                   |                         |
| 67f53bddbf26 | server-app  | app-v2               | Up 3 minutes      | 0.0.0.0:4000->8000/tcp  |
·--------------·-------------·----------------------·-------------------·-------------------------·
```