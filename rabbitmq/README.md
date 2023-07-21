# 참고
https://github.com/streadway/amqp 는 deprecated 되었다.
대신, https://github.com/rabbitmq/amqp091-go 를 사용해야 한다.

# 구조

기본 구조는 다음과 같다.  
publisher -> exchange -> route -> queue -> consumer

즉, publisher는 메시지를 exchange를 거쳐 큐로 보내고,
consumer는 queue에 직접적으로 연결되어 메시지를 가져온다.

## 기본 용어

- routing key
    - A.B.C 와 같이 . 으로 구분되는 키이다.
    - exchange와 queue를 연결한다.

## exchange type

- direct
- fanout
- topic
- headers

### direct

routing key 기반 메시지 전달
routing key가 정확히 일치해야 한다.

### fanout

브로드 캐스팅 방식
routing 되어있는 모든 큐에 메시지 전달(routing key 무시)

### topic

routing key 패턴 기반 메시지 전달

routing key 패턴

- *: 한 단어
- #: 이후 모든 단어

### headers

routing key + header 기반 메시지 전달

# amqp library 와의 rabbitmq 구조 매핑
```
ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table) error
QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
QueueBind(name, key, exchange string, noWait bool, args Table) error
```
