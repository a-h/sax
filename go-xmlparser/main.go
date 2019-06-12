package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tamerh/xml-stream-parser"
	"log"
	"os"
	"runtime/pprof"
)

var profile = flag.Bool("profile", false, "Set to true to enable profiling to xmlparser-cpuprofile.out")

func main() {
	flag.Parse()

	filename := "generate.xml"

	// Enable profiling.
	if *profile {
		f, err := os.Create("xmlparser-cpuprofile.out")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	count := countElements(filename)
	fmt.Printf("Go: Found %v elements\n", count)
}

func countElements(filename string) int {
	numElements := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer file.Close()

	buffer := bufio.NewReaderSize(file, 1024*1024) // 1048576
	parser := xmlparser.NewXmlParser(buffer, "ELEMENT")

	for range *parser.Stream() {
		numElements++
	}

	return numElements
}
