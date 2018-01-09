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

func generateInstructions(id int, instruction chan<- int) {

	for { // do forever

		opcode := (rand.Intn(5) + 1) // Randomly generate a new opcode (between 1 and 5)

		fmt.Printf("%d   Instruction: %d\n", id, opcode) // Report this to console display

		instruction <- opcode // Send the instruction for retirement
	}
}

//--------------------------------------------------------------------------------
// Executes the instruction by waiting number of seconds instruction gives
//...............................................................................
func executeInstruction(id int, execute <-chan int) {
	for {
		opcode := <-execute
		time.Sleep(time.Second * time.Duration(opcode))

	}
}

//--------------------------------------------------------------------------------
// Retires instructions by writing them to the console
//--------------------------------------------------------------------------------

func retireInstruction(id int, retired <-chan int) {

	for { // do forever
		// Receive an instruction from the generator
		opcode := <-retired

		fmt.Printf("%d       Retired: %d \n", id, opcode) // Report to console
	}
}

//////////////////////////////////////////////////////////////////////////////////
//  Main program, create required channels, then start goroutines in parallel.
//////////////////////////////////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().Unix()) // Seed the random number generator

	// Set up required channel
	opcodes := make([]chan int, 3) //arrayOfChannels

	for i := range opcodes {
		opcodes[i] = make(chan int)
	}
	//opcodes := make(chan int) // channel for flow of generated opcodes

	// Now start the goroutines in parallel.
	fmt.Printf("Start Go routines...\n")

	for i := 0; i < 3; i++ {
		go generateInstructions(i, opcodes[i])
		go executeInstruction(i, opcodes[i])
		go retireInstruction(i, opcodes[i])
	}

	for {
	}

} // end of main /////////////////////////////////////////////////////////////////
