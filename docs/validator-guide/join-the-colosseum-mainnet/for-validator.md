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
sed -i 's/minimum-gas-prices = "0stake"/minimum-gas-prices = "1ufct"/g' ~/.firmachain/config/app.toml
```

#### P2P options

FirmaChain discloses information on seed nodes for the purpose of P2P connection.\
The list of seed addresses can be found in **this link**.

```
#######################################################
###           P2P Configuration Options             ###
#######################################################
[p2p]
...
# Comma separated list of seed nodes to connect to
seeds = ""
```

example)

```
seeds = "id0000000000000000@13.51.211.18:26656,id0000000000000001@38.209.37.78:26656"
```

Please go to [Join the FirmaChain Node](broken-reference) document once all of the above steps are complete.
