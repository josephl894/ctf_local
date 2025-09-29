package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"module github.com/josephl894/ctf_local"
)

func mustHex32(s string) [32]byte {
	var out [32]byte
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != 32 {
		log.Fatalf("questionId must be 32-byte hex: %v", err)
	}
	copy(out[:], b)
	return out
}

func main() {
	oracleHex := flag.String("oracle", "0x1111111111111111111111111111111111111111", "oracle address")
	qidHex    := flag.String("qid", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "questionId (32-byte hex, no 0x)")
	outcomes  := flag.Uint("n", 2, "outcome slot count")
	indexSet  := flag.Uint("index", 1, "indexSet (e.g., 1 for YES, 2 for NO)")
	collHex   := flag.String("collateral", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", "collateral address")

	flag.Parse()

	oracle := common.HexToAddress(*oracleHex)
	qid := mustHex32(*qidHex)
	cond := ctfid.GetConditionID(oracle, qid, uint32(*outcomes))
	coll := common.HexToAddress(*collHex)

	col := ctfid.GetCollectionIDSingle(cond, uint32(*indexSet))
	pos := ctfid.GetPositionID(coll, col)

	fmt.Printf("conditionId : 0x%x\n", cond)
	fmt.Printf("collectionId: 0x%x\n", col)
	fmt.Printf("positionId  : 0x%x\n", pos)
}
