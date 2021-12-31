# For Genesis Validator

## Instruction

This document is written for FirmaChain’s Genesis Validators, and if you are indeed a Genesis Validator participant, please be sure to double check your gentx submission deadline and the starting time of the **Colosseum** mainnet. In order to submit your gentx file to the FirmaChain team, please do so after updating your binary file to the most recent version. ‘Deadline’ and the ‘Mainnet Launch’ schedule is as follows.

* _**Mainnet will launch on 01/11/2022 13:00 UTC.**_
* _**The deadline to submit your gentx is 01/11/2022 13:00 UTC.**_

_\* To submit your gentx file to the FirmaChain team, please follow the steps listed below._

## Setup FirmaChain

If you have a previous record of operating a node and hence have a .firmachain folder, we recommend you delete the existing folder and start from scratch.

```
rm -rf ~/.firmachain
```

### 1. Download FirmaChain’s most recent binary file.

```
cd ~
curl htps://build.firmachain.org | bash
```

### 2. Check the binary build version and the integrity before setting global command.

```
./firmachaind version
sha1sum ./firmachaind

sudo mv ~/firmachaind /usr/local/bin/firmachaind
```

### 3. Initialize(Reset) your firmachain node folder using the command provided below.

```
firmachaind init <your_moniker_name> —chain-id colosseum-1
```

### 4. Retrieve your wallet using your mnemonic.

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

### 5. Register Genesis account

```
firmachaind add-genesis-account <your_wallet_address> 10000000ufct
```

Please check whether the \<your\_wallet\_address> value is identical to the address you submitted during the KYC process.\
_**\* When registering a genesis account, for the amount field, you MUST enter 10000000ufct.**_\
&#x20;  _**If a different amount was put in, please start again by resetting your folder.**_

### 6. Create gentx file (★)

```
firmachaind gentx <key_name> --amount 10000000ufct --chain-id colosseum-1
--commission-rate 0.1
--commission-max-rate 0.2
--commission-max-change-rate 0.01
--moniker <moniker_name>
--website <website_link>
--details <description>
--security-contact <email>
```

### 7. Check gentx file

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

### 8. Create pull request

You must submit the gentx file to the FirmaChain team as a pull request to the https://github.com/firmachain/colosseum-1/gentx/ directory.\
_**\* Please change the file name to "vaildator\_name.json" and submit it.**_

_****_

Please go to [Join the FirmaChain Node](broken-reference)  document once all of the above steps are complete.
