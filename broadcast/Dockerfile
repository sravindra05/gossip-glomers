FROM golang:1.22

WORKDIR /broadcast

COPY deps.mk ./
COPY ./broadcast/go.mod ./broadcast/go.sum ./

RUN go mod download

COPY ./broadcast .

# get maelstrom dependencies
RUN make deps

# make node count configurable for broadcast
ARG NODE_COUNT
ENV NODE_COUNT=${NODE_COUNT}
# install binary
RUN go install .

# run test
CMD ["make", "run-broadcast"]

