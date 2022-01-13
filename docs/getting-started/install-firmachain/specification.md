# Specification

### Hardware requirements

Disk space depends on the [pruning strategy](https://hub.cosmos.network/main/hub-tutorials/join-mainnet.html#pruning-of-state) of your choice.

There are four strategies for pruning state. These strategies apply only to state and do not apply to block storage. To set pruning, adjust the `pruning` parameter in the `~/.firmachain/config/app.toml` file. The following pruning state settings are available:

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

### Increase maximum open files

In general, most operating systems limit the maximum number of files that can be opened.\
(The document was written based on `Ubuntu 20.04 LTS`, and the default value of this OS is 1024.)\
However, if you open `/etc/security/limits.conf` and modify the value as can be seen in the code below, you can raise the limit of the nofile feature

```
*                soft    nofile          65535
*                hard    nofile          65535
```
