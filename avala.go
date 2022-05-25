package main

import (
	"avala/lcg"
	"bufio"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
)

var method string
var initialNumber, seed, sampleSpace int

const defaultSampleSpace = 0xff_fff_fff

func main() {
	runtime.GOMAXPROCS(10)
	err := godotenv.Load(".env.dev")

	if err != nil {
		panic("Env is not loaded")
	}

	flag.StringVar(&method, "method", "lvg", "Method selector: (lvg)")
	flag.IntVar(&seed, "seed", 11, "Seed number (must be a co-prime to sampleSpace)")
	flag.IntVar(&initialNumber, "initial-number", 0x1234, "Initial number to start the sequence.")
	flag.IntVar(&sampleSpace, "sample-space", defaultSampleSpace, "Initial number to start the sequence.")
	flag.Parse()

	switch method {
	case "lvg":
		var primeSampleSpace int
		scanner := bufio.NewScanner(os.Stdin)

		if sampleSpace == defaultSampleSpace {
			primeSampleSpace = 4_294_967_311
		} else {
			primeSampleSpace = lcg.FindClosestSampleSpace(sampleSpace)
		}

		gen := lcg.LehmerGenerator(initialNumber, seed, primeSampleSpace)

		fmt.Printf("Press Enter to generate a new number.\nEnter 'exit' to shut down the application.")

		for scanner.Scan() {
			scannedText := scanner.Text()

			if scannedText == "exit" {
				break
			}

			random := gen()

			for {
				if random > sampleSpace {
					break
				}
				random = gen()
			}

			if random == initialNumber {
				// Would take about 4.3 billion for 8 digit hexadecimal to exhaust.
				fmt.Printf("Sequence is going to repeat.")
			}

			fmt.Printf("%08x\n", random)
		}
	}
}
