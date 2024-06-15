.PHONY:
build-echo:
	docker build -t echo -f maelstrom-echo/Dockerfile .

.PHONY:
run-echo: build-echo
	docker run echo

.PHONY:
build-uid:
	docker build -t uid -f uid-gen/Dockerfile .

.PHONY:
run-uid: build-uid
	docker run uid

.PHONY:
build-broadcast:
	docker build -t broadcast -f broadcast/Dockerfile .


NODE_COUNT ?= 1
.PHONY:
run-broadcast: build-broadcast
	docker run -e NODE_COUNT=$(NODE_COUNT) broadcast
