package main

import (
	"i/handler"
	"i/service"
)

// que carajos es ?
// La segregación de interfaces consiste en dividir interfaces grandes
// en interfaces más pequeñas y específicas. Esto permite que las clases implementen
// solo los métodos que necesitan, haciendo el sistema más flexible,
// mantenible y fácil de entender.
func main() {
	standardPrinter := &handler.StandardPrinter{
		Printer: &service.PrinterService{},
	}
	standardPrinter.Print("Hello, World!")

	multiFunction := &handler.MultiFunction{
		Printer: &service.PrinterService{},
		Scanner: &service.ScannerService{},
	}
	multiFunction.Print("Hello, World!")
	multiFunction.Scan("Hello, World!")

}
