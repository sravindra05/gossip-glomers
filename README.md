# gossip-glomers
A repository containing my solutions to the distributed systems challenge by [fly.io](https://fly.io/dist-sys/)

# Try it out
Each of the challenges have an associated make target that builds and runs a docker image for the challenge in question so I don't get to claim it runs only on my machine :) However, it goes without saying that `make` and `docker` are prerequisites.

### Challenge #1 Echo
```shell
make run-echo
```

### Challenge #2 UID Generation
```shell
make run-uid
```

### Challenge #3 Broadcast
This solution covers both a single node and multi-node broadcast so node count can be passed as an arg like so.
```shell
NODE_COUNT=1 make run-broadcast
```


