package main

import (
	"github.com/jviguy/brainfuck_interperter/brainfuck_interperter"
	"go.uber.org/atomic"
	"io/ioutil"
	"log"
	"os"
	"sync"
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
		b := &brainfuck_interperter.CustomFucker{Memory: memory, Ptr: ptr, Wg: sync.WaitGroup{}, Stdout: os.Stdout, Stdin: os.Stdin}
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
		b := &brainfuck_interperter.BrainFucker{Memory: memory, Ptr: ptr, Stdout: os.Stdout, Stdin: os.Stdin}
		b.Run(code)
	}
}
