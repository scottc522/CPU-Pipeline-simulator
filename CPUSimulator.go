// Pipeline Simulator
//b4005596 Scott Chapman
package main

// Imported packages
import (
	"fmt"       // for console I/O
	"math/rand" // for randomly creating opcodes
	"time"      // for the random number generator and 'executing' opcodes
)

//////////////////////////////////////////////////////////////////////////////////
//instruction struct
//opcode is how long the instruction runs for when executed
//////////////////////////////////////////////////////////////////////////////////
type instruction struct {
	id, opcode int
}

//----------------------------------------------------------------------------------
// Randomly generates an instruction and sends it to one of the execution pipelines
//----------------------------------------------------------------------------------
//It takes an array of channels. Each of the execution pieplines listens to one channel in that array
func generateInstructions(instructions []chan instruction) {
	//Retire is a channel that the first function in the sort and retire function will be listening too.
	retire := make(chan instruction)
	//Each one of the three execution pipelines is passed a different instruction channel and the same retire channel
	//The three go piplelines are activated
	go executeInstruction(0, instructions[0], retire)
	go executeInstruction(1, instructions[1], retire)
	go executeInstruction(2, instructions[2], retire)
	//The first step in the sort-retire function is activated
	go retireInstruction(retire)
	//Runs 101 times to generate instructions with ids from 1 to 100
	for i := 1; i < 101; i++ { // do forever
		var newInstruction instruction
		//The instruction id is equal to the itteration of the for loop.
		newInstruction.id = i
		// Randomly generates a new opcode (between 1 and 5) for this instruction
		newInstruction.opcode = (rand.Intn(5) + 1)
		//Sends this new instruction to one of the instructions channels.
		//The channel that is is sent to is determined by the modulous of it's id. Ensuring that each
		//execution pipleine is used evenly.
		instructions[newInstruction.id%3] <- newInstruction

	}
	//Set up a dummy instruction, and send it to each of the execute pipelines so they know to shutdown.
	var finalInstruction instruction
	finalInstruction.id = 101
	finalInstruction.opcode = 5
	time.Sleep(time.Second * 5 / 10)
	for k := 0; k < 3; k++ {

		instructions[k] <- finalInstruction

	}

}

//--------------------------------------------------------------------------------
// Executes the instruction by waiting number of seconds instruction gives
//...............................................................................
func executeInstruction(id int, execute <-chan instruction, retire chan<- instruction) {
	for {

		inst := <-execute
		//inst recieves an instruction from the execute channel
		//If the id of this instruction isn't 101 (not the id of the 'dummy' instruction then)
		if inst.id != 101 {
			//Sleep for one second multiplied by the number given by the opcode.
			time.Sleep(time.Second * time.Duration(inst.opcode))
			//Then pass the instruction into the retire channal.
			retire <- inst
		} else {
			//Else, it must be the 'dummy' instruction
			retire <- inst
			//pass the dummy instruction into the retire pipleine.
			//go to the key word f outside of the for loop. Effectivley ending this function
			goto F

		}

	}
F: //Print to show that pipeline has shut down
	fmt.Println(id, "execution has finished.")
}

//--------------------------------------------------------------------------------
// Pipe sort to ensure instructions are in the correct order
//--------------------------------------------------------------------------------

func retireInstruction(retired <-chan instruction) {

	for {
		//Create Array of 6 instruction channels
		pipeSort := make([]chan instruction, 6)
		//setting channels up.
		for i := 0; i < 6; i++ {
			pipeSort[i] = make(chan instruction)

		}

		//initalise go routines
		for i := 0; i < 5; i++ {
			go pipeSorter(i, pipeSort[i], pipeSort[i+1])
		}
		//inatialise retire function and pass in the last pipesort channel
		go retire(pipeSort[5])
		//Get the first instruction from the pipeline and give it to the variable myInstruction
		myInstruction := <-retired

		for {
			//Pass the next instruction from the channel into the next instruction variable.
			nextInstruction := <-retired
			//Check to see if the instruction is the dummy instruction
			if nextInstruction.id == 101 {
				//If it is the dummy instruction pass the current instruction into the start of the pipesort
				pipeSort[0] <- myInstruction
				//Pass the dummy instruction into the pipesort
				pipeSort[0] <- nextInstruction
				//Empty the dummy instructions from the other two execute functions.
				<-retired
				<-retired
				goto R
				//Escape the for loop so this function can end
			}
			//if myInstruction is not the dummy instruction then compare the current instruction
			//With the next instruction.
			if myInstruction.id > nextInstruction.id {
				//If the current instruction is bigger than the next instruction pass the next instruction
				// into the start of the pipeline sort.
				pipeSort[0] <- nextInstruction
			} else {
				//Otherwise pass the current instruction into the pipeline sort
				pipeSort[0] <- myInstruction
				//Set the current instruction to equal the current next instruction
				myInstruction = nextInstruction

			}

		}

	}
	//Lable R, so that inner for loop can be escaped when dummy instructions are reached.
R:
	fmt.Println("Pipeline entry is complete")
}

func pipeSorter(id int, myInstructions <-chan instruction, nextInstructions chan<- instruction) {

	//It's current instruction is the first thing passed to it from the instructions channel
	myInstruction := <-myInstructions

	for {
		//It takes its next instruction from the channel of instructions it has to process
		nextInstruction := <-myInstructions
		//Check to see if the next instruction is the dummy instruction
		if nextInstruction.id == 101 {
			//If it is the dummy instruction it passes its current instruction and the 'dummy' to the next
			//part of the pipeline via the nextInstructions channel.
			nextInstructions <- myInstruction
			nextInstructions <- nextInstruction
			//Goes to the lable F which will take it out of this for loop and allow the function to end
			goto F
		}
		//Checks which of the two instructions has the bigger id
		if myInstruction.id > nextInstruction.id {
			//If the current instruction has the bigger id it passes the next instruction to the next stage
			// of the pipeline via the next instructions channel.
			nextInstructions <- nextInstruction
		} else {
			//if the current instruction is smaller than the next one, pass the current instruction to the next stage
			//of the pipeline.
			nextInstructions <- myInstruction
			myInstruction = nextInstruction
		}
	}
F:
}

func retire(myInstructions <-chan instruction) {
	for {
		//Recieve instructions from the my instruction channel.
		myInstruction := <-myInstructions
		//If its the dummy then escape the for loop by going to F
		if myInstruction.id == 101 {
			goto F
		}
		//Other print out the instruction
		fmt.Println(myInstruction)
	F:
	}

}

//////////////////////////////////////////////////////////////////////////////////
//  Main program, create required channels, then start goroutines in parallel.
//////////////////////////////////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().Unix()) // Seed the random number generator

	//Create array of channels
	instructions := make([]chan instruction, 3) //arrayOfChannels
	//Set up instruction channels
	for i := range instructions {
		instructions[i] = make(chan instruction)
	}
	//Pass the array of chanels into the generate instruction funtion
	// and activate it
	go generateInstructions(instructions)
	//Keep this function going forever
	for {
	}

} // end of main /////////////////////////////////////////////////////////////////
