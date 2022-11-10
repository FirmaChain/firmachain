# Cosmovisor Example

The purpose of this document is to inform the users on how to upgrade the chain to v0.3.2 using **`cosmovisor`**. All processes mentioned in this document must be conducted on the binary v0.3.1 preconfigured chain server.

### Cosmovisor download & setup

#### 1. Download Cosmovisor

```
go get github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
```

**2. Using the command below**

```
sudo cp cosmovisor /usr/local/bin/cosmovisor
```

#### 3. Setup environment

```
sudo nano ~/.profile

export DAEMON_HOME=$HOME/.firmachain
export DAEMON_NAME=firmachaind
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true
```

Once you've applied all environment variables, please reload the .profile file by running the command provided below.

```
source ~/.profile
```

If all environment variables have been registered successfully, you can input the command provided below to get an output on the registered variables.

```
echo $DAEMON_NAME
```



:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

_<mark style="color:red;">**※ In case of 1.0.0 version.**</mark>_

* Please execute the command below.\
  If you don't execute this command, you can't run **`cosmovisor`**.

```
echo "{}" >> /home/firma/.firmachain/cosmovisor/current/upgrade-info.json
```

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::



In order to launch **`cosmovisor`**, please enter the command provided below.

```
$ cosmovisor start
12:55PM INF Configuration is valid:
Configurable Values:
DAEMON_HOME: /home/firma/.firmachain
DAEMON_NAME: firmachaind
DAEMON_ALLOW_DOWNLOAD_BINARIES: false
DAEMON_RESTART_AFTER_UPGRADE: true
DAEMON_POLL_INTERVAL: 300ms
UNSAFE_SKIP_BACKUP: false
DAEMON_PREUPGRADE_MAX_RETRIES: 0
Derived Values:
Root Dir: /home/firma/.firmachain/cosmovisor
Upgrade Dir: /home/firma/.firmachain/cosmovisor/upgrades
Genesis Bin: /home/firma/.firmachain/cosmovisor/genesis/bin/firmachaind
  Monitored File: /home/firmauser/.firmachain/data/upgrade-info.json
module=cosmovisor
12:55PM INF running app args=["start"]
module=cosmovisor path=/home/firmauser/.firmachain/cosmovisor/upgrades/v0.3.2/bin/firmachaind
12:55PM INF starting ABCI with Tendermint
12:55PM INF Starting multiAppConn service impl=multiAppConn module=proxy
12:55PM INF Starting localClient service connection=query impl=localClient module=abci-client
12:55PM INF Starting localClient service connection=snapshot impl=localClient module=abci-client
12:55PM INF Starting localClient service connection=mempool impl=localClient module=abci-client
12:55PM INF Starting localClient service connection=consensus impl=localClient module=abci-client
12:55PM INF Starting EventBus service impl=EventBus module=events
12:55PM INF Starting PubSub service impl=PubSub module=pubsub
12:55PM INF Starting IndexerService service impl=IndexerService module=txindex
12:55PM INF ABCI Handshake App Info hash="⽻F�����3�+�\x054��6v�#��(���m\x00q\bj" height=193 module=consensus protocol-version=0 software-version=0.3.2
12:55PM INF ABCI Replay Blocks appHeight=193 module=consensus stateHeight=193 storeHeight=193
```

Now, Cosmovisor is up and running.

### Upgrade version binary download & setup

The user must already be running FirmaChain v0.3.1 and this process illustrates the necessary preparations to upgrade the chain to v0.3.2 using Software Proposal.

#### 1. Folder Structure Setting

```
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/v0.3.2/bin
```

#### 2. Binary File Download

```
curl https://build.firmachain.dev/@v0.3.2 | bash
```

#### 3. Copy and Paste the Binary File to the Cosmovisor Upgrade Folder Directory

```
cp firmachaind $DAEMON_HOME/cosmovisor/upgrades/v0.3.2/bin
```

## Post-Upgrade

**`cosmovisor`** should be up and running using the upgraded file in v0.3.2/bin folder. Please use the command provided below to check the upgraded binary version.

```
$ cosmovisor version
Cosmovisor Version:
12:52PM INF Configuration is valid:
Configurable Values:
  DAEMON_HOME: /home/firmauser/.firmachain
  DAEMON_NAME: firmachaind
  DAEMON_ALLOW_DOWNLOAD_BINARIES: false
  DAEMON_RESTART_AFTER_UPGRADE: true
  DAEMON_POLL_INTERVAL: 300ms
  UNSAFE_SKIP_BACKUP: false
  DAEMON_PREUPGRADE_MAX_RETRIES: 0
Derived Values:
        Root Dir: /home/firmauser/.firmachain/cosmovisor
     Upgrade Dir: /home/firmauser/.firmachain/cosmovisor/upgrades
     Genesis Bin: /home/firmauser/.firmachain/cosmovisor/genesis/bin/firmachaind
  Monitored File: /home/firmauser/.firmachain/data/upgrade-info.json
 module=cosmovisor
12:52PM INF running app args=["version"] module=cosmovisor path=/home/firmauser/.firmachain/cosmovisor/upgrades/v0.3.2/bin/firmachaind
0.3.2
```
