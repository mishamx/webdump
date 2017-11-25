WebDump
=======

Developer tool for debug incoming http request

Install and run
---------------

### Download


```bash
curl -fsSL https://raw.github.com/mishamx/webdump/master/webdump
```

### Run

```bash
./webdump
```

### Env

Change webdump settings from ENV

```bash
export WEBDUMP_LISTEN=127.0.0.1:80
export WEBDUMP_TESTING_PATH: "/my/test/path"
export WEBDUMP_FILE_HEADER: "header.txt"
export WEBDUMP_FILE_FULL: "fullrequest.txt"
```

Use WebDump and ngrok

```bash
./webdump
# another terminal
ngrok http 3001
```

### Run from Docker

Create `docker-compose.yml` file
```yaml
version: '3.3'

services:
    webdump:
        image: mishamx/webdump:0.1
        environment:
            WEBDUMP_LISTEN: ":3001"
            WEBDUMP_TESTING_PATH: "/my/test/path"
            WEBDUMP_FILE_HEADER: "header.txt"
            WEBDUMP_FILE_FULL: "fullrequest.txt"
        ports:
            - "3001:3001"
        networks:
            - frontend
    ngrok:
        image: wernight/ngrok
        links:
            - "webdump"
        ports:
            - "4040:4040"
        networks:
            - frontend
        command: ["/bin/ngrok", "http", "webdump:3001"]

networks:
    frontend:
```

Run docker compose:
```bash
dockser-compose up
```

