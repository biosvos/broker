# 참고
https://github.com/streadway/amqp 는 deprecated 되었다.
대신, https://github.com/rabbitmq/amqp091-go 를 사용해야 한다.

# deploy rabbitmq in podman
```bash
podman run \
-d \
--name rabbitmq \
-p 5672:5672 \
-p 8080:8080 \
-e RABBITMQ_DEFAULT_USER=guest \
-e RABBITMQ_DEFAULT_PASS=guest \
docker.io/rabbitmq:management
```