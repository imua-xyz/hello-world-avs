# Hello World AVS
# Overview

Welcome to the Hello World AVS. This project shows you the simplest functionality you can expect from an AVS. It will give you a concrete understanding of the basic components.

## Workflow
1. AVS consumer requests a "Hello World" message to be generated and signed.
2. HelloWorld contract receives the request and emits a NewTaskCreated event for the request.
3. All Operators who are registered to the AVS and has staked, delegated assets takes this request. Operator generates the requested message, hashes it, and signs the hash with their private key.
4. Each Operator submits their signed hash back to the HelloWorld AVS contract.
5. If the Operator is registered to the AVS and has the minimum needed stake, the submission is accepted.

# Installation

`make install`

Or check out the latest release.

# Quick Start

1.`avs --config config.yaml`
2.`operator --config config.yaml`
3.`operator --config config.yaml`
4.`exokey import --key-type ecdsa {pri_key}`


### eg.
#### config.yaml (path/config.yaml)
```
# this sets the logger level (true = info, false = debug)
production: false
#The eoa address of Operator, used for sending transactions
operator_address: 0xce5b680d1fd259ada4820e9314bcf0723bdb1287
#The eoa address of avs owner, used for sending transactions
avs_owner_address: 0x3e108c058e8066DA635321Dc3018294cA82ddEdf
#Address of deployed AVS contract
avs_address: 0xfF8f8297BEF982ac6ED7a203e144D9fa4F0FcE31


# ETH RPC URL
eth_rpc_url: http://127.0.0.1:8545
eth_ws_url: ws://127.0.0.1:8546
avs_ecdsa_private_key_store_path: tests/keys/avs.ecdsa.key.json
operator_ecdsa_private_key_store_path: tests/keys/operator.ecdsa.key.json
bls_private_key_store_path: tests/keys/test.bls.key.json
node_api_ip_port_address: 0.0.0.0:9010
enable_node_api: true
register_operator_on_startup: false
  
#register avs parameters
avs_name: "hello-avs"
min_stake_amount: 1
avs_owner_addresses:
  - "exo18cggcpvwspnd5c6ny8wrqxpffj5zmhklprtnph"
  - "exo1sc9kjykz6qehauzmhjympsktdjaw4d99dksgrk"
asset_ids:
  - "0xdac17f958d2ee523a2206206994597c13d831ec7_0x65"
avs_unbonding_period: 7
min_self_delegation: 0
epoch_identifier: minute
avs_reward_address: 0x4f5CDaE0B1afeB0473dEb5AC4F9912409BBBBb72
avs_slash_address:  0x4f5CDaE0B1afeB0473dEb5AC4F9912409BBBBb72
task_address:  0x4f5CDaE0B1afeB0473dEb5AC4F9912409BBBBb72
params:
  - 5
  - 7
  - 8
  - 4

#create new task parameters
#Create task intervals,Unit second
create_task_interval: 500
task_response_period: 2
task_challenge_period: 2
threshold_percentage: 100
task_statistical_period: 2
```
All addresses are in Ethereum format.

# Private Key Management for AVS and Operator

## Overview
This document provides a comprehensive guide to generating and managing private keys for the AVS (Actively Validated Service) system using the `exokey` utility.

## Key Paths
- AVS ECDSA Private Key: `tests/keys/avs.ecdsa.key.json`
- Operator ECDSA Private Key: `tests/keys/operator.ecdsa.key.json`
- BLS Private Key: `tests/keys/test.bls.key.json`

## Prerequisites
- Ensure `exokey` is installed(use "make exokey" command)
- Have private keys ready for import

##  Steps

### 1. Import AVS ECDSA Private Key
```bash
# Import AVS private key
./exokey importKey --key-type ecdsa --private-key {avs_private_key}  --output-dir /tests/keys/avs.ecdsa.key.json
# Output: tests/keys/avs.ecdsa.key.json
```

### 2. Import Operator ECDSA Private Key
```bash
# Import Operator private key
./exokey importKey --key-type ecdsa --private-key {operator_private_key} --output-dir /tests/keys/operator.ecdsa.key.json
# Output: tests/keys/operator.ecdsa.key.json
```

### 3. Import BLS Private Key for operator
```bash
# Import BLS private key
./exokey  importKey --key-type bls --private-key {bls_private_key}  --output-dir /tests/keys/test.bls.key.json
# Output: tests/keys/test.bls.key.json
```

## Key Management Best Practices
- Keep private keys secure
- Never share private keys
- Use environment-specific key management
- Rotate keys periodically

## Verification
```bash
# Verify key files exist
ls tests/keys/
```
 
