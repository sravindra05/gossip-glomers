.PHONY: deps run-echo

MAELSTROM_BIN_PATH ?= ./maelstrom/maelstrom

deps:
	apt update && apt install -y openjdk-17-jdk gnuplot graphviz bzip2
	wget -O maelstrom.tar.bz2 https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2
	tar -xvf maelstrom.tar.bz2