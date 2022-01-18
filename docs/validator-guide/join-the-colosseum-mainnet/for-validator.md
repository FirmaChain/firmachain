# For Validator

## Instruction

This document illustrates the necessary steps a participant must take should they wish to participate in the node operation as a validator after the launch of the mainnet. If you wish to participate in the FirmaChain mainnet network, please refer to this document.

## Start Setup

### Initialize chain node

```
firmachaind init <your_node_name> --chain-id colosseum-1
```

### Restore Wallet

#### Restore Key (Validators must have mnemonic)

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

If you do not have a mnemonic or if you are creating a new wallet, please use the following command.

```
firmachaind keys add <key_name> --coin-type 7777777
```

### Node configuration file

You must modify your configuration file in order for you to join the FirmaChain network.

#### Change Minimum gas prices

Firstly, look at \~/.firmachain/config/app.toml file.\
You can reject any incoming transaction that is lower than the minimum gas price.

```
sed -i 's/minimum-gas-prices = "0stake"/minimum-gas-prices = "0.1ufct"/g' ~/.firmachain/config/app.toml
```

#### P2P options

FirmaChain discloses information on seed nodes for the purpose of P2P connection.\
The list of seed addresses can be found in [**this link**](https://github.com/FirmaChain/mainnet).

```
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
fffa9c85e3182342e4db7fc8027332c43a0cfa15@mainnet-seed1.firmachain.dev:26656,3ca16236b26a83ab8ab5de583c20a79b9843c274@mainnet-seed2.firmachain.dev:26656,8335d246b6703d112ae0726bfc2b6e3a5b0010c2@mainnet-seed3.firmachain.dev:26656
```

### Download genesis.json (★)

In order to participate in the mainnet you will need a genesis.json file. Genesis.json file can be found in the FirmaChain github repository and can be downloaded from server local using the following command

```
wget https://github.com/FirmaChain/mainnet/raw/master/genesis.json
```

### Replace genesis.json

```
mv ~/genesis.json ~/.firmachain/config/genesis.json
```

### Start FirmaChain

```
firmachaind start
```

### Register as Validator

```
firmachaind tx staking create-validator \
--pubkey $(firmachaind tendermint show-validator) \
--moniker <Your moniker name> \
--chain-id imperium-3 \
--commission-rate 0.10 \
--commission-max-rate 0.20 \
--commission-max-change-rate 0.01 \
--min-self-delegation 1 \
--identity <key base 64bit code> \
--amount <staking amount>ufct \
--fees 20000ufct \
--from <key_name>
```

"Keybase 64bit code" can be found through [this link](edit-validator-description.md#how-to-get-identity-64bit-code).

### Register as daemon (Optional)

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
