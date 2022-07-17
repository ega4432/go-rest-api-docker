# go-rest-api-docker

## Usage

```shell
$ git clone git@github.com:ega4432/go-rest-api-docker.git && cd go-rest-api-docker

$ cp .env.example .env

$ docker compose up --build
```

## Tips

### Connect db container

```shell
$ docker compose exec -it db /bin/bash -c "mysql -uroot -p<PASSWORD>"
```

### Import API test file with Thunder client

1. Install extension
2. Import JSON file
