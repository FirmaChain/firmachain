# Install

### Ways 1. Install FirmaChain from pre-built binary by curl

You can easily download a pre-built binary for your operating system using the command below.

#### download binary

```shell
# On Bash
curl https://build.firmachain.dev | bash

sudo mv ./firmachaind /usr/local/bin/firmachaind
```

#### check version & sha1sum

```
firmachaind version
0.3.3-6bcd9e15
```

```
firmachaind version --long
name: Firmachain
server_name: firmachaind
version: 0.3.3-6bcd9e15
commit: 6bcd9e15407f5207e8a2d03e6488301aecb9f98b
...
cosmos_sdk_version: v0.44.5-patch
```

```
sha1sum /usr/local/bin/firmachaind
b59534ef087becc442f04a6cd42e4fc153c18bb7 firmachaind
```

### Ways 2. Download from Github Release page

You can download a prebuilt binary from the link below. [https://github.com/firmachain/firmachain/v05/releases](https://github.com/firmachain/firmachain/v05/releases)​

<figure><img src="../../.gitbook/assets/image (8).png" alt=""><figcaption><p><a href="https://github.com/firmachain/firmachain/v05/releases">https://github.com/firmachain/firmachain/v05/releases</a></p></figcaption></figure>

Select, download and unzip a binary that suits your OS.\
Please confirm whether you’ve downloaded the correct version, using the method provided in Ways 1.
