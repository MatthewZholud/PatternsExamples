package main

import "fmt"

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	pr printer
}

func (m *mac) print()  {
	fmt.Println("Print request for mac")
	m.pr.printFile()
}

func (m *mac) srtPrinter(p printer)  {
	m.pr = p
}

type win struct {
	pr printer
}

func (m *win) print()  {
	fmt.Println("Print request for win")
	m.pr.printFile()
}

func (m *win) srtPrinter(p printer)  {
	m.pr = p
}





type printer interface {
	printFile()
}

type epson struct {
}

func (p *epson) printFile(){
	fmt.Println("Printing by epson printer")
}

type hp struct {
}

func (p *hp) printFile(){
	fmt.Println("Printing by hp printer")
}



func main()  {
	hpPrinter := &hp{}
	macComputer := &mac{}
	macComputer.srtPrinter(hpPrinter)
	macComputer.print()
}
