# Recover the chain node

## Instruction

FirmaChain provides a compressed data file that contains the block information up to a certain height. By downloading the compressed data file, you can reduce the time consumed to sync the network.

### Download 'data.tar'

```
cd ~
wget https://firmachain.org/mainnet/build/firmachain-augustus-20211230-v0.3.2.tar
```

### Verify the Integrity of the Compressed File

In order for you to check whether you have successfully downloaded the compressed file, you must verify the integrity of the file.

```
sha1sum data.tar
f10d9am29a0m2n18and0a82n10ausn2191ma9891fj83n6
```

### Decompress

You must decompress the verified data.tar file and must apply the file to the chain.

```
tar -xvf data.tar
```

_\* If the error message “tar: Unexpected EOF in archive” occurs during decompression, please download the compressed data file again._

### Moving the Folder

Please delete the existing data directory before executing this step.

```
rm -rf .firmachain/data
mv data ~/.firmachain/
```

### Start Chain

Once you start the chain, the blocks will accumulate from the backup data and you will be able to expedite your synchronization process.

```
firmachaind start
```
