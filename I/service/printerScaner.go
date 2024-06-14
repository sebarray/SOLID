package service

type Printer interface {
	Print(string)
}

type Scanner interface {
	Scan(string)
}

// SERVICE PRINTER
type PrinterService struct {
}

func (p *PrinterService) Print(text string) {
	println(text)
}

// SERVICE SCANNER
type ScannerService struct {
}

func (s *ScannerService) Scan(text string) {
	println(text)
}
