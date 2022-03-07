package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"strconv"
)

// реализация урезанного алгоритма Hashcash

const (
	targetBits = 24            //Сложность, на которой блок был добыт
	maxNonce   = math.MaxInt64 //Ограниченнее перебора
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

func (pow ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}

func (pow ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	return nonce, hash[:]
}

func (pow ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			[]byte(strconv.FormatInt(pow.block.Timestamp, 16)),
			[]byte(strconv.FormatInt(int64(targetBits), 16)),
			[]byte(strconv.FormatInt(int64(nonce), 16)),
		},
		[]byte{},
	)
}
