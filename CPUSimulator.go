// Pipeline Simulator OOO simple version 1.go
// The Super-Scalar Processor Simulator - simple out-of-order version, runs forever.
// A.Oram 2017
//fidninjjn
package main

//iigijb/dinagn
// Imported packages

import (
	"fmt"       // for console I/O
	"math/rand" // for randomly creating opcodes
	"time"      // for the random number generator and 'executing' opcodes
)

//////////////////////////////////////////////////////////////////////////////////
// Function/Process definitions
//////////////////////////////////////////////////////////////////////////////////

//----------------------------------------------------------------------------------
// Randomly generate an instruction 'opcode' between 1 and 5 and send to the retire function
//----------------------------------------------------------------------------------

func generateInstructions(instruction chan<- int) {

	for { // do forever

		opcode := (rand.Intn(5) + 1) // Randomly generate a new opcode (between 1 and 5)

		fmt.Printf("Instruction: %d\n", opcode) // Report this to console display

		instruction <- opcode // Send the instruction for retirement
	}
}

//--------------------------------------------------------------------------------
// Retires instructions by writing them to the console
//--------------------------------------------------------------------------------
func retireInstruction(retired <-chan int) {

	for { // do forever
		// Receive an instruction from the generator
		opcode := <-retired

		fmt.Printf("Retired: %d \n", opcode) // Report to console
	}
}

//////////////////////////////////////////////////////////////////////////////////
//  Main program, create required channels, then start goroutines in parallel.
//////////////////////////////////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().Unix()) // Seed the random number generator

	// Set up required channel

	opcodes := make(chan int) // channel for flow of generated opcodes

	// Now start the goroutines in parallel.
	fmt.Printf("Start Go routines...\n")

	go generateInstructions(opcodes)
	go retireInstruction(opcodes)

	for {
	}

} // end of main /////////////////////////////////////////////////////////////////
