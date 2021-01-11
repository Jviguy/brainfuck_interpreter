package brainfuck_interperter

import (
	"fmt"
	"log"
	"os"
	"sync"
)

//A struct used for running basic brainfuck code.
type BrainFucker struct {
	Ptr    uint16
	Memory map[uint16]uint8
}

func (b BrainFucker) Run(code string) {
	for i := 0; i < len(code); i++ {
		currentChar := code[i]
		switch currentChar {
		case '>':
			//> Increments pointer addr
			b.Ptr++
		case '<':
			//< Decreases Pointer Addr
			b.Ptr--
		case '+':
			//+ Increments the byte value at the said pointer
			b.Memory[b.Ptr]++
		case '-':
			//- Decreases the byte value at the said pointer
			b.Memory[b.Ptr]--
		case '.':
			//. Prints the current byte as a character to the os.Stdout
			fmt.Print(string(b.Memory[b.Ptr]))
		case ',':
			//reads one byte from the stdin and assigns it to our byte value
			by := make([]byte, 1)
			_, err := os.Stdin.Read(by)
			if err != nil {
				log.Fatal(err)
			}
			b.Memory[b.Ptr] = by[0]
		case '[':
			//[ if the value is zero we move to the next instruction after the matching ]
			if b.Memory[b.Ptr] == 0 {
				var loop = 1
				for loop > 0 {
					i--
					var currentChar = code[i]
					if currentChar == ']' {
						loop--
					} else if currentChar == '[' {
						loop++
					}
				}
			}
		case ']':
			//] if the value is non zero we move to the matching [
			if b.Memory[b.Ptr] != 0 {
				var loop = 1
				for loop > 0 {
					i--
					var currentChar = code[i]
					if currentChar == '[' {
						loop--
					} else if currentChar == ']' {
						loop++
					}
				}
			}
		default:
			log.Fatalf("Invalid Operator %v located at: char:%v", string(code[i]), i)
		}
	}
}

//A struct to run my custom brainfuck dialect.
type CustomFucker struct {
	Ptr    uint16
	Memory map[uint16]uint8
	Wg     sync.WaitGroup
}

func (b *CustomFucker) Run(code string) {
	for i := 0; i < len(code); i++ {
		currentChar := code[i]
		switch currentChar {
		case '>':
			//> Increments pointer addr
			b.Ptr++
		case '<':
			//< Decreases Pointer Addr
			b.Ptr--
		case '+':
			//+ Increments the byte value at the said pointer
			b.Memory[b.Ptr]++
		case '-':
			//- Decreases the byte value at the said pointer
			b.Memory[b.Ptr]--
		case '.':
			//. Prints the current byte as a character to the os.Stdout
			fmt.Print(string(b.Memory[b.Ptr]))
		case ',':
			//reads one byte from the stdin and assigns it to our byte value
			by := make([]byte, 1)
			_, err := os.Stdin.Read(by)
			if err != nil {
				log.Fatal(err)
			}
			b.Memory[b.Ptr] = by[0]
		case '[':
			//[ if the value is zero we move to the next instruction after the matching ]
			if b.Memory[b.Ptr] == 0 {
				var loop = 1
				for loop > 0 {
					i--
					var currentChar = code[i]
					if currentChar == ']' {
						loop--
					} else if currentChar == '[' {
						loop++
					}
				}
			}
		case ']':
			//] if the value is non zero we move to the matching [
			if b.Memory[b.Ptr] != 0 {
				var loop = 1
				for loop > 0 {
					i--
					var currentChar = code[i]
					if currentChar == '[' {
						loop--
					} else if currentChar == ']' {
						loop++
					}
				}
			}
		case '{':
			var start = i + 1
			//} starts the matching code in an another thread
			for {
				if i == len(code) {
					log.Fatalf("Unmatched [ found at char:%v", start-1)
				}
				currentChar = code[i]
				if currentChar == '}' {
					break
				}
				i++
			}
			go func() {
				b.Wg.Add(1)
				b.Run(code[start : i-1])
				b.Wg.Done()
			}()
		case '@':
			//@ awaits for all other threads to exit before continuing the program.
			b.Wg.Wait()
		default:
			log.Fatalf("Invalid Operator %v located at: char:%v", string(code[i]), i)
		}
	}
}
