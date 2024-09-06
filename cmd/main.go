package main

import (
	"brainf/config"
	"brainf/interpreter"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	flag.BoolVar(&debugFlag, "debug", false, "activate the debugger")
	flag.StringVar(&fileFlag, "file", "", "path to the brainf script")
	flag.Parse()
}

func main() {
	// Initialisation
	var err error
	config.Global.Set(debugFlag, fileFlag)

	_, err = os.Stat(config.Global.FilePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Panic(fmt.Errorf("brainf script not found at '%s'", config.Global.FilePath))
	}

	// Interpreter
	bfScript := interpreter.Transpile(config.Global.FilePath)
	mainUnivers := interpreter.Univers{}
	err = mainUnivers.Run(bfScript)

	if err != nil {
		log.Panic(err)
	}
}
