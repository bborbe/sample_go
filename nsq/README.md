# NSQ

## nsqd

```
docker rm nsqd
docker run \
--name nsqd \
-p 4150:4150 \
-p 4151:4151 \
docker.io/bborbe/nsq:latest /nsqd
```
