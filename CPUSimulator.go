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
//instruction struct
//////////////////////////////////////////////////////////////////////////////////
type instruction struct {
	id, opcode int
}

//////////////////////////////////////////////////////////////////////////////////
// Function/Process definitions
//////////////////////////////////////////////////////////////////////////////////

//----------------------------------------------------------------------------------
// Randomly generate an instruction 'opcode' between 1 and 5 and send to the retire function
//----------------------------------------------------------------------------------

func generateInstructions(instructions []chan instruction, done []chan bool) {
	go executeInstruction(instructions[0])
	go executeInstruction(instructions[1])
	go executeInstruction(instructions[2])
	for i := 1; i < 99; i++ { // do forever
		var newInstruction instruction
		newInstruction.id = i
		newInstruction.opcode = (rand.Intn(5) + 1) // Randomly generate a new opcode (between 1 and 5)
		fmt.Printf("Instruction: %d\n", newInstruction.opcode)
		instructions[newInstruction.id%3] <- newInstruction

		// Report this to console display
	}
	for {
	}
}

//--------------------------------------------------------------------------------
// Executes the instruction by waiting number of seconds instruction gives
//...............................................................................
func executeInstruction(execute <-chan instruction) {
	for {
		inst := <-execute
		fmt.Printf("executing instruction %d\n", inst.id)
		time.Sleep(time.Second * time.Duration(inst.opcode))

	}
}

//--------------------------------------------------------------------------------
// Retires instructions by writing them to the console
//--------------------------------------------------------------------------------

func retireInstruction(retired <-chan instruction) {

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
	instructions := make([]chan instruction, 3) //arrayOfChannels
	done := make([]chan bool, 3)
	for i := range instructions {
		instructions[i] = make(chan instruction)
		done[i] = make(chan bool)
	}

	go generateInstructions(instructions, done)
	//opcodes := make(chan int) // channel for flow of generated opcodes

	// Now start the goroutines in parallel.
	fmt.Printf("Start Go routines...\n")
	for {
	}

} // end of main /////////////////////////////////////////////////////////////////
