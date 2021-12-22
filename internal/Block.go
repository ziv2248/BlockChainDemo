package internal

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	payload := bytes.Join([][]byte{
		b.PrevBlockHash,
		b.Data,
		timestamp,
		[]byte("0")}, []byte{})
	hash := sha256.Sum256(payload)
	b.Hash = hash[:]
}

func CreateBlock(Data []byte, PrevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		Data,
		PrevBlockHash,
		[]byte{},
		0,
	}

	if len(PrevBlockHash) == 0 {
		block.SetHash()
	} else {
		pow := InitProofOfWork(block)
		nonce, hash := pow.Proof()
		block.Nonce = nonce
		block.Hash = hash[:]
	}

	return block
}
