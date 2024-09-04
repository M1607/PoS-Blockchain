/* This Go code manages blockchain operations using global variables for
Blockchain, TempBlocks, and channels for CandidateBlocks and announcements.
The PickWinner function selects a winning validator after a 20-second delay,
 creating a pool of eligible validators based on the number of tokens they’ve
 staked. Validators are added to the pool multiple times, proportional to their
 stake, to increase their chances of being chosen.


@author (M. Hirschfeld)
@version (September 4, 2024)
*/

package blockchain

import (
	"time"
)

// Global Variables for managing blockchain and temp blocks
var Blockchain []Block
var TempBlocks []Block
var CandidateBlocks = make(chan Block)
var announcements = make(chan string)

// Pick the winner based on the tokens staked
func PickWinner() {
	time.Sleep(20 * time.Second)

	var lotteryPool []string
	for _, block := range TempBlocks {
		// Check if validator is eligible
		if stake, ok := Validators[block.Validator]; ok {
			// Add validator to pool multiple times based on their stake
			for i := 0; i < stake; i++ {
				lotteryPool = append(lotteryPool, block.Validator)
			}
		}
	}
}
