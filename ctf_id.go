package ctfid

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetConditionID replicates CTF's getConditionId:
// keccak256( abi.encodePacked(oracle(20), questionId(32), outcomeSlotCount(uint256)) )
func GetConditionID(oracle common.Address, questionID [32]byte, outcomeSlotCount uint32) [32]byte {
	var buf []byte
	buf = append(buf, oracle.Bytes()...)    // 20 bytes
	buf = append(buf, questionID[:]...)     // 32 bytes

	// outcomeSlotCount as uint256 (32-byte big-endian)
	var u [32]byte
	binary.BigEndian.PutUint32(u[28:], outcomeSlotCount)
	buf = append(buf, u[:]...)              // 32 bytes

	return crypto.Keccak256Hash(buf)
}

// GetCollectionID (general form, includes parentCollectionId):
// keccak256( abi.encodePacked(parentCollectionId(32), conditionId(32), indexSet(uint256)) )
func GetCollectionID(parentCollectionID [32]byte, conditionID [32]byte, indexSet uint32) [32]byte {
	var buf []byte
	buf = append(buf, parentCollectionID[:]...)
	buf = append(buf, conditionID[:]...)

	var u [32]byte
	binary.BigEndian.PutUint32(u[28:], indexSet)
	buf = append(buf, u[:]...)

	return crypto.Keccak256Hash(buf)
}

// GetCollectionIDSingle (common single-condition shortcut):
// keccak256( abi.encodePacked(conditionId(32), indexSet(uint256)) )
func GetCollectionIDSingle(conditionID [32]byte, indexSet uint32) [32]byte {
	var buf []byte
	buf = append(buf, conditionID[:]...)

	var u [32]byte
	binary.BigEndian.PutUint32(u[28:], indexSet)
	buf = append(buf, u[:]...)

	return crypto.Keccak256Hash(buf)
}

// GetPositionID replicates CTF's getPositionId:
// keccak256( abi.encodePacked(collateralToken(20), collectionId(32)) ) as uint256
func GetPositionID(collateral common.Address, collectionID [32]byte) *big.Int {
	var buf []byte
	buf = append(buf, collateral.Bytes()...)
	buf = append(buf, collectionID[:]...)

	h := crypto.Keccak256Hash(buf)
	return new(big.Int).SetBytes(h.Bytes())
}
