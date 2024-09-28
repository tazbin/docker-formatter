# Docker Formatter

This Go project is a Docker formatter tool that formats the output of Docker commands such as `docker ps` and `docker ps -a`. The tool enhances readability and makes the output easier to parse and work with.

`docker ps` output is:

```bash
CONTAINER ID   IMAGE         COMMAND                  CREATED       STATUS       PORTS                                                                                                                NAMES
157b98bc2e57   backend-app   "docker-entrypoint.s…"   2 hours ago   Up 2 hours   3000/tcp                                                                                                             nodejs-app-container
220a34d3a427   nginx         "/docker-entrypoint.…"   2 hours ago   Up 2 hours   0.0.0.0:80->80/tcp, 0.0.0.0:4343->4343/tcp, 0.0.0.0:6000->6000/tcp, 0.0.0.0:7070->7070/tcp, 0.0.0.0:9090->9090/tcp   nginx-container
d7eefeabceff   postgres:14   "docker-entrypoint.s…"   2 hours ago   Up 2 hours   5432/tcp                                                                                                             postgres-container
67f53bddbf26   server-app    "docker-entrypoint.s…"   2 hours ago   Up 2 hours   0.0.0.0:4000->8000/tcp                                                                                               app-v2
```

our `docker fps` (formatted ps) output will be:
```bash
 # SHOWING ONLY RUNNING CONTAINERS:
·--------------·-------------·----------------------·------------·-------------------------·
| CONTAINER ID |    IMAGE    |    CONTAINER NAME    |   STATUS   |          PORTS          |
·--------------·-------------·----------------------·------------·-------------------------·
| 157b98bc2e57 | backend-app | nodejs-app-container | Up 2 hours | 3000/tcp                |
|              |             |                      |            |                         |
| 220a34d3a427 | nginx       | nginx-container      | Up 2 hours | 0.0.0.0:80->80/tcp,     |
|              |             |                      |            | 0.0.0.0:4343->4343/tcp, |
|              |             |                      |            | 0.0.0.0:6000->6000/tcp, |
|              |             |                      |            | 0.0.0.0:7070->7070/tcp, |
|              |             |                      |            | 0.0.0.0:9090->9090/tcp  |
|              |             |                      |            |                         |
| d7eefeabceff | postgres:14 | postgres-container   | Up 2 hours | 5432/tcp                |
|              |             |                      |            |                         |
| 67f53bddbf26 | server-app  | app-v2               | Up 2 hours | 0.0.0.0:4000->8000/tcp  |
·--------------·-------------·----------------------·------------·-------------------------·
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

#### 1. Clone the repository:

```bash
git clone https://github.com/tazbin/docker-formatter.git
```

Navigate to the directory
```bash
cd docker-formatter
```

#### 2. Create the executable
```bash
go build -o docker-formatter main.go
```

#### 3. Move the executable in the `usr/local/bin` directory
```bash
mv docker-formatter /usr/local/bin
```

#### 4. Update your shell configuration
You need to add the following snippet to your shell configuration file. Since I'm using **Zsh**, I will add it to the `~/.zshrc` file. 

For other shells, refer to the appropriate configuration file (`~/.bashrc` for **Bash** or `~/.profile` for **sh**).

Open your shell configuration file:
```bash
nano ~/.zshrc
```

Add the following function:
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

#### 5. Apply shell changes
After modifying `.zshrc`, run the following command to apply the changes
```bash
source ~/.zshrc
```

#### 6. Running custom docker commands

Running `docker fps` command will show output in formatter way
```bash
 # SHOWING ONLY RUNNING CONTAINERS:
·--------------·-------------·----------------------·------------·-------------------------·
| CONTAINER ID |    IMAGE    |    CONTAINER NAME    |   STATUS   |          PORTS          |
·--------------·-------------·----------------------·------------·-------------------------·
| 157b98bc2e57 | backend-app | nodejs-app-container | Up 2 hours | 3000/tcp                |
|              |             |                      |            |                         |
| 220a34d3a427 | nginx       | nginx-container      | Up 2 hours | 0.0.0.0:80->80/tcp,     |
|              |             |                      |            | 0.0.0.0:4343->4343/tcp, |
|              |             |                      |            | 0.0.0.0:6000->6000/tcp, |
|              |             |                      |            | 0.0.0.0:7070->7070/tcp, |
|              |             |                      |            | 0.0.0.0:9090->9090/tcp  |
|              |             |                      |            |                         |
| d7eefeabceff | postgres:14 | postgres-container   | Up 2 hours | 5432/tcp                |
|              |             |                      |            |                         |
| 67f53bddbf26 | server-app  | app-v2               | Up 2 hours | 0.0.0.0:4000->8000/tcp  |
·--------------·-------------·----------------------·------------·-------------------------·
```

Running `docker fps -a` command will show output in formatter way
```bash
 # SHOWING ALL CONTAINERS (RUNNING, STOPPED, EXITED etc):
·--------------·------------------------·----------------------·--------------------------·-------------------------·
| CONTAINER ID |         IMAGE          |    CONTAINER NAME    |          STATUS          |          PORTS          |
·--------------·------------------------·----------------------·--------------------------·-------------------------·
| 157b98bc2e57 | backend-app            | nodejs-app-container | Up 2 hours               | 3000/tcp                |
|              |                        |                      |                          |                         |
| 220a34d3a427 | nginx                  | nginx-container      | Up 2 hours               | 0.0.0.0:80->80/tcp,     |
|              |                        |                      |                          | 0.0.0.0:4343->4343/tcp, |
|              |                        |                      |                          | 0.0.0.0:6000->6000/tcp, |
|              |                        |                      |                          | 0.0.0.0:7070->7070/tcp, |
|              |                        |                      |                          | 0.0.0.0:9090->9090/tcp  |
|              |                        |                      |                          |                         |
| d7eefeabceff | postgres:14            | postgres-container   | Up 2 hours               | 5432/tcp                |
|              |                        |                      |                          |                         |
| 27d6ce5e3b2c | redis                  | redis-container      | Exited (0) 4 minutes ago |                         |
|              |                        |                      |                          |                         |
| 67f53bddbf26 | server-app             | app-v2               | Up 2 hours               | 0.0.0.0:4000->8000/tcp  |
·--------------·------------------------·----------------------·--------------------------·-------------------------·
```