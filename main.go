package main

import (
	"github.com/jviguy/brainfuck_interpreter/brainfuck_interpreter"
	"go.uber.org/atomic"
	"log"
	"os"
	"sync"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Error No Input File Stated!")
	}
	//The said code to be executed
	if os.Args[1] == "-c" {
		if len(os.Args) < 2 {
			log.Fatal("Error No Input File Stated!")
		}
		bytes, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatal("Error The Stated Input file couldn't be found!")
		}
		code := string(bytes)
		//Memory pointer.
		var ptr atomic.Value
		ptr.Store(uint16(0))
		//The said memory the interep will use.
		memory := make(map[uint16]uint8)
		b := &brainfuck_interpreter.CustomFucker{Memory: memory, Ptr: ptr, Wg: sync.WaitGroup{}, Stdout: os.Stdout, Stdin: os.Stdin}
		b.Run(code)
	} else {
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal("Error The Stated Input file couldn't be found!")
		}
		code := string(bytes)
		//Memory pointer.
		var ptr atomic.Value
		ptr.Store(uint16(0))
		//The said memory the interep will use.
		memory := make(map[uint16]uint8)
		b := &brainfuck_interpreter.BrainFucker{Memory: memory, Ptr: ptr, Stdout: os.Stdout, Stdin: os.Stdin}
		b.Run(code)
	}
}
