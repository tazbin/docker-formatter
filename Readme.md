# Docker Formatter

This Go project is a Docker formatter tool that formats the output of Docker commands such as `docker ps` and `docker ps -a`. The tool enhances readability and makes the output easier to parse and work with.

`docker ps` output is:

```bash
CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS          PORTS                NAMES
9a7c666f88a1   backend-app   "docker-entrypoint.s…"   34 seconds ago   Up 32 seconds   3000/tcp             nodejs-app-container
800c5893e5dd   nginx         "/docker-entrypoint.…"   34 seconds ago   Up 32 seconds   0.0.0.0:80->80/tcp   nginx-container
3e8f1d3736e6   postgres:14   "docker-entrypoint.s…"   34 seconds ago   Up 33 seconds   5432/tcp             postgres-container
dcac680aa299   redis         "docker-entrypoint.s…"   34 seconds ago   Up 33 seconds   6379/tcp             redis-container
```

our `docker fps` (formatted ps) output will be:
```bash
 # SHOWING ONLY RUNNING CONTAINERS:
·--------------·-------------·----------------------·--------------·--------------------·
| CONTAINER ID |    IMAGE    |    CONTAINER NAME    |    STATUS    |       PORTS        |
·--------------·-------------·----------------------·--------------·--------------------·
| 9a7c666f88a1 | backend-app | nodejs-app-container | Up 2 seconds | 3000/tcp           |
|              |             |                      |              |                    |
| 800c5893e5dd | nginx       | nginx-container      | Up 1 second  | 0.0.0.0:80->80/tcp |
|              |             |                      |              |                    |
| 3e8f1d3736e6 | postgres:14 | postgres-container   | Up 2 seconds | 5432/tcp           |
|              |             |                      |              |                    |
| dcac680aa299 | redis       | redis-container      | Up 2 seconds | 6379/tcp           |
|              |             |                      |              |                    |
·--------------·-------------·----------------------·--------------·--------------------·
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
git clone <https://github.com/tazbin/docker-formatter.git>
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
```

6. Now running `docker fps` command will show output in formatter way
```bash
 # SHOWING ONLY RUNNING CONTAINERS:
·--------------·-------------·----------------------·--------------·--------------------·
| CONTAINER ID |    IMAGE    |    CONTAINER NAME    |    STATUS    |       PORTS        |
·--------------·-------------·----------------------·--------------·--------------------·
| 9a7c666f88a1 | backend-app | nodejs-app-container | Up 2 seconds | 3000/tcp           |
|              |             |                      |              |                    |
| 800c5893e5dd | nginx       | nginx-container      | Up 1 second  | 0.0.0.0:80->80/tcp |
|              |             |                      |              |                    |
| 3e8f1d3736e6 | postgres:14 | postgres-container   | Up 2 seconds | 5432/tcp           |
|              |             |                      |              |                    |
| dcac680aa299 | redis       | redis-container      | Up 2 seconds | 6379/tcp           |
|              |             |                      |              |                    |
·--------------·-------------·----------------------·--------------·--------------------·
```