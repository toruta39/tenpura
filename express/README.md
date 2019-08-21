# my-express-app

## Usage

```
docker-compose build
docker-compose up
```

After every time `package-lock.json` is updated, there's a need to rebuild 
the docker image.

The container will expose the server at port 3000, modify 
`docker-compose.yaml` to switch to another port.

## TODO

- logging
- build
- test
- deploy
- production readiness
- ci/cd
