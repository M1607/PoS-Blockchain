/* This program sets up a Proof-of-Stake blockchain system. It begins by 
creating the genesis block, which is the first block in the blockchain. 
The server listens for validator connections on a specified port and handles 
token stakes and block proposals. Proposed blocks are collected, and the system 
selects a validator to add their block based on staked tokens. The server keeps
 running indefinitely to manage the blockchain.


@author (M. Hirschfeld)
@version (September 4, 2024)
*/ 

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"PoS_Blockchain/blockchain"
)

func main() {
	// Create the genesis block
	genesisBlock := blockchain.Block{
		Index:		0,
		Timestamp: 	time.Now().String(),
		BPM:		0,
		PrevHash: 	"",
		Hash: 		blockchain.CalculateHash(time.Now().String()),
	}
	blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)

	// Start the TCP Server
	port := os.Getenv("ADDR")
	if port == "" {
		port = "9000"
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	// Listen for incoming connections
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go blockchain.HandleConn(conn)
		}
	}()

	//Start block selection process
	go func() {
		for {
			block := <-blockchain.CandidateBlocks
			blockchain.TempBlocks = append(blockchain.TempBlocks, block)
		}
	}()

	// Announce the winnners and update hte blockchain
	go blockchain.PickWinner()

	// Keep the main function running
	select {}
}