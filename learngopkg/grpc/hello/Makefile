# 制作私钥
server.key:
	openssl genrsa -out server.key 2048

# 生成公钥
server.pem:
	openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

.PHONY: init
init: server.key server.pem

.PHONY: server
server:
	cd server && (go run . >>server.log &)

.PHONY: client
client:
	cd client && (go run . >>client.log &)

.PHONY: stop
stop:
	kill $$(ps | grep -P "hello(client|server)" | awk '{print $$1}')