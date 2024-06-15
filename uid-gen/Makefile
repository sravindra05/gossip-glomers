include deps.mk

.PHONY:
run-uid-gen:
	$(MAELSTROM_BIN_PATH) test -w unique-ids --bin $(GOPATH)/bin/uid-gen --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition