# Deploy own network

#### Have you completed the previous step?

* If you are unsure about the minimum specifications or the ‘firmachaind’ binary files, please visit the following link. [Install FirmaChain](broken-reference)

## How to set-up

In order to become a Validator, you must follow the direction below.

### Initialize chain node

```
firmachaind init <node_name> --chain-id <Your chain id>
```

### Change gas prices

```
vi ~/.firmachain/config/app.toml

###############################################################################
###                           Base Configuration                            ###
###############################################################################

minimum-gas-prices = "1ufct"
```

### Create Wallet

#### Create your key (If you don’t have a mnemonic)

```
firmachaind keys add <key_name> --coin-type 7777777
Enter keyring passphrase: XXXXXXXX
Re-enter keyring passphrase: XXXXXXXX
# must be at least 8 characters

- name: firmachain
  type: local
  address:
  pubkey: ''
  mnemonic: ""
  
**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

glad music grace lawn squeeze book very text wire okay ozone morning permit tumble guard hurry various gallery kitten surprise brain piano level picture
```

* Please save your mnemonic phase in a safe location once you’ve created your key. If you lose your mnemonic, you won’t be able to retrieve it again.

#### Restore key (If you have a mnemonic)

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

### genesis.json setting

#### change denom

```
vi ~/.firmachain/config/genesis.json

:%s/"stake"/"ufct"/g
4 substitutions on 4 lines
```

#### change gov & voting time

```
vi ~/.firmachain/config/genesis.json

"max_deposit_period": "120s"
"voting_period": "120s"
```

* Changing the gov & voting time should only be done on the privately owned network.

### Register as Validator

```
firmachaind add-genesis-account <wallet_address> <balances>ufct
```



### Setting Staking volume

```
firmachaind gentx <key_name> <balances>ufct --chain-id <chain_id>
Enter keyring passphrase:
Genesis transaction written to "/home/user/.firmachain/config/gentx/gentx-09835f25a665ff508cefac029906ad52686a8612.json"
```



### Run collect-gentxs to integrate genesis.json

```
firmachaind collect-gentxs

{"app_message":{"auth":{"accounts":[{"@type":"/cosmos.auth.v1beta1.BaseAccount","account_number":"0","address":"firma1vxwzvk87atvzv7prz26l6efvz9ne3llmf0wwdx","pub_key":null,"sequence":"0"}],"params":{"max_memo_characters":"256","sig_verify_cost_ed25519":"590","sig_verify_cost_secp256k1":"1000","tx_sig_limit":"7","tx_size_cost_per_byte":"10"}},"authz":{"authorization":[]},"bank":{"balances":[{"address":"firma1vxwzvk87atvzv7prz26l6efvz9ne3llmf0wwdx","coins":[{"amount":"100000000000000","denom":"ufct"}]}],"denom_metadata":[],"params":{"default_send_enabled":true,"send_enabled":[]},"supply":[]},"capability":{"index":"1","owners":[]},"contract":{"contractFileList":[],"contractLogCount":"0","contractLogList":[]},"crisis":{"constant_fee":{"amount":"1000","denom":"stake"}},"distribution":{"delegator_starting_infos":[],"delegator_withdraw_infos":[],"fee_pool":{"community_pool":[]},"outstanding_rewards":[],"params":{"base_proposer_reward":"0.010000000000000000","bonus_proposer_reward":"0.040000000000000000","community_tax":"0.020000000000000000","withdraw_addr_enabled":true},"previous_proposer":"","validator_accumulated_commissions":[],"validator_current_rewards":[],"validator_historical_rewards":[],"validator_slash_events":[]},"evidence":{"evidence":[]},"feegrant":{"allowances":[]},"genutil":{"gen_txs":[{"auth_info":{"fee":{"amount":[],"gas_limit":"200000","granter":"","payer":""},"signer_infos":[{"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AqWJv/BLwXKHNy0q2969XrOGVN0twh9gqo1VRTLamw9l"},"sequence":"0"}]},"body":{"extension_options":[],"memo":"09835f25a665ff508cefac029906ad52686a8612@192.168.20.104:26656","messages":[{"@type":"/cosmos.staking.v1beta1.MsgCreateValidator","commission":{"max_change_rate":"0.010000000000000000","max_rate":"0.200000000000000000","rate":"0.100000000000000000"},"delegator_address":"firma1vxwzvk87atvzv7prz26l6efvz9ne3llmf0wwdx","description":{"details":"","identity":"","moniker":"test1","security_contact":"","website":""},"min_self_delegation":"1","pubkey":{"@type":"/cosmos.crypto.ed25519.PubKey","key":"oEJ7OgjBwvPnw1PGQYb646/Dc/B0pNGVMO2VL+zm/Cs="},"validator_address":"firmavaloper1vxwzvk87atvzv7prz26l6efvz9ne3llmhu94dg","value":{"amount":"50000000000000","denom":"ufct"}}],"non_critical_extension_options":[],"timeout_height":"0"},"signatures":["3Tbk+HJbZdkAxHnKl+IGsyZKyi+/wSnzPNZITWUazNUBt63+ERjmzp+aj5ld7f20gndrqN4tVxVLS/ARtK0Yuw=="]}]},"gov":{"deposit_params":{"max_deposit_period":"172800s","min_deposit":[{"amount":"10000000","denom":"stake"}]},"deposits":[],"proposals":[],"starting_proposal_id":"1","tally_params":{"quorum":"0.334000000000000000","threshold":"0.500000000000000000","veto_threshold":"0.334000000000000000"},"votes":[],"voting_params":{"voting_period":"172800s"}},"ibc":{"channel_genesis":{"ack_sequences":[],"acknowledgements":[],"channels":[],"commitments":[],"next_channel_sequence":"0","receipts":[],"recv_sequences":[],"send_sequences":[]},"client_genesis":{"clients":[],"clients_consensus":[],"clients_metadata":[],"create_localhost":false,"next_client_sequence":"0","params":{"allowed_clients":["06-solomachine","07-tendermint"]}},"connection_genesis":{"client_connection_paths":[],"connections":[],"next_connection_sequence":"0","params":{"max_expected_time_per_block":"30000000000"}}},"mint":{"minter":{"annual_provisions":"0.000000000000000000","inflation":"0.130000000000000000"},"params":{"blocks_per_year":"6311520","goal_bonded":"0.670000000000000000","inflation_max":"0.200000000000000000","inflation_min":"0.070000000000000000","inflation_rate_change":"0.130000000000000000","mint_denom":"stake"}},"nft":{"nftItemCount":"0","nftItemList":[]},"params":null,"slashing":{"missed_blocks":[],"params":{"downtime_jail_duration":"600s","min_signed_per_window":"0.500000000000000000","signed_blocks_window":"100","slash_fraction_double_sign":"0.050000000000000000","slash_fraction_downtime":"0.010000000000000000"},"signing_infos":[]},"staking":{"delegations":[],"exported":false,"last_total_power":"0","last_validator_powers":[],"params":{"bond_denom":"stake","historical_entries":10000,"max_entries":7,"max_validators":100,"unbonding_time":"1814400s"},"redelegations":[],"unbonding_delegations":[],"validators":[]},"transfer":{"denom_traces":[],"params":{"receive_enabled":true,"send_enabled":true},"port_id":"transfer"},"upgrade":{},"vesting":{}},"chain_id":"test-1","gentxs_dir":"/home/firma/.firmachain/config/gentx","moniker":"test1","node_id":"09835f25a665ff508cefac029906ad52686a8612"}
```



### Start your own network node

Congratulations! You are all set to start developing on FirmaChain.

```
firmachaind start
5:05AM INF starting ABCI with Tendermint
5:05AM INF Starting multiAppConn service impl=multiAppConn module=proxy
5:05AM INF Starting localClient service connection=query impl=localClient module=abci-client
5:05AM INF Starting localClient service connection=snapshot impl=localClient module=abci-client
5:05AM INF Starting localClient service connection=mempool impl=localClient module=abci-client
5:05AM INF Starting localClient service connection=consensus impl=localClient module=abci-client
5:05AM INF Starting EventBus service impl=EventBus module=events
5:05AM INF Starting PubSub service impl=PubSub module=pubsub
5:05AM INF Starting IndexerService service impl=IndexerService module=txindex
5:05AM INF ABCI Handshake App Info hash= height=0 module=consensus protocol-version=0 software-version=0.2.8
5:05AM INF ABCI Replay Blocks appHeight=0 module=consensus stateHeight=0 storeHeight=0
5:05AM INF asserting crisis invariants inv=0/11 module=x/crisis name=gov/module-account
5:05AM INF asserting crisis invariants inv=1/11 module=x/crisis name=staking/module-accounts
5:05AM INF asserting crisis invariants inv=2/11 module=x/crisis name=staking/nonnegative-power
5:05AM INF asserting crisis invariants inv=3/11 module=x/crisis name=staking/positive-delegation
5:05AM INF asserting crisis invariants inv=4/11 module=x/crisis name=staking/delegator-shares
5:05AM INF asserting crisis invariants inv=5/11 module=x/crisis name=bank/nonnegative-outstanding
5:05AM INF asserting crisis invariants inv=6/11 module=x/crisis name=bank/total-supply
5:05AM INF asserting crisis invariants inv=7/11 module=x/crisis name=distribution/nonnegative-outstanding
5:05AM INF asserting crisis invariants inv=8/11 module=x/crisis name=distribution/can-withdraw
5:05AM INF asserting crisis invariants inv=9/11 module=x/crisis name=distribution/reference-count
5:05AM INF asserting crisis invariants inv=10/11 module=x/crisis name=distribution/module-account
5:05AM INF asserted all invariants duration=0.815086 height=0 module=x/crisis
5:05AM INF created new capability module=ibc name=ports/transfer
5:05AM INF port binded module=x/ibc/port port=transfer
5:05AM INF claimed capability capability=1 module=transfer name=ports/transfer
5:05AM INF Completed ABCI Handshake - Tendermint and App are synced appHash= appHeight=0 module=consensus
5:05AM INF Version info block=11 p2p=8 tendermint_version=0.34.12
5:05AM INF This node is a validator addr=A442D1FD6E2BD4C03101680E3815A7664930244E module=consensus pubKey=oEJ7OgjBwvPnw1PGQYb646/Dc/B0pNGVMO2VL+zm/Cs=
5:05AM INF P2P Node ID ID=09835f25a665ff508cefac029906ad52686a8612 file=/home/firma/.firmachain/config/node_key.json module=p2p
5:05AM INF Adding persistent peers addrs=[] module=p2p
5:05AM INF Adding unconditional peer ids ids=[] module=p2p
5:05AM INF Add our address to book addr={"id":"09835f25a665ff508cefac029906ad52686a8612","ip":"0.0.0.0","port":26656} book=/home/firma/.firmachain/config/addrbook.json module=p2p
5:05AM INF Starting Node service impl=Node
5:05AM INF Starting pprof server laddr=localhost:6060
5:05AM INF Starting RPC HTTP server on 127.0.0.1:26657 module=rpc-server
5:05AM INF Starting P2P Switch service impl="P2P Switch" module=p2p
5:05AM INF Starting Consensus service impl=ConsensusReactor module=consensus
5:05AM INF Reactor  module=consensus waitSync=false
5:05AM INF Starting State service impl=ConsensusState module=consensus
5:05AM INF Starting baseWAL service impl=baseWAL module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
5:05AM INF Starting Group service impl=Group module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
5:05AM INF Searching for height height=1 max=0 min=0 module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
5:05AM INF Searching for height height=0 max=0 min=0 module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
5:05AM INF Found height=0 index=0 module=consensus wal=/home/firma/.firmachain/data/cs.wal/wal
5:05AM INF Catchup by replaying consensus messages height=1 module=consensus
5:05AM INF Replay: Done module=consensus
5:05AM INF Starting TimeoutTicker service impl=TimeoutTicker module=consensus
5:05AM INF Starting Evidence service impl=Evidence module=evidence
5:05AM INF Starting StateSync service impl=StateSync module=statesync
5:05AM INF Starting PEX service impl=PEX module=pex
5:05AM INF Starting AddrBook service book=/home/firma/.firmachain/config/addrbook.json impl=AddrBook module=p2p
5:05AM INF Starting Mempool service impl=Mempool module=mempool
5:05AM INF Starting BlockchainReactor service impl=BlockchainReactor module=blockchain
5:05AM INF Ensure peers module=pex numDialing=0 numInPeers=0 numOutPeers=0 numToDial=10
5:05AM INF Saving AddrBook to file book=/home/firma/.firmachain/config/addrbook.json module=p2p size=0
5:05AM INF No addresses to dial. Falling back to seeds module=pex
5:05AM INF Timed out dur=4994.122207 height=1 module=consensus round=0 step=1
5:05AM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"AD43A0590BACD05761A5672B941BE46770D107F8D780CC72E3C34932F93A4D44","parts":{"hash":"9D8CB5FE3316747098D358D37284AE052F29F252137ECBDAB12BF542D7AC146F","total":1}},"height":1,"pol_round":-1,"round":0,"signature":"Mp/MiA097Cms3FdOhwDrrDJ3wIZxRhIk2Na18mCjf1chbWqWpq1Hf+5Z9g4pgVHQHCZU9LS6dBKizcG9+5QPDg==","timestamp":"2021-11-24T05:05:13.785504124Z"}
5:05AM INF received complete proposal block hash=AD43A0590BACD05761A5672B941BE46770D107F8D780CC72E3C34932F93A4D44 height=1 module=consensus
5:05AM INF finalizing commit of block hash=AD43A0590BACD05761A5672B941BE46770D107F8D780CC72E3C34932F93A4D44 height=1 module=consensus num_txs=0 root=E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855
5:05AM INF minted coins from module account amount=2059725ufct from=mint module=x/bank
5:05AM INF executed block height=1 module=state num_invalid_txs=0 num_valid_txs=0
5:05AM INF commit synced commit=436F6D6D697449447B5B31303020313839203230352031363120313438203130322032303820373820323420313320363420323330203139382031393620313035203139302031333620323534203135322037382031303320323237203135203230312032333120323132203138362031313020313720313231203739203138325D3A317D
5:05AM INF committed state app_hash=64BDCDA19466D04E180D40E6C6C469BE88FE984E67E30FC9E7D4BA6E11794FB6 height=1 module=state num_txs=0
5:05AM INF indexed block height=1 module=txindex
```

