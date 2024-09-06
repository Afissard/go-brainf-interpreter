package interpreter

import (
	"brainf/config"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInstruction(currentChar string) (instructionVal uint, err error) {
	err = nil
	var i uint = 0
	for i < uint(len(instructionSet)) {
		if currentChar == instructionSet[i] {
			return i, err
		}
		i++
	}
	err = fmt.Errorf("%s isn't an instruction character", currentChar)
	return 0, err
}

func Transpile(filePath string) (encodedScript []uint) {
	if config.Global.Debug {
		log.Println("transposing the script ...")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanRunes)
	for reader.Scan() {
		instructionVal, err := getInstruction(reader.Text())

		if err == nil {
			encodedScript = append(encodedScript, instructionVal)
		}
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

	return encodedScript
}
