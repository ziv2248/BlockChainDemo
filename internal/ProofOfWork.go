package internal

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

const MAX_NONCE = int(^uint(0) >> 1)
const DEFECAULTY = 12

type ProofOfWork struct {
	target *big.Int
	block  *Block
}

func InitProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DEFECAULTY))
	fmt.Printf("題目: hash < %d\n", target)

	pow := &ProofOfWork{target, b}
	return pow
}

func (pow *ProofOfWork) prepareData(nounce int) []byte {
	timestamp := []byte(strconv.FormatInt(pow.block.Timestamp, 10))

	payload := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data, timestamp,
		[]byte(strconv.Itoa(nounce))}, []byte{})
	hash := sha256.Sum256(payload)

	data := hash[:]
	return data
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func (pow *ProofOfWork) Proof() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nounce := 0

	fmt.Printf("目標Data: \"%s\" 正在為您挖礦中...\n", pow.block.Data)

	for nounce < MAX_NONCE {
		if nounce%128 == 1 {
			fmt.Println("正在為您挖礦中...")
		}
		data := pow.prepareData(nounce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖到了! hash: %x\n", hash)
			break
		} else {
			nounce++
		}
	}
	return nounce, hash[:]
}
