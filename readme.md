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

FirmaChain now builds on [Tendermint](https://github.com/tendermint/tendermint) consensus and the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) toolkits and [Ignite](https://github.com/ignite/cli/)

---

## Installation

## - Install Go ##

Go v1.18+ or higher is required for FirmaChain.

</br>

## - Official Build  ##


https://docs.firmachain.org/master/getting-started/install-firmachain/install
https://docs.firmachain.org/master/validator-guide/upgrade/v0.3.5

</br>


## - Development Build : using make ##

```bash
git clone https://github.com/firmachain/firmachain
cd firmachain
make install
```

</br>

## - Development Build : using ignite ##

**Step 1: Install ignite v0.23.0**
```
// 1. download and unarchive the file below. 
https://github.com/ignite/cli/releases/tag/v0.23.0

// Then run this command to move the ignite executable to /usr/local/bin/:
sudo mv ignite /usr/local/bin/
```

**Step 2: Install FirmaChain**

```bash
git clone https://github.com/firmachain/firmachain
cd firmachain
ignite chain build
```

After build, you can find the firmachaind file on "~/go/bin/firmachaind".

If you want to run FirmaChain on dev env, just run this command.

```bash
ignite chain serve
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
- Testnet **Imperium-4** is ready for online. (private testnet for only chain developer)

## EcoSystem
- [FirmaChain-BlockExplorer](https://github.com/FirmaChain/firmachain-explorer)
- [FirmaChain-Station](https://github.com/FirmaChain/firma-station) 
- [FirmaChain-Faucet (testnet only)](https://github.com/FirmaChain/firmachain-faucet) 

## Community
- [Website](https://firmachain.org/#/)
- [Medium](https://medium.com/firmachain)
- [Telegram](https://t.me/firmachain_announcement)
- [Twitter](https://twitter.com/firmachain)
