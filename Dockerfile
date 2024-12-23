# how to run
# docker build -t firmachain .

# docker run -it -p 26657:26657 -p 26656:26656 -v ~/.firmachain:/root/.firmachain firmachain firmachaind init
# docker run -it -p 26657:26657 -p 26656:26656 -v ~/.firmachain:/root/.firmachain firmachain firmachaind start

# to enter docker
# docker exec -it firmachain bash

# run container as a daemon (td option : -t -> Assign Termail to Container, -d: run on the background)
# > docker run -td -p 26657:26657 -p 26656:26656 -v ~/.firmachain:/root/.firmachain firmachain firmachaind start

# Use multi-stage build
FROM golang:1.16 as builder

RUN apt-get update && apt-get install -y git

# Download from GitHub instead of using COPY
RUN rm firmachain -rf
RUN git clone https://github.com/firmachain/firmachain/v05 /firmachain
WORKDIR "/firmachain"

# Always run on latest version
RUN LEDGER_ENABLED=false make 

# Create final container
FROM ubuntu:latest

# It is ok to COPY files from a build container (when using multi-stage builds)
COPY --from=builder /go/bin/firmachaind /usr/local/bin/firmachaind

# rest server / grpc / tendermint p2p / tendermint rpc
EXPOSE 1317 9090 26656 26657

# Run firmachind by default 
# ex) docker run firmachain
CMD ["/usr/local/bin/firmachaind"]
