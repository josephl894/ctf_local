# CTF-ID Go Utils

This module provides the **core ID calculations** used by the **Gnosis Conditional Tokens Framework (CTF)** — the plumbing for building prediction-market backends.

## Features
- **Condition ID** – unique ID for each market question  
- **Collection ID (single)** – outcome subset (e.g. YES / NO)  
- **Collection ID (parent)** – for hierarchical combinations  
- **Position ID** – ERC-1155 token ID = `keccak256(collateral | collection)`  
- **BN254 Combine** – order-independent & associative set-combine for multi-condition markets  
- Unit-tested with deterministic & BN254-combine properties  
- CLI to print IDs for given inputs

## Quick Start
```bash
# install deps
go mod tidy

# run tests
go test ./...

# build CLI
go build -o bin/ctf_id_cli ./cmd/ctf_id_cli

# example run
./bin/ctf_id_cli   -oracle 0x1111111111111111111111111111111111111111   -qid aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   -n 2 -index 1   -collateral 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
```

Output:
```
conditionId : 0x...
collectionId: 0x...
positionId  : 0x...
```

## When to Use
- **Binary / single-condition** market → `GetConditionID` + `GetCollectionIDSingle` + `GetPositionID`
- **Multi-condition / subset tokens** → combine base collections with `CombineCollectionsBN254`

## Project Layout
```
ctf_local/
  ctf_id.go              # keccak-based IDs
  ctf_id_bn254.go        # BN254 unordered combine
  ctf_test.go            # unit tests
  cmd/ctf_id_cli/main.go # demo CLI
```

## Notes
- Keccak helpers mirror Solidity `abi.encodePacked` (address 20B + bytes32 + uint256 BE).
- BN254 combine makes `A⊕B == B⊕A` and associative for complex markets.
- CLI is for local testing / generating golden vectors; not a production service.

---

> This README should be enough for a teammate to install, test, build, and understand **what** the code does and **why** it exists.
