include deps.mk

NODE_COUNT ?= 1

.PHONY:
run-broadcast:
	$(MAELSTROM_BIN_PATH) test -w broadcast --bin $(GOPATH)/bin/broadcast --node-count $(NODE_COUNT) --time-limit 20 --rate 10