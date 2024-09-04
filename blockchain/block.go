/* This Go package defines the core structure and functions for managing 
a blockchain. The Block struct holds the block's index, timestamp, BPM, 
previous and current hashes, and the validator's address. The CalculateHash 
function generates a SHA256 hash, while GenerateBlock creates a new block 
by incrementing the index and calculating its hash. CalculateBlockHash 
computes a hash based on the block’s data, and IsBlockValid ensures the 
integrity of the chain by validating the new block's index, hash, and 
previous hash. These components enable block creation and validation in 
the blockchain.


@author (M. Hirschfeld)
@version (September 4, 2024)
*/ 

package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index int
	Timestamp string
	BPM int
	PrevHash string
	Hash string
	Validator string
}

// Hashing function to calculate has for block
func CalculateHash (input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Generate a new block
func GenerateBlock (oldBlock Block, BPM int, address string) Block {
	newBlock := Block{
		Index:	oldBlock.Index + 1,
		Timestamp: time.Now().String(),
		BPM:		BPM,
		PrevHash: 	oldBlock.Hash,
		Validator: 	address,
	}
	newBlock.Hash = CalculateBlockHash(newBlock)
	return newBlock
}

// Calculate Block's Hash
func CalculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return CalculateHash(record)
}

// Check block validity by comparing hash and previous hash
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}