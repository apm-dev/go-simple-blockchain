package bc

import (
	"errors"
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
	InitialCoins  = 100
)

type Blockchain interface {
	GetChain() []*block
	NewTrx(sender string, recipient string, amount float32)
	Mine() error
}

type blockchain struct {
	memPool   MemPool
	pow       POW
	chain     Chain
	myAddress string
}

func NewBlockchain(yourAddress string) Blockchain {
	return &blockchain{
		memPool:   newMemPool(),
		pow:       newPOW(),
		chain:     newChain(createGenesisBlock(yourAddress)),
		myAddress: yourAddress,
	}
}

func createGenesisBlock(address string) block {
	gHash, _ := crypto.Hash(crypto.HashSHA256, "this is your bc:"+address+":genesis block")
	return block{
		Index:     1,
		Timestamp: time.Now().UTC().Unix(),
		Trxs: []*trx{{
			Sender:    "",
			Recipient: address,
			Amount:    InitialCoins,
		}},
		Nonce:        0,
		PreviousHash: gHash,
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

func (b *blockchain) Mine() error {
	lastBlock := b.chain.lastBlock()
	if lastBlock == nil {
		return errors.New("there is no genesis block in the chain")
	}
	pHash, err := crypto.Hash(crypto.HashSHA256, lastBlock)
	if err != nil {
		return err
	}
	rewardTrx := trx{
		Sender:    "",
		Recipient: b.myAddress,
		Amount:    MiningReward,
	}
	b.memPool.addTrx(&rewardTrx)
	newBlock := block{
		Index:        lastBlock.Index + 1,
		Timestamp:    time.Now().UTC().Unix(),
		Trxs:         b.memPool.getTrxs(),
		Nonce:        0,
		PreviousHash: pHash,
	}

	newBlock.Nonce = b.pow.findNonce(newBlock)
	err = b.chain.addBlock(newBlock)
	if err != nil {
		return err
	}
	b.memPool.clear()
	return nil
}
