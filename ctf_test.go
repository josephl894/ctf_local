package ctfid

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func mustHex32(s string) [32]byte {
	var out [32]byte
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != 32 {
		panic("bad hex32")
	}
	copy(out[:], b)
	return out
}

func TestConditionIDDeterministic(t *testing.T) {
	oracle := common.HexToAddress("0x1111111111111111111111111111111111111111")
	qid := mustHex32("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	got1 := GetConditionID(oracle, qid, 2)
	got2 := GetConditionID(oracle, qid, 2)
	if got1 != got2 {
		t.Fatalf("conditionId not deterministic")
	}
}

func TestConditionIDDifferentInputs(t *testing.T) {
	oracle := common.HexToAddress("0x1111111111111111111111111111111111111111")
	qid1 := mustHex32("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	qid2 := mustHex32("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	id1 := GetConditionID(oracle, qid1, 2)
	id2 := GetConditionID(oracle, qid2, 2)
	if id1 == id2 {
		t.Fatalf("different questionIds should not collide")
	}
}

func Test_GetCollectionID_SingleVsParentZeroAreDifferent(t *testing.T) {
    oracle := common.HexToAddress("0x2222222222222222222222222222222222222222")
    qid := mustHex32("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc")
    cond := GetConditionID(oracle, qid, 2)

    // Base (no parent) collection ID for YES
    single := GetCollectionIDSingle(cond, 1)

    // “Parent=0” form is a different constructor and should NOT equal single
    var zero [32]byte
    withParentZero := GetCollectionID(zero, cond, 1)

    if single == withParentZero {
        t.Fatalf("single-collection form must differ from parent=0 constructor")
    }
}


func TestPositionIDDependsOnCollateral(t *testing.T) {
	oracle := common.HexToAddress("0x3333333333333333333333333333333333333333")
	qid := mustHex32("dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd")
	cond := GetConditionID(oracle, qid, 2)
	cYes := GetCollectionIDSingle(cond, 1)

	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // sample
	dai  := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")

	pUSDC := GetPositionID(usdc, cYes)
	pDAI  := GetPositionID(dai, cYes)
	if pUSDC.Cmp(pDAI) == 0 {
		t.Fatalf("different collateral must give different positionId")
	}
	if pUSDC.Cmp(big.NewInt(0)) <= 0 || pDAI.Cmp(big.NewInt(0)) <= 0 {
		t.Fatalf("position ids must be positive")
	}
}
