package ctf_local

import (
	"crypto/sha256"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"golang.org/x/crypto/sha3"
)

// mapIDToG1: keccak(id||"CTF:BN254") → reduce mod r → s*G1
func mapIDToG1(id [32]byte) (bn254.G1Affine, error) {
	// 1) domain-separated keccak
	h := sha3.NewLegacyKeccak256()
	h.Write(id[:])
	h.Write([]byte("CTF:BN254"))
	sum := h.Sum(nil) // 32 bytes

	// 2) reduce to scalar mod r using fr.Element
	var s fr.Element
	s.SetBytes(sum) // automatically reduces mod r

	// 3) scalar multiply base generator
	var gen bn254.G1Affine
	gen.SetGenerator()
	var out bn254.G1Affine
	// Convert scalar to big.Int for mul
	b := new(big.Int)
	s.BigInt(b)
	out.ScalarMultiplication(&gen, b)
	return out, nil
}

// g1Add returns (A + B) in G1
func g1Add(a, b bn254.G1Affine) bn254.G1Affine {
	var aj, bj, sumj bn254.G1Jac
	aj.FromAffine(&a)
	bj.FromAffine(&b)
	sumj.AddAssign(&aj)
	sumj.AddAssign(&bj)
	var sum bn254.G1Affine
	sum.FromJacobian(&sumj)
	return sum
}

// pointToID: compress G1 point and keccak → [32]byte
func pointToID(p bn254.G1Affine) [32]byte {
	// Use uncompressed (x||y) and keccak; you can switch to compressed if you prefer.
	xy := p.Marshal() // []byte; gnark-crypto returns uncompressed (x||y)
	k := sha3.NewLegacyKeccak256()
	k.Write(xy)
	var out [32]byte
	copy(out[:], k.Sum(nil))
	return out
}

// CombineCollectionsBN254: unordered/associative combine of two collection IDs.
// same inputs in any order → same output.
func CombineCollectionsBN254(parentA, parentB [32]byte) ([32]byte, error) {
	if parentA == parentB {
		// idempotent combine
		return parentA, nil
	}
	a, err := mapIDToG1(parentA)
	if err != nil {
		return [32]byte{}, err
	}
	b, err := mapIDToG1(parentB)
	if err != nil {
		return [32]byte{}, err
	}
	sum := g1Add(a, b)
	return pointToID(sum), nil
}
