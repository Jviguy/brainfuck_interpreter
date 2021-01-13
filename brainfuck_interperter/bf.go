package brainfuck_interperter

import (
	"go.uber.org/atomic"
	"log"
	"os"
	"sync"
)

//A struct used for running basic brainfuck code.
type BrainFucker struct {
	Ptr    atomic.Value
	Stdout *os.File
	Stdin  *os.File
	Memory map[uint16]uint8
}

func (b BrainFucker) Run(code string) {
	for i := 0; i < len(code); i++ {
		currentChar := code[i]
		switch currentChar {
		case '>':
			//> Increments pointer addr
			b.Ptr.Store(b.Ptr.Load().(uint16) + 1)
		case '<':
			//< Decreases Pointer Addr
			b.Ptr.Store(b.Ptr.Load().(uint16) - 1)
		case '+':
			//+ Increments the byte value at the said pointer
			b.Memory[b.Ptr.Load().(uint16)]++
		case '-':
			//- Decreases the byte value at the said pointer
			b.Memory[b.Ptr.Load().(uint16)]--
		case '.':
			//. Prints the current byte as a character to the os.Stdout
			_, err := b.Stdout.WriteString(string(b.Memory[b.Ptr.Load().(uint16)]))
			if err != nil {
				log.Fatal(err)
			}
		case ',':
			//reads one byte from the stdin and assigns it to our byte value
			by := make([]byte, 1)
			_, err := b.Stdin.Read(by)
			if err != nil {
				log.Fatal(err)
			}
			b.Memory[b.Ptr.Load().(uint16)] = by[0]
		case '[':
			//[ if the value is zero we move to the next instruction after the matching ]
			if b.Memory[b.Ptr.Load().(uint16)] == 0 {
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
			if b.Memory[b.Ptr.Load().(uint16)] != 0 {
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
	Ptr    atomic.Value
	Stdout *os.File
	Stdin  *os.File
	Memory map[uint16]uint8
	Wg     sync.WaitGroup
}

func (b *CustomFucker) Run(code string) {
	var i atomic.Value
	for i.Store(0); i.Load().(int) < len(code); i.Store(i.Load().(int) + 1) {
		currentChar := code[i.Load().(int)]
		switch currentChar {
		case '>':
			//> Increments pointer addr
			b.Ptr.Store(b.Ptr.Load().(uint16) + 1)
		case '<':
			//< Decreases Pointer Addr
			b.Ptr.Store(b.Ptr.Load().(uint16) - 1)
		case '+':
			//+ Increments the byte value at the said pointer
			b.Memory[b.Ptr.Load().(uint16)]++
		case '-':
			//- Decreases the byte value at the said pointer
			b.Memory[b.Ptr.Load().(uint16)]--
		case '.':
			//. Prints the current byte as a character to the os.Stdout
			_, err := b.Stdout.WriteString(string(b.Memory[b.Ptr.Load().(uint16)]))
			if err != nil {
				log.Fatal(err)
			}
		case ',':
			//reads one byte from the stdin and assigns it to our byte value
			by := make([]byte, 1)
			_, err := b.Stdin.Read(by)
			if err != nil {
				log.Fatal(err)
			}
			b.Memory[b.Ptr.Load().(uint16)] = by[0]
		case '[':
			//[ if the value is zero we move to the next instruction after the matching ]
			if b.Memory[b.Ptr.Load().(uint16)] == 0 {
				var loop = 1
				for loop > 0 {
					i.Store(i.Load().(int) - 1)
					var currentChar = code[i.Load().(int)]
					if currentChar == ']' {
						loop--
					} else if currentChar == '[' {
						loop++
					}
				}
			}
		case ']':
			//] if the value is non zero we move to the matching [
			if b.Memory[b.Ptr.Load().(uint16)] != 0 {
				var loop = 1
				for loop > 0 {
					i.Store(i.Load().(int) - 1)
					var currentChar = code[i.Load().(int)]
					if currentChar == '[' {
						loop--
					} else if currentChar == ']' {
						loop++
					}
				}
			}
		case '{':
			var start = i.Load().(int) + 1
			//} starts the matching code in an another thread
			for {
				if i.Load().(int) == len(code) {
					log.Fatalf("Unmatched [ found at char:%v", start-1)
				}
				currentChar = code[i.Load().(int)]
				if currentChar == '}' {
					break
				}
				i.Store(i.Load().(int) + 1)
			}
			go func() {
				b.Wg.Add(1)
				b.Run(code[start : i.Load().(int)-1])
				b.Wg.Done()
			}()
		case '@':
			//@ awaits for all other threads to exit before continuing the program.
			b.Wg.Wait()
		default:
			log.Fatalf("Invalid Operator %v located at: char:%v", string(code[i.Load().(int)]), i.Load().(int))
		}
	}
}
