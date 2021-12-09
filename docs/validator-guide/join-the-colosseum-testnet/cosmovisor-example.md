# Cosmovisor Example

The purpose of this document is to inform the users on how to upgrade the chain to v0.3.2 using Cosmovisor. All processes mentioned in this document must be conducted on the binary v0.3.1 preconfigured chain server.

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

In order to launch Cosmovisor, please enter the command provided below.

```
cosmovisor start
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
curl https://build.firmachain.org/@v0.3.2 | bash
```

#### 3. Copy and Paste the Binary File to the Cosmovisor Upgrade Folder Directory

```
cp .firmachaind $DAEMON_HOME/cosmovisor/upgrades/v0.3.2/bin
```

## Post-Upgrade

Cosmovisor should be up and running using the upgraded file in v0.3.2/bin folder. Please use the command provided below to check the upgraded binary version.

```
cosmovisor version
# 0.3.2
```
