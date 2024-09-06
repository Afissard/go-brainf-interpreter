package interpreter

import (
	"brainf/config"
	"fmt"
	"log"
)

func (u *Univers) Run(script []uint) error {
	if config.Global.Debug {
		log.Println("running the script ...")
	}
	loopEntry := []int{} // list all the loop entry cells pos
	loopExit := []int{}  // list all the loop exit cells pos
	for i := 0; i < len(script); i++ {
		if config.Global.Debug {
			u.showMemory()
		}

		switch script[i] {
		case MovRight:
			if u.pointer+1 >= uint(len(u.memory)-1) {
				u.pointer = 0
			} else {
				u.pointer++
			}
		case MovLeft:
			if u.pointer == 0 {
				u.pointer = uint(len(u.memory) - 1)
			} else {
				u.pointer--
			}
		case IncCell:
			// TODO: prevent overflow
			u.memory[u.pointer]++
		case DecCell:
			// TODO: prevent overflow
			u.memory[u.pointer]--
		case Output:
			fmt.Printf("%c", u.memory[u.pointer])
		case Input:
			var in rune
			fmt.Print("Waiting for input: ")
			fmt.Scanf("%c", &in)
			u.memory[u.pointer] = int8(in)
		case JmpPast:
			loopEntry = append(loopEntry, i)
			if u.memory[u.pointer] == 0 { // go to the exit the loop
				i = loopExit[len(loopExit)-1] // jump to the exit instruction
			}
		case JmpBack:
			loopExit = append(loopExit, i)
			if u.memory[u.pointer] != 0 { // go back into the loop
				i = loopEntry[len(loopEntry)-1]
			} else { // exit the loop
				loopEntry = loopEntry[:len(loopEntry)-1]
				loopExit = loopExit[:len(loopExit)-1]
			}
		default:
			return fmt.Errorf("unknown instruction : %d : %s", script[i], instructionSet[script[i]])
		}
	}
	return nil
}
