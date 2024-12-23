# Usage on FirmaChain-CLI

## Install FirmaChain CLI

The following is an explanation on how to issue a transaction on your Ledger Wallet.

#### 1. Install Starport

```
git clone https://github.com/firmachain/starport

cd starport

make install

# use to global
sudo mv ~/go/bin/starport /usr/local/bin/
```

_\* If you have not installed go-lang, please install go-lang by running the command provided below._

```
sudo snap install go --classic
```

#### 2. Install FirmaChain

```
git clone https://github.com/firmachain/firmachain/v05.git

cd firmachain/

starport chain build

# use to global
sudo mv ~/go/bin/firmachaind /usr/local/bin/
```

#### 3. Start FirmaChain (customized version)

```
firmachaind start
```

## Create an address with the Ledger

Please unlock your Ledger Nano S/X/S Plus by entering the pin code. Additionally, if either your Ledger Live application is running or if you have a running application that connects your Ledger Nano S/X/S Plus (e.g. Ledger Live), please close all such applications.&#x20;

_\* Ledger Nano S/X/S Plus only supports 1:1 connection._

### Preparations

This stage outlines the necessary steps you must take to create a Ledger Nano S/X/S Plus wallet.

**initialize chain**

```
firmachaind init test-node-001 --chain-id roma-1
```

**change gas prices**

```
vim ~/.firmachain/config/app.toml
# minimum-gas-prices = "0.01ufct"
```

### Open the FirmaChain application.

You can navigate through the apps by using the two buttons located on the top of the ledger. Once you choose the FirmaChain app, you can open the application by clicking the two buttons simultaneously.

_\* Currently, the "Firma Chain app" is distributed in developer mode, and the "Pending Ledger review" message is visible in developer mode. Press the top two buttons at the same time to move on._

<figure><img src="../../.gitbook/assets/image (10).png" alt=""><figcaption></figcaption></figure>

### Create FirmaChain Wallet using the Ledger Tag

By running the command provided below, you can create the key information connected with your Ledger Nano S/X/S Plus on local.

```
firmachaind keys add <key_name> --coin-type 7777777 --ledger
```

If you are using mnemonic, please run the command provided below.

```
firmachaind keys add <key_name> --recover --coin-type 7777777 --ledger
```

Once you complete the steps above, an address will appear on the screen of your connected Ledger Nano S/X/S Plus. By clicking on the button on the right, you can check your wallet information. Finally, approve by simultaneously clicking on the two buttons located on the top of your Ledger.

![](<../../.gitbook/assets/image (8) (1).png>)

![](<../../.gitbook/assets/image (33).png>)

Now, the address created with your \<key\_name> will include the information of your Ledger Nano S/X/S Plus.

## Receive FirmaChain Tokens <a href="#b6bc" id="b6bc"></a>

In order to receive tokens, you must be in possession of your wallet address information.\
Using the following command, you can check your wallet information.

```
firmachaind keys list
```

Following is the information of your Ledger Nano S/X/S Plus you can obtain by running the above command.

```
- name: <key_name>
  type: ledger
  address: <your_nano_key_address>
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Biqdmqjkl12na092mald9andiqaxz"}'
  mnemonic: ""
```

_\* All keys saved on local will be displayed. The type of your Ledger Nano S/X/S Plus wallet information will be shown as ‘ledger.’_

## Send Tokens

In order to send tokens, your node must be participating as a full node. \
If you’ve followed all the necessary steps outlined in this document, you will be able to continue the remaining process on local.

### Execute command

Do you wish to send FirmaChain tokens? If so, please run the command provided below.

```
firmachaind tx bank send <ledger_nano_key_name> \
<send_to_address> 10000ufct \
--from <legder_nano_key_name> \
--fees 2000ufct \
--chain-id imperium-2 \
--note hello_firmachain
```

_\* --note_ is optional. _Only the flag has been changed to --note from the original memo._\
&#x20; _The item you must pay attention to at this stage is ._ \
&#x20; _The  must show the name of your Ledger Nano S/X/S Plus._

If you’ve successfully sent the above command (transaction) from your OS, please check your Ledger Nano S/X/S Plus. You approve the transaction request from your Ledger Nano S/X/S Plus. Using the top right button of your Ledger Nano S/X/S Plus, please check whether all transaction information is correct before proceeding to the ‘Approve or Reject’ stage.

_\* You can proceed by simultaneously clicking on the two buttons located on the top of your Ledger._

![](<../../.gitbook/assets/image (7).png>)

From the FirmaChain full node, you can check the change of your wallet balance.\
Please run the command below to check your balance.

```
firmachaind q bank balances <your_nano_key_address>
```

If you run the command provided above, you will be able to see the following results.

```
balances:
- amount: "399999936000"
  denom: ufct
pagination:
  next_key: null
  total: "0"
```

Ledger Nano S/X/S Plus can be used with FirmaChain CLI.

### Support

If you have any issue with your Ledger installation and use station, please join to our community channel and email.

Telegram : [https://t.me/firmachain\_global](https://t.me/firmachain\_global)\
Mailto : [contact@firmachain.org](mailto:contact@firmachain.org?bcc=contact@firmachain.org)
