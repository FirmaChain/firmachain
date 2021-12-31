# Join the FirmaChain Node

### 1. Download genesis.json (★)

In order to participate in the mainnet you will need a genesis.json file. Genesis.json file can be found in the FirmaChain github repository and can be downloaded from server local using the following command

```
cd ~
wget https://github.com/FirmaChain/firmachain-mainnet-colosseum/raw/master/genesis.json
```

_\* Once we receive all the gentx files from our Genesis Validators and compile them into a genesis, we will provide the genesis.json file._

### 2. Replace genesis.json

```
mv ~/genesis.json ~/.firmachain/config/genesis.json
```

### 3. Increase maximum open files

In general, most operating systems limit the maximum number of files that can be opened.\
(The document was written based on `Ubuntu 20.04 LTS`, and the default value of this OS is 1024.)\
However, if you open `/etc/security/limits.conf` and modify the value as can be seen in the code below, you can raise the limit of the nofile feature

```
*                soft    nofile          65535
*                hard    nofile          65535
```

### 4. Start FirmaChain

```
firmachaind start
```

### 5. Register as daemon (Optional)

It is absolutely crucial that the FirmaChain nodes remain active at all times. The simplest solution would be to register this as a system. After a reboot or any other type of event, the service registered on the system will be activated and hence, FirmaChain will be able to start the operation of the nodes.

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
LimitNOFILE=65535

Environment="DAEMON_HOME=$HOME/.firmachain"
Environment="DAEMON_NAME=firmachaind"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="UNSAFE_SKIP_BACKUP=false"

[Install]
WantedBy=multi-user.target
EOF
```

Now you are all set to operate FirmaChain using daemon. Please join our network using the command provided below.

```
sudo systemctl daemon-reload
sudo systemctl restart firmachaind
```
