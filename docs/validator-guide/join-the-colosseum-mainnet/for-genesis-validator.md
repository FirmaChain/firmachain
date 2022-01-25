# For Genesis Validator

## Instruction

This document is written for FirmaChain’s Genesis Validators, and if you are indeed a Genesis Validator participant, please be sure to double check your gentx submission deadline and the starting time of the **Colosseum** mainnet. In order to submit your gentx file to the FirmaChain team, please do so after updating your binary file to the most recent version. ‘Deadline’ and the ‘Mainnet Launch’ schedule is as follows.

* _**Mainnet will launch on 01/26/2022 14:00 UTC.**_
* _**The deadline to submit your gentx is 01/19/2022 14:00 UTC.**_

_\* To submit your gentx file to the FirmaChain team, please follow the steps listed below._

## Setup FirmaChain

If you have a previous record of operating a node and hence have a .firmachain folder, we recommend you delete the existing folder and start from scratch.

```
rm -rf ~/.firmachain
```

### Download FirmaChain’s most recent binary file.

```
cd ~
curl https://build.firmachain.org | bash
```

### Check the binary build version and the integrity before setting global command.

```
./firmachaind version
sha1sum ./firmachaind

sudo mv ~/firmachaind /usr/local/bin/firmachaind
```

### Initialize(Reset) your firmachain node folder using the command provided below.

```
firmachaind init <your_moniker_name> --chain-id colosseum-1
```

### Retrieve your wallet using your mnemonic.

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

### Register Genesis account

```
firmachaind add-genesis-account <your_wallet_address> 10000000ufct
```

Please check whether the \<your\_wallet\_address> value is identical to the address you submitted during the KYC process.\
_**\* When registering a genesis account, for the amount field, you MUST enter 10000000ufct.**_\
&#x20;  _**If a different amount was put in, please start again by resetting your folder.**_

### Create gentx file (★)

```
firmachaind gentx <key_name> 10000000ufct --chain-id colosseum-1 \
--commission-rate 0.1 \
--commission-max-rate 0.2 \
--commission-max-change-rate 0.01 \
--moniker <moniker_name> \
--website <website_link> \
--details <description> \
--security-contact <email> \
--identity <key base 64bit code>
```

"Keybase 64bit code" can be found through [this link](edit-validator-description.md#how-to-get-identity-64bit-code).

### Check gentx file

If all of the above steps were completed without any error, you will be able to find a json file starting with ‘gentx-’ in the \~/.firmachain/config/gentx/ path.

```
{
  "body": {
    "messages": [
      {
        "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
        "description": {
          "moniker": "test-node",
          "identity": "",
          "website": "firmachain.org",
          "security_contact": "contact@firmachain.org",
          "details": "hello world\\n"
        },
        "commission": {
          "rate": "0.100000000000000000",
          "max_rate": "0.200000000000000000",
          "max_change_rate": "0.010000000000000000"
        },
        "min_self_delegation": "1",
        "delegator_address": "firma1qvrmp6hyyveds0vgenc5hjy0qtrc76hy33fadr",
        "validator_address": "firmavaloper1qvrmp6hyyveds0vgenc5hjy0qtrc76hy0zzxdd",
        "pubkey": {
          "@type": "/cosmos.crypto.ed25519.PubKey",
          "key": "Vlow94hC+SC2sOb0mHuoxGXG/nfzboWK5HfjFSBJE88="
        },
        "value": {
          "denom": "ufct",
          "amount": "10000000"
        }
      }
    ],
    "memo": "1157a6860320735421d7ef49a12b86e27727827d@192.168.20.111:26656",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [
      {
        "public_key": {
          "@type": "/cosmos.crypto.secp256k1.PubKey",
          "key": "AlnRwPXCoLNNTrFDZHmmosFwxsSB6ovibUF16zvAE7S+"
        },
        "mode_info": {
          "single": {
            "mode": "SIGN_MODE_DIRECT"
          }
        },
        "sequence": "0"
      }
    ],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": [
    "cUTd4hFxYfayvvnHKF+cSUlJR+6DpC6GEb0RHsoBNmhI8WlnGa8KcRic/vMu0Zd/CGxP8B7yYKu3a+dqOpGitg=="
  ]
}
```

### Create pull request

You must submit the gentx file to the FirmaChain team as a pull request to the https://github.com/FirmaChain/mainnet/gentxs/ directory.\
_**\* Please change the file name to "vaildator\_name.json" and submit it.**_

### Download genesis.json (★)

After collecting the gentx from our genesis validators, we will share a consolidated genesis.json file via mainnet git. Genesis validators should download the consolidated file.

```
wget https://github.com/FirmaChain/mainnet/raw/main/genesis.json
```

### Replace genesis.json

```
mv ~/genesis.json ~/.firmachain/config/genesis.json
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
Details are available on [github](https://github.com/FirmaChain/mainnet).

```
#######################################################
###           P2P Configuration Options             ###
#######################################################
...

# example: 159.89.10.97:26656
external_address = ""

# Comma separated list of seed nodes to connect to
seeds = "seed list"
```

#### Input your 'external address'

```
external_address = "your_external_address:26656"
```

#### Seed list (copy seeds)

```
f89dcc15241e30323ae6f491011779d53f9a5487@mainnet-seed1.firmachain.dev:26656,04cce0da4cf5ceb5ffc04d158faddfc5dc419154@mainnet-seed2.firmachain.dev:26656,940977bdc070422b3a62e4985f2fe79b7ee737f7@mainnet-seed3.firmachain.dev:26656
```

### Start FirmaChain

```
firmachaind start
```

### Register as daemon (Optional)

It is absolutely crucial that the FirmaChain nodes remain active at all times. The simplest solution would be to register this as a system. After a reboot or any other type of event, the service registered on the system will be activated and hence, FirmaChain will be able to start the operation of the nodes.

```
sudo tee /etc/systemd/system/firmachaind.service > /dev/null <<EOF  
[Unit]
Description=Firmachain Node
After=network-online.target

[Service]
User=$USER
ExecStart=$(which firmachaind) start
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
