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