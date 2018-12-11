#!/usr/bin/env bash
#
docker rm -f object4d-minio
docker run --name object4d-minio -p 9000:9000 -d minio/minio server /data/minio
docker rm -f object4d-redis
docker run  --name object4d-redis --net=bridge --restart=always -p 36379:6379 -d redis

docker ps -a|grep object4d