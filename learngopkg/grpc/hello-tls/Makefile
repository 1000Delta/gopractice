# 制作私钥
server.key:
	openssl genrsa -out server.key 2048

# 生成公钥
server.pem:
	openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

.PHONY: init
init: server.key server.pem

server/helloserver:
	cd server && go build -o helloserver .

client/helloclient:
	cd client && go build -o helloclient .

.PHONY: server
server: server/helloserver
	cd server && (./helloserver >>server.log &)

.PHONY: client
client: client/helloclient
	cd client && (./helloclient >>client.log &)

.PHONY: stop
stop:
	kill $$(ps | grep -P "hello(client|server)" | awk '{print $$1}')

.PHONY: clean
.IGNORE: clean
clean:
	rm server/helloserver
	rm client/helloclient