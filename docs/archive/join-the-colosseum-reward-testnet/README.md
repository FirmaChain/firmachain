# Join the Colosseum (reward-testnet)

### Colosseum testnet has officially been terminated as of December 29, 2021.



#### Have you completed the previous step?

* If you are unsure about the minimum specifications or the ‘firmachaind’ binary files, please visit the following link. [Install FirmaChain](../../getting-started/install-firmachain/)&#x20;

## How to setup

Please follow the directions below in order to become a validator.

### Initialize chain node

```
firmachaind init <node_name> --chain-id colosseum-1
```

### Change gas prices

```
vi ~/.firmachain/config/app.toml

###############################################################################
###                           Base Configuration                            ###
###############################################################################

minimum-gas-prices = "0.1ufct"
```

#### Support RPC-API

REST API endpoint is not activated in the default settings. Therefore, if you want to support, please follow the directions below.

#### This process is not mandatory.

```
vi ~/.firmachain/config/app.toml
```

```
###############################################################################
###                           API Configuration                             ###
###############################################################################

[api]

# Enable defines if the API server should be enabled.
enable = false
```

### Restore Wallet

#### Restore Key (Validators must have mnemonic)&#x20;

```
firmachaind keys add <key_name> --recover --coin-type 7777777
> Enter your bip39 mnemonic
glad music grace lawn squeeze book very text wire okay ozone morning permit tumble guard hurry various gallery kitten surprise brain piano level picture

Enter keyring passphrase: XXXXXXXX
Re-enter keyring passphrase: XXXXXXXX

- name: firmatester
  type: local
  address: 
  pubkey: ''
  mnemonic: ""
```

### genesis.json file setting

#### Download genesis.json and copy to firmachain config folder

genesis.json download

```
wget https://github.com/firmachain/firmachain/v05-testnet-colosseum/raw/main/genesis.json
```

copy genesis.json from firmachain config folder

```
cp ./genesis.json ~/.firmachain/config/genesis.json
```

#### P2P Setting

We provide three seed nodes for P2P connection. You can find more information from this link.

```
vi ~/.firmachain/config/config.toml

#######################################################
###           P2P Configuration Options             ###
#######################################################
[p2p]

# Address to listen for incoming connections
laddr = "tcp://0.0.0.0:26656"

# Comma separated list of seed nodes to connect to
seeds = "seed list"
```

#### Seed list (copy seeds)

```
1a8e340bf37d7a419b7b5a4db0f323099a060443@colosseum-seed1.firmachain.dev:26656,3e8c571232bdd6b48073213476173fd46b1c8293@colosseum-seed2.firmachain.dev:26656,458c78173fd416e91fed08c215cc935556d25414@colosseum-seed3.firmachain.dev:26656
```

### Start Colosseum network

```
firmachaind start
7:42AM INF starting ABCI with Tendermint
7:42AM INF Starting multiAppConn service impl=multiAppConn module=proxy
7:42AM INF Starting localClient service connection=query impl=localClient module=abci-client
7:42AM INF Starting localClient service connection=snapshot impl=localClient module=abci-client
7:42AM INF Starting localClient service connection=mempool impl=localClient module=abci-client
7:42AM INF Starting localClient service connection=consensus impl=localClient module=abci-client
7:42AM INF Starting EventBus service impl=EventBus module=events
7:42AM INF Starting PubSub service impl=PubSub module=pubsub
7:42AM INF Starting IndexerService service impl=IndexerService module=txindex
7:42AM INF ABCI Handshake App Info hash= height=0 module=consensus protocol-version=0 software-version=0.2.8
7:42AM INF ABCI Replay Blocks appHeight=0 module=consensus stateHeight=0 storeHeight=0
7:42AM INF asserting crisis invariants inv=0/11 module=x/crisis name=staking/module-accounts
7:42AM INF asserting crisis invariants inv=1/11 module=x/crisis name=staking/nonnegative-power
7:42AM INF asserting crisis invariants inv=2/11 module=x/crisis name=staking/positive-delegation
7:42AM INF asserting crisis invariants inv=3/11 module=x/crisis name=staking/delegator-shares
7:42AM INF asserting crisis invariants inv=4/11 module=x/crisis name=bank/nonnegative-outstanding
7:42AM INF asserting crisis invariants inv=5/11 module=x/crisis name=bank/total-supply
7:42AM INF asserting crisis invariants inv=6/11 module=x/crisis name=distribution/nonnegative-outstanding
7:42AM INF asserting crisis invariants inv=7/11 module=x/crisis name=distribution/can-withdraw
7:42AM INF asserting crisis invariants inv=8/11 module=x/crisis name=distribution/reference-count
7:42AM INF asserting crisis invariants inv=9/11 module=x/crisis name=distribution/module-account
7:42AM INF asserting crisis invariants inv=10/11 module=x/crisis name=gov/module-account
7:42AM INF asserted all invariants duration=1.309414 height=0 module=x/crisis
7:42AM INF created new capability module=ibc name=ports/transfer
7:42AM INF port binded module=x/ibc/port port=transfer
7:42AM INF claimed capability capability=1 module=transfer name=ports/transfer
7:42AM INF Completed ABCI Handshake - Tendermint and App are synced appHash= appHeight=0 module=consensus
7:42AM INF Version info block=11 p2p=8 tendermint_version=0.34.12
7:42AM INF This node is a validator addr=A442D1FD6E2BD4C03101680E3815A7664930244E module=consensus pubKey=oEJ7OgjBwvPnw1PGQYb646/Dc/B0pNGVMO2VL+zm/Cs=
7:42AM INF P2P Node ID ID=09835f25a665ff508cefac029906ad52686a8612 file=/home/firma/.firmachain/config/node_key.json module=p2p
7:42AM INF Adding persistent peers addrs=[] module=p2p
7:42AM INF Adding unconditional peer ids ids=[] module=p2p
7:42AM INF Add our address to book addr={"id":"09835f25a665ff508cefac029906ad52686a8612","ip":"0.0.0.0","port":26656} book=/home/firma/.firmachain/config/addrbook.json module=p2p
7:42AM INF Starting Node service impl=Node
7:42AM INF Starting pprof server laddr=localhost:6060
7:42AM INF Starting P2P Switch service impl="P2P Switch" module=p2p
7:42AM INF Starting Mempool service impl=Mempool module=mempool
7:42AM INF Starting BlockchainReactor service impl=BlockchainReactor module=blockchain
7:42AM INF Starting RPC HTTP server on 127.0.0.1:26657 module=rpc-server
7:42AM INF Starting Consensus service impl=ConsensusReactor module=consensus
7:42AM INF Reactor  module=consensus waitSync=false
7:42AM INF Starting State service impl=ConsensusState module=consensus
7:42AM INF Starting baseWAL service impl=baseWAL module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
7:42AM INF Starting Group service impl=Group module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
7:42AM INF Searching for height height=1 max=0 min=0 module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
7:42AM INF Searching for height height=0 max=0 min=0 module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
7:42AM INF Found height=0 index=0 module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
7:42AM INF Catchup by replaying consensus messages height=1 module=consensus
7:42AM INF Replay: Done module=consensus
7:42AM INF Starting TimeoutTicker service impl=TimeoutTicker module=consensus
7:42AM INF Starting Evidence service impl=Evidence module=evidence
7:42AM INF Starting StateSync service impl=StateSync module=statesync
7:42AM INF Starting PEX service impl=PEX module=pex
7:42AM INF Starting AddrBook service book=/home/firma/.firmachain/config/addrbook.json impl=AddrBook module=p2p
7:42AM INF Saving AddrBook to file book=/home/firma/.firmachain/config/addrbook.json module=p2p size=0
7:42AM INF Ensure peers module=pex numDialing=0 numInPeers=0 numOutPeers=0 numToDial=10
7:42AM INF No addresses to dial. Falling back to seeds module=pex
7:42AM INF Timed out dur=4997.41631 height=1 module=consensus round=0 step=1
7:42AM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"AD43A0590BACD05761A5672B941BE46770D107F8D780CC72E3C34932F93A4D44","parts":{"hash":"9D8CB5FE3316747098D358D37284AE052F29F252137ECBDAB12BF542D7AC146F","total":1}},"height":1,"pol_round":-1,"round":0,"signature":"EmQrypojuR8EmThxgvSQIbGOZIX1+lbrb6xTqC08/jYHdqDX5/2amQ+NdD+i+kY0o7hg0k+5LDKKN3CAupnrDQ==","timestamp":"2021-11-24T07:42:15.203109231Z"}
7:42AM INF received complete proposal block hash=AD43A0590BACD05761A5672B941BE46770D107F8D780CC72E3C34932F93A4D44 height=1 module=consensus
7:42AM INF finalizing commit of block hash=AD43A0590BACD05761A5672B941BE46770D107F8D780CC72E3C34932F93A4D44 height=1 module=consensus num_txs=0 root=E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855
7:42AM INF minted coins from module account amount=2059725ufct from=mint module=x/bank
7:42AM INF executed block height=1 module=state num_invalid_txs=0 num_valid_txs=0
7:42AM INF commit synced commit=436F6D6D697449447B5B31303020313839203230352031363120313438203130322032303820373820323420313320363420323330203139382031393620313035203139302031333620323534203135322037382031303320323237203135203230312032333120323132203138362031313020313720313231203739203138325D3A317D
7:42AM INF committed state app_hash=64BDCDA19466D04E180D40E6C6C469BE88FE984E67E30FC9E7D4BA6E11794FB6 height=1 module=state num_txs=0
7:42AM INF indexed block height=1 module=txindex
```

### Register as Validator

```
firmachaind tx staking create-validator \
--pubkey $(firmachaind tendermint show-validator) \
--moniker <Your moniker name> \
--chain-id colosseum-1 \
--commission-rate 0.10 \
--commission-max-rate 0.20 \
--commission-max-change-rate 0.01 \
--min-self-delegation 1 \
--amount <staking amount>ufct \
--fees 20000ufct \
--from <key_name>
```
