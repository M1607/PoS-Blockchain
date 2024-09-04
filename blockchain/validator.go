/* This Go package handles connections from blockchain validators. It 
prompts the validator to enter their token balance and BPM (beats per 
minute), validating and storing them. If inputs are invalid, the validator 
is removed. A new block is created based on the last block, using the 
validator's BPM and address, and is sent to the CandidateBlocks channel 
for potential inclusion in the blockchain.


@author (M. Hirschfeld)
@version (September 4, 2024)
*/ 

package blockchain

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

// Validators hold information about the token staked by validators
var Validators = make(map[string]int)

// Handle the connection for each validator
func HandleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Enter token balance
	fmt.Fprint(conn, "Enter token balance: ")
	tokenBalance, _ := reader.ReadString('\n')
	stake, err := strconv.Atoi(tokenBalance)
	if err != nil {
		fmt.Fprintf(conn, "Invalid token balance \n")
		return
	}

	// Assign validator an address
	validatorAddress := CalculateHash((tokenBalance))
	Validators[validatorAddress] = stake

	// Enter BPM
	fmt.Fprint(conn, "Enter BPM: ")
	bpmInput, _ := reader.ReadString('\n')
	BPM, err := strconv.Atoi(bpmInput)
	if err != nil {
		delete(Validators, validatorAddress)
		fmt.Fprintf(conn, "Invalid BPM\n")
		return
	}

	// Create new block
	lastBlock := Blockchain[len(Blockchain)-1]
	newBlock := GenerateBlock(lastBlock, BPM, validatorAddress)

	// Send block to be considered
	CandidateBlocks <- newBlock
}