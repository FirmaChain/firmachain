# Cosmovisor Guide

**`cosmovisor`** is a small process manager for Cosmos SDK application binaries that monitors the governance module for incoming chain upgrade proposals. If it sees a proposal that gets approved, **`cosmovisor`** can automatically download the new binary, stop the current binary, switch from the old binary to the new one, and finally restart the node with the new binary.

## Setup

### **1.** Download Cosmovisor

Using the command below, users can download the latest **`cosmovisor`** binary file. We recommend you to use the most recent version of **`cosmovisor`**.

```
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest
```

You can also enter the Tag, as shown in the code below to download the original version.

```
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v0.1.0
```

Please copy the downloaded **`cosmovisor`** binary using the command below.

```
sudo cp cosmovisor /usr/local/bin/cosmovisor
```

_\* If you are using go v1.15 or earlier, you will need to use go get, and you may want to run the command outside a project directory._

### **2. Setting up environmental variables**

**`cosmovisor`** reads its configuration from environment variables:

* DAEMON\_HOME
  * This is the folder directory of the running chain.
  * `export DAEMON_HOME=$HOME/.firmachain`
* DAEMON\_NAME
  * This is the directory of the binary file once the chain starts running.
  * `export DAEMON_NAME=firmachaind`
* DAEMON\_ALLOW\_DOWNLOAD\_BINARIES (optional, default = false)
  * This setting asks whether you would automatically download the binary file.
  * Due to update issues, Cosmos does not recommend you true the AutoDownload. [(Link)](https://docs.cosmos.network/master/run-node/cosmovisor.html#auto-download)
* DAEMON\_RESTART\_AFTER\_UPGRADE (optional, default = true)
  * If you set the variable as true, the chain restarts using the new binary, post-upgrade.
  *   If you set the variable as false, the chain must be restarted manually by the manager post-upgrade.

      _\* IMPORTANT : Restart will only happen after the upgrade and if any error occurs during the upgrade, the chain will not restart._
* DAEMON\_POLL\_INTERVAL (optional, default = 300ms)
  * This is the length of the space to poll the upgrade plan file. (e.g. 1s).
  * The value can either be a number (in milliseconds) or a duration (e.g. 1s).
* UNSAFE\_SKIP\_BACKUP (defaults to false)
  * If you set the variable as true, the upgrade will start without any data backup.\
    If you wish to backup your data before upgrade, please set this variable as false or delete this settings category.
  * If you set the variable as false, it becomes easier to roll back and therefore, we recommend you use this feature.
* DAEMON\_PREUPGRADE\_MAX\_RETRIES (defaults to 0)
  * This option sets the number of pre upgrade call attempts.
  * If the pre upgrade call attempt reaches the set limit due to consistent failure, **`cosmovisor`** will fail to upgrade.

Following is a sample setting of the above variables.\
Options can be modified in ways that suits your purpose after reading the description of each variable mentioned above.\
DAEMON\_HOME and DAEMON\_NAME must be used as is. (Do not modify!)

```
export DAEMON_HOME=$HOME/.firmachain
export DAEMON_NAME=firmachaind
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true
export DAEMON_POLL_INTERVAL=300
export UNSAFE_SKIP_BACKUP=false
export DAEMON_PREUPGRADE_MAX_RETRIES=0
```

The explanation above describes the option variables to run **`cosmovisor`** and in order to safely manage and use **`cosmovisor`** we recommend you add to the \~/.profile file.

```
sudo nano ~/.profile
```

### **3. Copying 'firmachaind' files in the proper folders**

**`cosmovisor`** must be able to read and run the firmachain binary. Please execute the command provided below.

```
mkdir -p ~/.firmachain/cosmovisor/genesis/bin/
cp $(which firmachaind) ~/.firmachain/cosmovisor/genesis/bin/
```



:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

_<mark style="color:red;">**※ In case of 1.0.0 version.**</mark>_

* Please execute the command below.\
  If you don't execute this command, you can't run **`cosmovisor`**.

```
echo "{}" >> /home/firma/.firmachain/cosmovisor/current/upgrade-info.json
```

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::



Whether you’ve properly configured the settings mentioned above can be determined using the command provided below.

```
cosmovisor version
```

In order to run with **`cosmovisor`**, you must quite the original firmachaind process.

```
ps -ef | grep firmachaind
pkill <process id>
```

Finally, start **`cosmovisor`**

```
cosmovisor start
```

### **4.** Registering Cosmovisor to the system (Optional)

If your server went down and you have to restart, you can let it restart automatically by registering the service. Please follow the example provided below.

```
sudo tee /etc/systemd/system/firmachaind.service > /dev/null <<EOF  
[Unit]
Description=Firmachain Node
After=network-online.target

[Service]
User=$USER
ExecStart=$(which cosmovisor) start
Restart=always
RestartSec=3
LimitNOFILE=4096

Environment="DAEMON_HOME=$HOME/.firmachain"
Environment="DAEMON_NAME=firmachaind"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="UNSAFE_SKIP_BACKUP=false"

[Install]
WantedBy=multi-user.target
EOF
```

Once you’ve created the system file, please register the file to the system by using the command provided below.

```
sudo systemctl daemon-reload
```

Finally, restart.

```
sudo systemctl restart firmachaind
```
