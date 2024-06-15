FROM golang:1.22

WORKDIR /uid-gen

COPY deps.mk ./
COPY ./uid-gen/go.mod ./uid-gen/go.sum ./uid-gen/Makefile ./

RUN go mod download

# get maelstrom dependencies
RUN make deps

COPY ./uid-gen .

# install binary
RUN go install .

# run test
CMD ["make", "run-uid-gen"]

