# how to run
# docker build -t firmachain .
# docker run -p 1317:1317 -p 26657:26657  firmachain

# Use multi-stage build
FROM golang:1.16 as builder

RUN apt-get update && apt-get install -y git wget

# Download from GitHub instead of using COPY
RUN rm firmachain -rf
RUN git clone https://github.com/firmachain/firmachain/v05 /firmachain
WORKDIR "/firmachain"

# Checkout a specific version
RUN git checkout v0.2.3
RUN make 
#RUN LEDGER_ENABLED=false make 

RUN wget https://github.com/firmachain/firmachain/v05-devnet-imperium/raw/master/genesis.json;

# Create final container
FROM ubuntu:latest

# It is ok to COPY files from a build container (when using multi-stage builds)
COPY --from=builder /go/bin/firma /usr/local/bin/firmachaind

# rest server / grpc / tendermint p2p / tendermint rpc
EXPOSE 1317 9090 26656 26657

RUN firmachaind unsafe-reset-all
RUN firmachaind init --chain-id imperium-2 moniker

COPY --from=builder /firmachain/genesis.json root/.firmachain/config/genesis.json

RUN sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "1.0ufct"/g' ~/.firmachain/config/app.toml
RUN sed -i 's/persistent_peers = ""/persistent_peers = "65db86dfd58c1e3732cfd81bf2129bba222b2cc7@192.168.20.101:26656"/g' ~/.firmachain/config/config.toml
RUN sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' ~/.firmachain/config/config.toml
RUN sed -i ':a;N;$!ba;s/# Enable defines if the API server should be enabled.\nenable = false/# Enable defines if the API server should be enabled.\nenable = true/g' ~/.firmachain/config/app.toml

CMD ["/usr/local/bin/firmachaind", "start"]