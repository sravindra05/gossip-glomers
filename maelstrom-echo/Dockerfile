FROM golang:1.22

WORKDIR /maelstrom-echo

COPY deps.mk ./
COPY ./maelstrom-echo/go.mod ./maelstrom-echo/go.sum ./

RUN go mod download

COPY ./maelstrom-echo .

# get maelstrom dependencies
RUN make deps

# install binary
RUN go install .

# run test
CMD ["make", "run-echo"]

