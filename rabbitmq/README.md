# 구조
기본 구조는 다음과 같다.  
publisher -> exchange -> route -> queue -> consumer 

즉, publisher는 메시지를 exchange를 거쳐 큐로 보내고,
consumer는 queue에 직접적으로 연결되어 메시지를 가져온다.