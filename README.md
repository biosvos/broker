## 구조

publiser -> subscriber 로 직접 전달한다.

같은 topic에 연결된 subscriber는 동시에 같은 delivery를 받는다.

### broker

- broker에 인증/연결을 담당한다.

### publisher

- message publish 만을 담당한다.

### subscriber

- delivery subscribe 만을 담당한다.