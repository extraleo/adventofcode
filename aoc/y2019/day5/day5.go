package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func day5() []int {
	inputs:=utils.SpiltInputGetInt("input.txt", ",")
	for address := 0; address < len(inputs); address++ {
		opcode := inputs[address] % 100
		mode1st := 0
		mode2nd := 0
		if inputs[address] >= 100 {
			if ((inputs[address] % 1000) - (inputs[address] % 100)) == 100 {
				mode1st = 1
			}
			if ((inputs[address] % 10000) - (inputs[address] % 1000)) == 1000 {
				mode2nd = 1
			}
		}

		switch opcode {
		case 1:
			// +
			parameter1 := inputs[address+1]
			if mode1st == 0 {
				parameter1 = inputs[parameter1]
			}
			parameter2 := inputs[address+2]
			if mode2nd == 0 {
				parameter2 = inputs[parameter2]
			}

			inputs[inputs[address+3]] = parameter1 + parameter2
			address = address + 3
		case 2:
			// *
			parameter1 := inputs[address+1]
			if mode1st == 0 {
				parameter1 = inputs[parameter1]
			}
			parameter2 := inputs[address+2]
			if mode2nd == 0 {
				parameter2 = inputs[parameter2]
			}

			inputs[inputs[address+3]] = parameter1 * parameter2
			address = address + 3
		case 3:
			// Read
			scanner := bufio.NewScanner(os.Stdin)
			for {
				fmt.Print("Input: ")
				scanner.Scan()
				text := scanner.Text()
				if len(text) != 0 {
					inputs[inputs[address+1]] = utils.Atoi(text)
					break
				}

			}
			address = address + 1
		case 4:
			// Write
			fmt.Printf("(%d)\n", inputs[inputs[address+1]])
			address = address + 1
		case 5:
			// Jump-if-true
			parameter1 := inputs[address+1]
			if mode1st == 0 {
				parameter1 = inputs[parameter1]
			}
			parameter2 := inputs[address+2]
			if mode2nd == 0 {
				parameter2 = inputs[parameter2]
			}
			if parameter1 != 0 {
				address = parameter2 - 1 
				continue
			}
			address = address + 2
		case 6:
			// Jump-if-false
			parameter1 := inputs[address+1]
			if mode1st == 0 {
				parameter1 = inputs[parameter1]
			}
			parameter2 := inputs[address+2]
			if mode2nd == 0 {
				parameter2 = inputs[parameter2]
			}

			if parameter1 == 0 {
				address = parameter2 - 1
				continue
			}
			address = address + 2
		case 7:
			// less-than
			parameter1 := inputs[address+1]
			if mode1st == 0 {
				parameter1 = inputs[parameter1]
			}
			parameter2 := inputs[address+2]
			if mode2nd == 0 {
				parameter2 = inputs[parameter2]
			}
			inputs[inputs[address+3]] = 0
			if parameter1 < parameter2 {
				inputs[inputs[address+3]] = 1
			}
			address = address + 3
		case 8:
			// equals
			parameter1 := inputs[address+1]
			if mode1st == 0 {
				parameter1 = inputs[parameter1]
			}
			parameter2 := inputs[address+2]
			if mode2nd == 0 {
				parameter2 = inputs[parameter2]
			}

			inputs[inputs[address+3]] = 0
			if parameter1 == parameter2 {
				inputs[inputs[address+3]] = 1
			}
			address = address + 3
		case 99:
			// Halt
			return inputs
		}
	}

	return inputs
}


func main() {
	day5()
}
