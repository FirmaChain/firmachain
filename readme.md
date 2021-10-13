# FirmaChain

### A BLOCKCHAIN BASED E-CONTRACT PLATFORM

FirmaChain adds a signing and verifying e-contract function to the blockchain. 

Unlocking new limits of electronic contracts with blockchain technology, FirmaChain seeks to resolve all the social and legal issues (contracts, notarial, etc.) with written contracts through the use of electronic contracts based on FirmaChain’s data blockchain.

FirmaChain now builds on [Tendermint](https://github.com/tendermint/tendermint) consensus and the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) toolkits and [Starport](https://github.com/tendermint/starport)

</br>

---

## Installation

### 1. From curl command from https://build.firmachain.org

You can download a pre-built binary for your operating system easily below command. 
```bash
curl https://build.firmachain.org | bash
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

- Testnet **Imperium-2** is ready for online. (private testnet for only chain developer)
- Testnet **Colosseum-1** is ready for online. (public testnet for validator)
- Mainnet **Augustus-1** launched at 29.09.2020
	- FirmaChain is currently operating on Ethereum as an ERC-20 Token **FCT**. All contracts and tokens will be migrated to our main network once it is complete network stabilization.

## EcoSystem
- [FirmaChain-BlockExplorer](https://github.com/FirmaChain/firmachain-explorer)
- [FirmaChain-Station (on development)](https://github.com/FirmaChain/firma-station) 
- [FirmaChain-Faucet (devnet, testnet only)](https://github.com/FirmaChain/firmachain-faucet) 

## Community
- [Website](https://firmachain.org/#/)
- [Medium](https://medium.com/firmachain)
- [Telegram](https://t.me/firmachain_announcement)
- [Twitter](https://twitter.com/firmachain)
