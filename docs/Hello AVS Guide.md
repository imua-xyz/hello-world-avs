# Hello AVS Guide 

# **Hello World AVS**

Welcome to the **Hello World AVS**. This project demonstrates the simplest functionality you can expect from an AVS, providing a concrete understanding of its basic components.

**GitHub Repository**:

- [Hello World AVS](https://github.com/ExocoreNetwork/hello-world-avs/commits/main/)
- [Exocore CLI](https://github.com/ExocoreNetwork/exocore)

### **AVS User Flow Example**

1. **Request**: An AVS consumer requests the generation and signing of a "Hello World" message.
2. **Event Emission**: The `HelloWorld` contract receives this request and emits a `NewTaskCreated` event.
3. **Task Processing by Operators**:
    - All registered Operators who have staked or delegated assets to the AVS pick up the request.
    - Each Operator generates the requested message, hashes it, and signs the hash using their private key.
4. **Submission**: Operators submit their signed hashes back to the `HelloWorld` AVS contract.
5. **Validation**: If the Operator is registered and meets the minimum staking requirements, their submission is accepted.

# Steps

## Run the node

<aside>
💡  

Prerequistion: [Install dependency](https://docs.exocore.network/validator-setup/compiling-binary-from-source) or download it directly below.

</aside>  

[exocored](Hello%20AVS%20Guide%20(1)%2013d612d2b5df80fe9e3aebb4441852a3/exocored.txt)

```powershell
git clone [https://github.com/ExocoreNetwork/exocore.git](https://github.com/ExocoreNetwork/exocore.git)
git checkout develop
cd exocore
make build
cp build/exocored /usr/bin/exocored
./local_node.sh
```  

Check Available Accounts  检查可用帐户

```powershell
exocored keys list --home ~/.tmp-exocored
- address: exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
  name: dev0
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AopCZPMxuF3pvVk+LV4MNXrejD6HCjMR+8tkB7Y+F4/F"}'
  type: local
  ...
```  

## Register Operator

Fund the account if the balance of the dev0 account is 0, check it with `exocored q bank balances exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx`

```powershell
exocored tx operator register-operator --meta-info "Operator1" --from dev0 --commission-rate 0.5 --commission-max-rate 1 --commission-max-change-rate 1 --home ~/.tmp-exocored --keyring-backend test --fees 50000000000hua --chain-id exocoretestnet_233-1
```  

## Show Operator ECDSA Key

```powershell
exocored keys unsafe-export-eth-key dev0 --home ~/.tmp-exocored
DBDE1905049DEE771ED652F3DC05D57EBFEC0FE1B0ED0310482E9C58DB398834
```  

## Import Keys

```powershell
git clone https://github.com/ExocoreNetwork/hello-world-avs.git
cd hello-world-avs
make build
# avs owner private key sample: D196DCA836F8AC2FFF45B3C9F0113825CCBB33FA1B39737B948503B263ED75AE
# operator private key sample: DBDE1905049DEE771ED652F3DC05D57EBFEC0FE1B0ED0310482E9C58DB398834
# bls private key sample: 1c0599ffc52d512fd5b549fa050833e7d3bba12969d09d70a16441384e5a8a3a
./exokey importKey --key-type ecdsa --private-key {avsOwner_private_key} --output-dir tests/keys/avs.ecdsa.key.json
./exokey importKey --key-type ecdsa --private-key {operator_private_key} --output-dir tests/keys/operator.ecdsa.key.json
./exokey  importKey --key-type bls --private-key {bls_private_key}  --output-dir tests/keys/test.bls.key.json  

```

Continue to configure the `config.yaml` file in `hello-world-avs/config.yaml`

- **operator_address**
    - It is the EIP-55 address of the bench32 address of the operator, check it with

    ```powershell
    exocored debug addr exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
    Address bytes: [175 8 51 107 198 168 209 105 195 74 78 158 77 98 118 227 98 244 117 126]
    Address (hex): AF08336BC6A8D169C34A4E9E4D6276E362F4757E
    Address (EIP-55): 0xAF08336BC6A8D169c34A4e9e4d6276e362F4757E
    Bech32 Acc: exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
    Bech32 Val: exovaloper14uyrx67x4rgkns62f60y6cnkud30gat7nkaagu
    ```  

- **avs_owner_addresses**
    - The bench32 address of the operator must be included in **avs_owner_addresses**

```powershell
# this sets the logger level (true = info, false = debug)
production: false
#The eoa address of Operator, used for sending transactions
operator_address: 0xAF08336BC6A8D169c34A4e9e4d6276e362F4757E
#The eoa address of avs owner, used for sending transactions
avs_owner_address: 0x3e108c058e8066DA635321Dc3018294cA82ddEdf
#Address of deployed AVS contract, if it is empty, will deploy AVS on fly
avs_address: 
# ETH RPC URL
eth_rpc_url: http://127.0.0.1:8545
eth_ws_url: ws://127.0.0.1:8546
avs_ecdsa_private_key_store_path: tests/keys/avs.ecdsa.key.json
operator_ecdsa_private_key_store_path: tests/keys/operator.ecdsa.key.json
bls_private_key_store_path: tests/keys/test.bls.key.json
node_api_ip_port_address: 0.0.0.0:9010
enable_node_api: false
register_operator_on_startup: false
#register avs parameters
avs_name: "hello-avs"
min_stake_amount: 1
avs_owner_addresses:
  - "exo18cggcpvwspnd5c6ny8wrqxpffj5zmhklprtnph"
  - "exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx"
asset_ids:
  - "0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"
avs_unbonding_period: 7
min_self_delegation: 0
epoch_identifier: minute
avs_reward_address: 
avs_slash_address: 
task_address: 
params:
  - 5
  - 7
  - 8
  - 4
#create new task parameters
#Create task intervals,Unit second
create_task_interval: 200
task_response_period: 2
task_challenge_period: 2
threshold_percentage: 100
task_statistical_period: 2
# depoist and delegate params
deposit_amount: 100
delegate_amount: 100
staker: 0xa53f68563D22EB0dAFAA871b6C08a6852f91d627  

```

## Register AVS and Create Task

```powershell
git clone https://github.com/ExocoreNetwork/hello-world-avs.git
cd hello-world-avs
make build
./avs/main --config config.yaml
```  

Result:

```powershell
./avs/main --config config.yaml 
2024-11-13T12:27:10.955Z        INFO    logging/zap_logger.go:49        AVS_ADDRESS env var not set. will deploy avs contract
2024-11-13T12:27:11.644Z        INFO    logging/zap_logger.go:69        tx hash: 0x81477902f1c3432b933ef58a12f9e9990dd7d3913253af687ee412927a6f8f48
2024-11-13T12:27:11.644Z        INFO    logging/zap_logger.go:69        contract address: 0xd83404Cde3A28b751a661521Cb0aD3Cc35B7fa95
2024-11-13T12:27:17.352Z        DEBUG   logging/zap_logger.go:45        Estimating gas and nonce        {"tx": "0x08aa5b66c6f344519670509138e868dbe9e3cfef4bd275a6eb10430db11571b0"}
2024-11-13T12:27:17.365Z        DEBUG   logging/zap_logger.go:45        Getting signer for tx   {"tx": "0xc14fa15c9308af5b9853f776966896bd1236f606a094c125d6c56ebed1ce2abb"}
2024-11-13T12:27:18.064Z        DEBUG   logging/zap_logger.go:45        Sending transaction     {"tx": "0xc14fa15c9308af5b9853f776966896bd1236f606a094c125d6c56ebed1ce2abb"}
2024-11-13T12:27:20.072Z        INFO    logging/zap_logger.go:69        tx hash: 0x08aa5b66c6f344519670509138e868dbe9e3cfef4bd275a6eb10430db11571b0
2024-11-13T12:27:20.072Z        INFO    logging/zap_logger.go:69        Starting avs.
2024-11-13T12:27:20.072Z        INFO    logging/zap_logger.go:69        Avs owner set to send new task every 200 seconds
2024-11-13T12:27:20.072Z        INFO    logging/zap_logger.go:49        Avs sending new task
2024-11-13T12:27:20.074Z        WARN    logging/zap_logger.go:53        AVS USD value is zero or negative       {"value": "0.000000000000000000", "attempt": 1, "max_attempts": 10}
```  

It will show the contract address of the created AVS, and also the `config.yaml` will be overwritten with `avs_address`, `avs_reward_address` , `avs_slash_address`, `task_address`, Since there is no operator opt in to the AVS, so the AVS USD value is zero, it is not allowed to create new task.

## Operator Opt into AVS

```powershell
 exocored tx operator opt-into-avs <avs_contract_address> --from dev0 --home ~/.tmp-exocored --keyring-backend test --fees 50000000000hua --chain-id exocoretestnet_233-1
```  

Verify operator opt into the avs:

```powershell
exocored q operator get-avs-list exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
avs_list:
- 0xd83404cde3a28b751a661521cb0ad3cc35b7fa95
```  

## Deposit and Make Delegation

Prerequisite:

```powershell
git clone https://github.com/ExocoreNetwork/hello-world-avs.git
cd hello-world-avs
make build
```  

Deposit and delegate with operator module in hello-world-avs

```powershell
./operator/main --config config.yaml
```  

Result:

```powershell
2024-11-13T12:34:44.882Z        INFO    logging/zap_logger.go:49        current epoch  is :     {"currentEpoch": 2837}
2024-11-13T12:34:44.882Z        INFO    logging/zap_logger.go:49        Execute Phase One Submission Task       {"currentEpoch": 2837, "startingEpoch": 2836, "taskResponsePeriod": 2}
2024-11-13T12:34:44.882Z        INFO    logging/zap_logger.go:49        Submitting task response for task response period       {"taskAddr": "0xd83404Cde3A28b751a661521Cb0aD3Cc35B7fa95", "taskId": 2, "operator-addr": "0xAF08336BC6A8D169c34A4e9e4d6276e362F4757E"}
2024-11-13T12:34:45.811Z        DEBUG   logging/zap_logger.go:45        Estimating gas and nonce        {"tx": "0x44c374ddf58b24aef97d1845eda88ab9478ac610dc4194275bfb6bfb255b6491"}
2024-11-13T12:34:45.821Z        DEBUG   logging/zap_logger.go:45        Getting signer for tx   {"tx": "0x2fbc117007fb43faa196ad4855bdb68cbfe97148587ebaf967f18bda9796af9c"}
2024-11-13T12:34:45.968Z        INFO    logging/zap_logger.go:69        tx hash: 0x94dbf2ce7c41fbe01abf8977f8623629e357f35eee727839156dd487f4bfb35a
2024-11-13T12:34:45.971Z        INFO    logging/zap_logger.go:49        current epoch  is :     {"currentEpoch": 2837}
2024-11-13T12:34:45.971Z        INFO    logging/zap_logger.go:49        Execute Phase Two Submission Task       {"currentEpoch": 2837, "startingEpoch": 2833, "taskResponsePeriod": 2, "taskStatisticalPeriod": 2}
2024-11-13T12:34:45.971Z        INFO    logging/zap_logger.go:49        Submitting task response for statistical period {"taskAddr": "0xd83404Cde3A28b751a661521Cb0aD3Cc35B7fa95", "taskId": 1, "operator-addr": "0xAF08336BC6A8D169c34A4e9e4d6276e362F4757E"}
```  

check the result of `./avs/main --config config.yaml` again.

```powershell
2024-11-13T12:30:09.594Z        INFO    logging/zap_logger.go:69        Starting avs.
2024-11-13T12:30:09.594Z        INFO    logging/zap_logger.go:69        Avs owner set to send new task every 200 seconds
2024-11-13T12:30:09.594Z        INFO    logging/zap_logger.go:49        Avs sending new task
2024-11-13T12:30:10.305Z        DEBUG   logging/zap_logger.go:45        Estimating gas and nonce        {"tx": "0xb273fc6702500277b74236b18d92c5ee529e7f949d4b794645059a2afa99400a"}
2024-11-13T12:30:10.320Z        DEBUG   logging/zap_logger.go:45        Getting signer for tx   {"tx": "0x1ec222369d5459b6e92b094da6d4feacb362b3635022cf51b8c20eb20f801fbd"}
2024-11-13T12:30:11.057Z        DEBUG   logging/zap_logger.go:45        Sending transaction     {"tx": "0x1ec222369d5459b6e92b094da6d4feacb362b3635022cf51b8c20eb20f801fbd"}
2024-11-13T12:30:13.061Z        INFO    logging/zap_logger.go:49        Transaction not yet mined       {"hash": "0x871859c4beda7559ae434e6c5506e16b43e2cd6e20e77b9d34a7a36bba6fb9d3"}
2024-11-13T12:30:15.063Z        INFO    logging/zap_logger.go:69        tx hash: 0xb273fc6702500277b74236b18d92c5ee529e7f949d4b794645059a2afa99400a
2024-11-13T12:33:29.593Z        INFO    logging/zap_logger.go:49        sendNewTask-num:        {"taskNum": 2}
2024-11-13T12:33:29.593Z        INFO    logging/zap_logger.go:49        Avs sending new task
2024-11-13T12:33:30.529Z        DEBUG   logging/zap_logger.go:45        Estimating gas and nonce        {"tx": "0x212ec28c1ee4eae4733d2d2274c4e036c814e40a1ae4f287e566cdccc347296d"}
2024-11-13T12:33:30.543Z        DEBUG   logging/zap_logger.go:45        Getting signer for tx   {"tx": "0xf50abe8ff4082bdf5b5d7aab1730e7c6aab2720e5ae18686742deda559491706"}
2024-11-13T12:33:31.395Z        DEBUG   logging/zap_logger.go:45        Sending transaction     {"tx": "0xf50abe8ff4082bdf5b5d7aab1730e7c6aab2720e5ae18686742deda559491706"}
2024-11-13T12:33:33.401Z        INFO    logging/zap_logger.go:69        tx hash: 0x212ec28c1ee4eae4733d2d2274c4e036c814e40a1ae4f287e566cdccc347296d
2024-11-13T12:36:49.591Z        INFO    logging/zap_logger.go:49        sendNewTask-num:        {"taskNum": 3}
```  

## Check TaskInfo with exocore

```powershell
exocored q avs TaskInfo 0xd83404Cde3A28b751a661521Cb0aD3Cc35B7fa95 1
actual_threshold: "7766279631452241920"
err_signed_operators: []
hash: bN9XyPiULWhVmBM38FXdvn9gE46xQe91Sewfs+LAkXs=
name: Z1Ujo
no_signed_operators: []
operator_active_power:
  operator_power_list:
  - active_power: "315786.000000000000000000"
    operator_addr: exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
opt_in_operators:
- exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
signed_operators:
- exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
starting_epoch: "2833"
task_challenge_period: "2"
task_contract_address: 0xd83404cde3a28b751a661521cb0ad3cc35b7fa95
task_id: "1"
task_response_period: "2"
task_statistical_period: "2"
task_total_power: "315786.000000000000000000"
threshold_percentage: "100"
```  

## Check Task Submit Info with exocore

If the phase is `PHASE_DO_COMMIT`, it is the expected result that two phase submit result is completed.

```powershell
exocored q avs SubmitTaskResult 0xd83404Cde3A28b751a661521Cb0aD3Cc35B7fa95 1 exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
info:
  bls_signature: jncXW+w8ZHcVnU/gK3F2GUktj0ZEdNxPLost64TN9Pl/KDBhl04ae/3PTxeHZ36LEU10IW2+/wdJO7njDnPUZFf9MokjbhgoTHlWsyCpJLoKBpadpBrAWcsYid856TwE
  operator_address: exo14uyrx67x4rgkns62f60y6cnkud30gat7a20fdx
  phase: PHASE_DO_COMMIT
  task_contract_address: 0xd83404Cde3A28b751a661521Cb0aD3Cc35B7fa95
  task_id: "1"
  task_response: eyJUYXNrSUQiOjEsIk51bWJlclN1bSI6MX0=
  task_response_hash: 0x91b8eabcc462c7ded2c0427c3bbd3e4b05c5a73ea55571ff114b43dd9aeff767
```  