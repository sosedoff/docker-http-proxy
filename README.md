# docker-http-proxy

Auxilary docker container to proxy HTTP requests to a backend

## Usage

```
docker run \
  -d \
  -e TARGET=https://some-backend-service.com
  --restart=always \
  sosedoff/docker-http-proxy
```