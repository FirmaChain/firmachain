# FirmaChain

[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/firmachain/firmachain)](https://github.com/firmachain/firmachain/v05/releases)
[![license](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/firmachain/firmachain/v05/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/firmachain/firmachain/v05/.svg)](https://pkg.go.dev/github.com/firmachain/firmachain/v05/)
[![report](https://goreportcard.com/badge/github.com/firmachain/firmachain/v05)](https://goreportcard.com/report/github.com/firmachain/firmachain/v05)
![Lines of code](https://img.shields.io/tokei/lines/github/firmachain/firmachain)

![ci](https://github.com/firmachain/firmachain/v05/actions/workflows/test.yml/badge.svg)
<a href="https://codecov.io/gh/firmachain/firmachain">
    <img src="https://codecov.io/gh/firmachain/firmachain/branch/master/graph/badge.svg">
</a>


## A BLOCKCHAIN BASED E-CONTRACT PLATFORM

FirmaChain adds a signing and verifying e-contract function to the blockchain.

Unlocking new limits of electronic contracts with blockchain technology, FirmaChain seeks to resolve all the social and legal issues (contracts, notarial, etc.) with written contracts through the use of electronic contracts based on FirmaChain’s data blockchain.

FirmaChain now builds on [CometBFT](https://github.com/cometbft/cometbft) consensus and the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) toolkits.

## Installation

### Install Go ###

**Go v1.21 or higher** is required for building and running FirmaChain.

**NOTE** : we updated Go requirements from `v1.18` to `v1.21` in `Firmachain v0.4.0`

</br>

### Official Build  ###

We are not providing official prebuilt binaries from `v0.3.5`. Please check the guide to build your own binary.

https://docs.firmachain.org/master/node-and-validators-guide/run-a-full-node/installation-firmachaind

Currently available tags:
| Version | Tag
| - | - |
| v0.4.0 | `v0.4.0` |
| v0.3.5 | `v0.3.5-patch` |
| v0.3.3 (Deprecated) | `v0.3.3-patch` |

</br>

### Build Guide : using make ###

```bash
git clone https://github.com/firmachain/firmachain/v05
cd firmachain
git checkout (desired tags)

make install
...
```

After the successful build, you will find `firmachaind` binary on the `~/go/bin` build path.

</br>

## Validator & Node
We make governance with chosen validators.
We are currently looking for validators to contribute to our network. 

If you are interested, please contact us through contact@firmachain.org

</br>

## Network 

- Firmachain Mainnet **Colosseum-1** launched on 2022.01.26
    - FirmaChain is currently operating on Ethereum as an ERC-20 Token FCT. 
    - All contracts and tokens will be migrated to our main network once it is complete network stabilization.
- Firmachain Pubilc Testnet **Imperium-4** is also available.

</br>

## Ecosystem
- [Firmachain Block Explorer](https://github.com/firmachain/firmachain/v05-explorer)
- [Firma Station](https://github.com/FirmaChain/firma-station) 
- [Firmachain Faucet (Testnet, Imperium-4)](https://github.com/firmachain/firmachain/v05-faucet) 

</br>

## Community
- [Website](https://firmachain.org/#/)
- [Medium](https://medium.com/firmachain)
- [Telegram](https://t.me/firmachain_announcement)
- [Twitter](https://twitter.com/firmachain)
