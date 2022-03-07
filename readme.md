# FirmaChain

[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/firmachain/firmachain)](https://github.com/firmachain/firmachain/releases)
[![license](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/firmachain/firmachain/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/firmachain/firmachain/.svg)](https://pkg.go.dev/github.com/firmachain/firmachain/)
[![report](https://goreportcard.com/badge/github.com/firmachain/firmachain)](https://goreportcard.com/report/github.com/firmachain/firmachain)
![Lines of code](https://img.shields.io/tokei/lines/github/firmachain/firmachain)

![ci](https://github.com/firmachain/firmachain/actions/workflows/test.yml/badge.svg)
<a href="https://codecov.io/gh/firmachain/firmachain">
    <img src="https://codecov.io/gh/firmachain/firmachain/branch/master/graph/badge.svg">
</a>


### A BLOCKCHAIN BASED E-CONTRACT PLATFORM

FirmaChain adds a signing and verifying e-contract function to the blockchain. 

Unlocking new limits of electronic contracts with blockchain technology, FirmaChain seeks to resolve all the social and legal issues (contracts, notarial, etc.) with written contracts through the use of electronic contracts based on FirmaChain’s data blockchain.

FirmaChain now builds on [Tendermint](https://github.com/tendermint/tendermint) consensus and the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) toolkits and [Starport](https://github.com/tendermint/starport)

</br>

---

## Installation

### 1. From curl command from https://build.firmachain.dev

You can download a pre-built binary for your operating system easily below command. 
```bash
curl https://build.firmachain.dev | bash
```

or you can find the latest binaries on the [releases](https://github.com/firmachain/firmachain/releases) page.

### 2. From Source

**Step 1. Install Golang**

Go v1.16+ or higher is required for FirmaChain and Starport.


**Step 2: Install Starport (customized version)**
```
git clone https://github.com/firmachain/starport
cd starport
make install
sudo mv ~/go/bin/starport /usr/local/bin
```

**Step 3: Install FirmaChain**


```bash
git clone https://github.com/firmachain/firmachain
cd firmachain
starport chain build
```

After build, you can find the firmachaind file on "~/go/bin/firmachaind".

If you want to run FirmaChain on dev env, just run this command.

```bash
starport chain serve
```

</br>

---
</br>

## Validator & Node
We make governance with chosen validators.
We are currently looking for validators to contribute to our network. 

If you are interested, please contact us through contact@firmachain.org


## Network 

- Maintnet **Colosseum-1** launced at 2022.01.26
    - FirmaChain is currently operating on Ethereum as an ERC-20 Token FCT. 
    - All contracts and tokens will be migrated to our main network once it is complete network stabilization.
- Testnet **Imperium-3** is ready for online. (private testnet for only chain developer)

## EcoSystem
- [FirmaChain-BlockExplorer](https://github.com/FirmaChain/firmachain-explorer)
- [FirmaChain-Station (on development)](https://github.com/FirmaChain/firma-station) 
- [FirmaChain-Faucet (devnet, testnet only)](https://github.com/FirmaChain/firmachain-faucet) 

## Community
- [Website](https://firmachain.org/#/)
- [Medium](https://medium.com/firmachain)
- [Telegram](https://t.me/firmachain_announcement)
- [Twitter](https://twitter.com/firmachain)
