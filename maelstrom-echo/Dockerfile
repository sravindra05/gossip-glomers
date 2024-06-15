FROM golang:1.22

WORKDIR /maelstrom-echo

COPY deps.mk ./
COPY ./maelstrom-echo/go.mod ./maelstrom-echo/go.sum ./maelstrom-echo/Makefile ./

RUN go mod download

# get maelstrom dependencies
RUN make deps

COPY ./maelstrom-echo .

# install echo binary
RUN go install .

# run test
CMD ["make", "run-echo"]

