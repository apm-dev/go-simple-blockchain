package bc

import (
	"github.com/apm-dev/go-simple-blockchain/src/crypto"
	"time"
)

const (
	//	To adjust the difficulty of the algorithm,
	//	we could modify the number of leading zeroes.
	//	But 4 is sufficient.
	//	You’ll find out that the addition of a single leading zero
	//	makes a huge difference to the time required to find a solution.
	DifficultyLvl = "0000"
	MiningReward  = 12.5
)

type Blockchain interface {
	GetChain() []*block
	NewTrx(sender string, recipient string, amount float32)
	Mine(minerPKH string) error
}

type blockchain struct {
	memPool   MemPool
	pow       POW
	chain     Chain
}

func NewBlockchain() Blockchain {
	return &blockchain{
		memPool:   newMemPool(),
		pow:       newPOW(),
		chain:     newChain(),
	}
}

func (b *blockchain) GetChain() []*block {
	return b.chain.getBlocks()
}

func (b *blockchain) NewTrx(sender string, recipient string, amount float32) {
	b.memPool.addTrx(&trx{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	})
}

func (b *blockchain) Mine(minerPKH string) error {
	var pHash string
	var newBlockIndex int32

	lastBlock := b.chain.lastBlock()
	//	Check is it genesis block or not
	if lastBlock == nil {
		pHash = ""
		newBlockIndex = 1
	} else {
		var err error
		pHash, err = crypto.Hash(crypto.HashSHA256, lastBlock)
		if err != nil {
			return err
		}
		newBlockIndex = lastBlock.Index + 1
	}

	//	Create reward trx and add it to mempool
	rewardTrx := trx{
		Sender:    "",
		Recipient: minerPKH,
		Amount:    MiningReward,
	}
	b.memPool.addTrx(&rewardTrx)

	//	create new block template
	newBlock := block{
		Index:        newBlockIndex,
		Timestamp:    time.Now().UTC().Unix(),
		Trxs:         b.memPool.getTrxs(),
		Nonce:        0,
		PreviousHash: pHash,
	}
	//	Find and set a nonce for new block
	newBlock.Nonce = b.pow.findNonce(newBlock)

	//	Add new block to the chain
	err := b.chain.addBlock(newBlock)
	if err != nil {
		return err
	}
	//	When block created should clear mempool for new trxs
	b.memPool.clear()
	return nil
}
