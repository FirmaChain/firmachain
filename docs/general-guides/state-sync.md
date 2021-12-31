# State Sync

## Instruction

‘[State Sync](https://docs.tendermint.com/v0.34/tendermint-core/state-sync.html#state-sync)’ feature was included in the recent release of the tendermint core v0.34 update.\
The feature allows users to look up information on the most recent credible block height instead of having to look up all block information from the past. The ‘state sync’ feature reduces the time required to synchronize the network to a matter of minutes.

## State syncing a node

FirmaChain is structured to create a status sync snapshot of a portion of its nodes to allow a new node to join the network using the ‘State Sync’ feature. In order to join the network through ‘State Sync’ the following information must be retrieved in advance.

* At least 2 available RPC servers.
* A trusted height.
* The block ID hash of the trusted height.

If you don't have a jq package, please install the jq package.

```
sudo apt install jq -y
```

By running the command below, you will be able to retrieve data on the ‘trust height’ and ‘trust\_hash’ via RPC.

```
curl -s http://colosseum-seed2.firmachain.dev:26657/block | \
jq -r '.result.block.header.height + "\n" + .result.block_id.hash'
```

Following is a sample of the retrieved data.

```
476569
2DC71A2F41B2E74718AC50ED33960902C75282EED5AD56F18B4BFCB33F9B796A
```

Please set the chain to allow ‘State Sync’ feature and put in the retrieved data in the config.toml file.

```
#######################################################
###         State Sync Configuration Options        ###
#######################################################
[statesync]
...
# starting from the height of the snapshot.
enable = true
...
# weeks) during which they can be financially punished (slashed) for misbehavior.
rpc_servers = "192.168.20.108:26657,192.168.20.109:26657"
trust_height = 476569
trust_hash = "2DC71A2F41B2E74718AC50ED33960902C75282EED5AD56F18B4BFCB33F9B796A"
trust_period = "168h0m0s"
```

Now we are all set. Please start the chain. \
Once you start the chain, it will automatically search and restore the status sync snapshot from the network. (This process takes approximately 2 to 5 minutes.)

```
firmachaind start

Discovering snapshots for 15s
Discovered new snapshot        height=3000 format=1 hash=0F14A473
Discovered new snapshot        height=2000 format=1 hash=C6209AF7
```
