package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var profile = flag.Bool("profile", false, "Set to true to enable profiling to cpuprofile.out")

func main() {
	flag.Parse()

	filename := "generate.xml"

	// Enable profiling.
	if *profile {
		f, err := os.Create("cpuprofile.out")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	count := countHouses(filename)
	fmt.Printf("Go: Found %v houses\n", count)
}

func countHouses(filename string) int {
	houses := 0

	file, err := os.Open(filename)
	handle(err)
	defer file.Close()
	buffer := bufio.NewReaderSize(file, 1024*1024*256) // 33554432
	decoder := xml.NewDecoder(buffer)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "House" {
				houses++
			}
		}
	}

	return houses
}

func handle(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
