package handler

import "i/service"

// multi function printer
type MultiFunction struct {
	Printer service.Printer
	Scanner service.Scanner
}

func (m *MultiFunction) Print(text string) {
	m.Printer.Print(text)
}

func (m *MultiFunction) Scan(text string) {
	m.Scanner.Scan(text)
}

//	standard printer
type StandardPrinter struct {
	Printer service.Printer
}

func (s *StandardPrinter) Print(text string) {
	s.Printer.Print(text)
}
