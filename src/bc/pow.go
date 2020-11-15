package bc

import "github.com/apm-dev/go-simple-blockchain/src/crypto"

type POW interface {
	findNonce(b block) int64
	isNonceValid(b block) bool
}

type pow struct {
}

func newPOW() POW {
	return &pow{}
}

//	Find correct nonce by incrementing nonce then hash entire block
//	and check with difficulty level
func (p *pow) findNonce(b block) int64 {
	b.Nonce = 0
	for p.isNonceValid(b) == false {
		b.Nonce += 1
	}
	return b.Nonce
}

func (p *pow) isNonceValid(b block) bool {
	hash, _ := crypto.Hash(crypto.HashSHA256, b)
	return hash[len(hash)-len(DifficultyLvl):] == DifficultyLvl
}
