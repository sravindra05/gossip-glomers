FROM golang:1.22

WORKDIR /uid-gen

COPY deps.mk ./
COPY ./uid-gen/go.mod ./uid-gen/go.sum ./

RUN go mod download

COPY ./uid-gen .

# get maelstrom dependencies
RUN make deps

# install binary
RUN go install .

# run test
CMD ["make", "run-uid-gen"]

