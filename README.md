# connect-orgs

## Compiling

```
$ make clean && make
```

## Building Docker image

```
$ make docker-build
```

## Pushing Docker image to Container Registry

Note that only the owner of the repository is allowed to push the image. 

```
$ make docker-push
```

## Running

### For local environment

#### Run connect-infra
```
$ git clone https://github.com/mrexmelle/connect-infra
$ cd connect-infra
$ docker compose up
```

#### Run local service
```
$ ./connect-orgs serve
```

### For docker environment

#### Run connect-infra
```
$ git clone https://github.com/mrexmelle/connect-infra
$ cd connect-infra
$ docker compose up
```

#### Run service in docker
```
$ make docker-build
$ docker compose up
```
Note that you cannot alter the docker image in the container registry. Only the owner of the repository is allowed to do so.

If error happens in `core` service due to failure to connect to database, restart it:
```
$ docker compose restart core
```
The failure might happen due to database service isn't ready when `core` attempts to connect to it.


## API Documentation
Once the service runs, the API documentation is available in `$HOST:$PORT/swagger/index.html`

Note that the API documentation is only available if the service is run with `local` profile, i.e. when either `APP_PROFILE` environment is defined as `local` or undefined.