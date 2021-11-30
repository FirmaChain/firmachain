# Install FirmaChain

## Specification

### Hardware requirements

Disk space depends on the [pruning strategy](https://hub.cosmos.network/main/gaia-tutorials/join-mainnet.html#pruning-of-state) of your choice.

There are four strategies for pruning state. These strategies apply only to state and do not apply to block storage. To set pruning, adjust the `pruning` parameter in the `~/.gaia/config/app.toml` file. The following pruning state settings are available:

1. `everything`: Prune all saved states other than the current state.
2. `nothing`: Save all states and delete nothing.
3. `default`: Save the last 100 states and the state of every 10,000th block.
4. `custom`: Specify pruning settings with the `pruning-keep-recent`, `pruning-keep-every`, and `pruning-interval` parameters.

| Pruning strategy	 | Minimum disk space | Recommended disk space |
| ----------------- | ------------------ | ---------------------- |
| everything        | 20 GB              | 40 GB                  |
| default           | 80 GB              | 120 GB                 |
| nothing           | 120 GB             | > 240 GB               |

By default, every node is in `default` mode which is the recommended setting for most environments.



Apart from disk space, the following requirements should be met.

| Minimum CPU cores | Recommended CPU cores |
| ----------------- | --------------------- |
| 2                 | 4                     |

| Minimum RAM | Recommended RAM |
| ----------- | --------------- |
| 4 GB        | 8 GB            |

### Service Ports

| Type | Port number | Description            |
| ---- | ----------- | ---------------------- |
| SSH  | 22          | for connect ssh        |
| API  | 1317        | http api               |
| P2P  | 26656       | for peer communication |

## Installation

### Ways 1. Install FirmaChain from pre-built binary by curl

You can easily download a pre-built binary for your operating system using the command below.

#### download binary

```shell
# On Bash
curl https://build.firmachain.org/@v0.3.1 | bash

sudo mv ./firmachaind /usr/local/bin/firmachaind
```

#### check version & sha1sum

```
firmachaind version
0.3.1
```

```
firmachaind version --long
name: Firmachain
server_name: firmachaind
version: 0.3.1
commit: 987f4c165213aa7e928683b1c7128834dc12e460
...
cosmos_sdk_version: v0.44.3
```

```
sha1sum /usr/local/bin/firmachaind
ed0faa6537bf7ecf63febe334abf980ca199307e  firmachaind
```

### Ways 2. Download from Github Release page

You can download a prebuilt binary from the link below. [https://github.com/FirmaChain/firmachain/releases/tag/v0.3.1](https://github.com/FirmaChain/firmachain/releases/tag/v0.3.1)​

![Screenshot from Github release page](https://firmachain.gitbook.io/\~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MizWtAhqgIPKt343EjZ%2Fuploads%2Fi3jIsz9TLJOFKG8Rax6y%2Fstep2\_download\_image.png?alt=media\&token=dee95f51-021b-4c12-9318-c2e3c4247703)

Select, download and unzip a binary that suits your OS.\
Please confirm whether you’ve downloaded the correct version, using the method provided in Ways 1.
