package main

import "os"
import "fmt"

func main() {
	f, _ := os.Create("generate.xml")
	defer f.Close()
	lines := 10000000 // 10M

	f.WriteString(`<?xml version="1.0" encoding="utf-8"?>` + "\n")
	f.WriteString(`<LIST>` + "\n")
	for i := 0; i < lines; i++ {
		f.WriteString(`<ELEMENT ATTRIBUTE1="" ATTRIBUTE2="" />`)
	}
	f.WriteString("\n" + `</LIST>` + "\n")

	fmt.Println("Done")
}
