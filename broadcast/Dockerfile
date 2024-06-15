FROM golang:1.22

WORKDIR /broadcast

COPY deps.mk ./
COPY ./broadcast/go.mod ./broadcast/go.sum ./broadcast/Makefile ./

RUN go mod download

# get maelstrom dependencies
RUN make deps

COPY ./broadcast .

# make node count configurable for broadcast
ENV NODE_COUNT=1

# install binary
RUN go install .

# run test
CMD ["make", "run-broadcast"]

