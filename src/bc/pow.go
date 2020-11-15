package bc

import "github.com/apm-dev/go-simple-blockchain/src/crypto"

//	To adjust the difficulty of the algorithm,
//	we could modify the number of leading zeroes.
//	But 4 is sufficient.
//	You’ll find out that the addition of a single leading zero
//	makes a huge difference to the time required to find a solution.
const DifficultyLvl = "0000"

var instance POW

type POW interface {
	FindNonce(b block) int64
	isNonceValid(b block) bool
}

type pow struct {
}

//	Singleton pattern
func GetPOW() POW {
	if instance == nil {
		instance = &pow{}
	}
	return instance
}

//	Find correct nonce by incrementing nonce then hash entire block
//	and check with difficulty level
func (p *pow) FindNonce(b block) int64 {
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
