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
	retire := make(chan instruction)
	go executeInstruction(instructions[0], retire)
	go executeInstruction(instructions[1], retire)
	go executeInstruction(instructions[2], retire)
	go retireInstruction(retire)

	for i := 1; i < 100; i++ { // do forever
		var newInstruction instruction
		newInstruction.id = i
		//	fmt.Println(newInstruction.id)
		newInstruction.opcode = (rand.Intn(5) + 1) // Randomly generate a new opcode (between 1 and 5)
		//	fmt.Printf("Instruction: %d\n", newInstruction.opcode)
		instructions[newInstruction.id%3] <- newInstruction
		// Report this to console display
	}

}

//--------------------------------------------------------------------------------
// Executes the instruction by waiting number of seconds instruction gives
//...............................................................................
func executeInstruction(execute <-chan instruction, retire chan<- instruction) {
	for {

		inst := <-execute
		//	fmt.Printf("executing instruction %d\n", inst.id)

		time.Sleep(time.Second * time.Duration(inst.opcode) / 2)
		retire <- inst

	}
}

//--------------------------------------------------------------------------------
// Retires instructions by writing them to the console
//--------------------------------------------------------------------------------

func retireInstruction(retired <-chan instruction) {

	for {
		//retireie := <-retired
		//fmt.Println(retireie.id)
		// do forever
		// Receive an instruction from the generator

		pipeSort := make([]chan instruction, 6) //arrayOfChannels

		for i := 0; i < 6; i++ {
			pipeSort[i] = make(chan instruction)

		}
		for i := 0; i < 5; i++ {
			go pipeSorter(i, pipeSort[i], pipeSort[i+1])
		}
		go retire(pipeSort[5])
		myInstruction := <-retired
		//for {
		//	fmt.Println(myInstruction.id)
		//	myInstruction = <-retired
		//}
		for {

			//fmt.Println(myInstruction)

			//select {
			//case <-retired:
			nextInstruction := <-retired
			//for {
			//	fmt.Println(nextInstruction)
			//	nextInstruction = <-retired
			//}
			//fmt.Println(myInstruction.id)
			//fmt.Println(nextInstruction.id)
			//	fmt.Println("Options are ", myInstruction.id, "and  ", nextInstruction.id)
			if myInstruction.id > nextInstruction.id {
				pipeSort[0] <- nextInstruction
				//fmt.Println("i chose ", nextInstruction)
			} else {

				//	fmt.Println("I chose ", nextInstruction)
				pipeSort[0] <- myInstruction
				myInstruction = nextInstruction
				//	}

			}

		}

		///fmt.Printf("Retired: %d \n", opcode.id) // Report to console
	}
}
func pipeSorter(id int, myInstructions <-chan instruction, nextInstructions chan<- instruction) {

	//fmt.Println(id, myInstruction.id)
	myInstruction := <-myInstructions
	for {
		//println(id)
		//select {
		//case <-myInstructions:

		nextInstruction := <-myInstructions
		//fmt.Println("My choice is between  ", myInstruction.id, "and  ", nextInstruction.id)
		if myInstruction.id < nextInstruction.id {
			//	fmt.Println("I CHOSE  ", nextInstruction.id)
			nextInstructions <- nextInstruction
		} else {
			//	fmt.Println("I CHOSE  ", myInstruction.id)
			nextInstructions <- myInstruction
			myInstruction = nextInstruction
		}
	}

}

func retire(myInstructions <-chan instruction) {
	for {
		//<-myInstructions
		myInstruction := <-myInstructions
		fmt.Println(myInstruction)
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
